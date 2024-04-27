package hasher

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type (
	SecretCrypto interface {
		Name() string
		Hash(raw string, arg any) string
		Validate(raw, hashed string) bool
	}
	bcryptCrypto struct {
	}
	argon2Crypto struct {
	}
	Argon2Argument struct {
		Memory      uint32
		Iterations  uint32
		Parallelism uint8
		SaltLength  uint32
		KeyLength   uint32
	}
)

var (
	DefaultArgon2Argument = Argon2Argument{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
)

func (a argon2Crypto) Name() string {
	return "$a2id$"
}

func (a argon2Crypto) Hash(raw string, arg any) string {
	if a, ok := arg.(Argon2Argument); ok {
		salt, err := generateRandomBytes(a.SaltLength)
		if err != nil {
			return ""
		}
		hash := argon2.IDKey([]byte(raw), salt, a.Iterations, a.Memory, a.Parallelism, a.KeyLength)
		b64Salt := base64.RawStdEncoding.EncodeToString(salt)
		b64Hash := base64.RawStdEncoding.EncodeToString(hash)
		return fmt.Sprintf("$a2id$%x$%x$%x$%x$%s$%s", argon2.Version, a.Memory, a.Iterations, a.Parallelism, b64Salt, b64Hash)
	}
	return ""

}
func (a argon2Crypto) Validate(raw, hashed string) bool {
	var version uint
	var memory, iterations uint32
	var parallelism uint8
	var salt, hash string
	idx := strings.LastIndex(hashed, "$")
	if idx == -1 {
		return false
	}
	hash = hashed[idx+1:]
	n, err := fmt.Sscanf(hashed[:idx], "$a2id$%x$%x$%x$%x$%s", &version, &memory, &iterations, &parallelism, &salt)
	if err != nil || n != 5 {
		return false
	}
	bSalt, err := base64.RawStdEncoding.DecodeString(salt)
	if err != nil {
		return false
	}
	bHash, err := base64.RawStdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}
	newHash := argon2.IDKey([]byte(raw), bSalt, iterations, memory, parallelism, uint32(len(bHash)))
	if subtle.ConstantTimeCompare(newHash, bHash) == 1 {
		return true
	}
	return false
}
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s bcryptCrypto) Name() string {
	return "$2a$"
}
func (s bcryptCrypto) Hash(raw string, arg any) string {
	if v, ok := arg.(int); !ok {
		return ""
	} else {
		bin := []byte(raw)
		if len(bin) > 72 {
			bin = bin[:72]
		}
		if b, err := bcrypt.GenerateFromPassword(bin, v); err != nil {
			return ""
		} else {
			return string(b)
		}
	}
}

func (s bcryptCrypto) Validate(raw, hashed string) bool {
	bin := []byte(raw)
	if len(bin) > 72 {
		bin = bin[:72]
	}
	return bcrypt.CompareHashAndPassword([]byte(hashed), bin) == nil
}

var (
	Argon2id SecretCrypto = argon2Crypto{}
	BCrypt   SecretCrypto = bcryptCrypto{}
)

func PasswordValidate(raw, hashed string) bool {
	if strings.HasPrefix(hashed, "$a2id$") {
		return Argon2id.Validate(raw, hashed)
	} else {
		return BCrypt.Validate(raw, hashed)
	}
}

func TotpGenerate(opts totp.GenerateOpts) (string, error) {
	v, err := totp.Generate(opts)
	if err != nil {
		return "", err
	}
	return v.URL(), err
}
func BcryptHash(raw string, cost int) string {
	return BCrypt.Hash(raw, cost)
}
func Argon2Hash(raw string, opt *Argon2Argument) string {
	ag := Argon2Argument{}
	if opt != nil {
		ag = *opt
	}
	if ag.Memory == 0 {
		ag.Memory = DefaultArgon2Argument.Memory
	}
	if ag.Iterations == 0 {
		ag.Iterations = DefaultArgon2Argument.Iterations
	}
	if ag.Parallelism == 0 {
		ag.Parallelism = DefaultArgon2Argument.Parallelism
	}
	if ag.SaltLength == 0 {
		ag.SaltLength = DefaultArgon2Argument.SaltLength
	}
	if ag.KeyLength == 0 {
		ag.KeyLength = DefaultArgon2Argument.KeyLength
	}

	return Argon2id.Hash(raw, ag)
}
func TotpValidate(code, def string) bool {
	if strings.HasPrefix(def, "otpauth://") {
		if k, err := otp.NewKeyFromURL(def); err != nil {
			return totp.Validate(code, def)
		} else if v, err := totp.ValidateCustom(code, k.Secret(), time.Now(), totp.ValidateOpts{
			Period:    uint(k.Period()),
			Skew:      1.,
			Digits:    k.Digits(),
			Algorithm: k.Algorithm(),
		}); err != nil {
			return false
		} else {
			return v
		}
	}
	return totp.Validate(code, def)
}
func TotpCode(def string) string {
	if strings.HasPrefix(def, "otpauth://") {
		if k, err := otp.NewKeyFromURL(def); err != nil {
			panic(err)
		} else if v, err := totp.GenerateCodeCustom(k.Secret(), time.Now(), totp.ValidateOpts{
			Period:    uint(k.Period()),
			Skew:      1.,
			Digits:    k.Digits(),
			Algorithm: k.Algorithm(),
		}); err != nil {
			panic(err)
		} else {
			return v
		}
	} else if v, err := totp.GenerateCode(def, time.Now()); err != nil {
		panic(err)
	} else {
		return v
	}
}

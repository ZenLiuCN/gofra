package hasher

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
)

func ObscureInt(i ...int) string {
	buf := binary.AppendVarint(nil, int64(len(i)))
	for _, i2 := range i {
		buf = binary.AppendVarint(buf, int64(i2))
	}
	return base64.URLEncoding.EncodeToString(buf)
}

func ClarifyInt(v string) (o []int, e error) {
	var buf []byte
	buf, e = base64.URLEncoding.DecodeString(v)
	if e != nil {
		return
	}
	r := bytes.NewReader(buf)
	var n int64
	n, e = binary.ReadVarint(r)
	if e != nil {
		return
	}
	var x int64
	for i := int64(0); i < n; i++ {
		x, e = binary.ReadVarint(r)
		if e != nil {
			return
		}
		o = append(o, int(x))
	}
	return
}

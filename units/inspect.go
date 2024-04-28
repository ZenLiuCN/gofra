package units

import (
	"os"
	"path/filepath"
)

func ExecutableName() string {
	return filepath.Base(os.Args[0])
}

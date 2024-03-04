package reader

import (
	"fmt"
	"io"
)

type rsFunction func(io.Reader, io.Writer) error

var rsFunctions = map[string]rsFunction{
	"R1":    R1,
	"R2":    R2,
	"R2Bis": R2Bis,
	"R3":    R3,
}

func Read(version string, in io.Reader, out io.Writer) error {
	rsFunc, ok := rsFunctions[version]
	if !ok {
		return fmt.Errorf("%s invalid function selection Please choose from R1 to R7", version)
	}
	return rsFunc(in, out)
}

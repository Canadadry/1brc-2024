package reader

import (
	"bufio"
	"io"
)

func R3Bis(in io.Reader, out io.Writer) error {
	return R3(bufio.NewReaderSize(in, 1000000), out)
}

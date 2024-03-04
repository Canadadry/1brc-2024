package reader

import (
	"bufio"
	"io"
)

func R2Bis(in io.Reader, out io.Writer) error {
	return R2(bufio.NewReaderSize(in, 1000000), out)
}

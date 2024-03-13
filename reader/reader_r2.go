package reader

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func readOneByte(r io.Reader) byte {
	buf, ok := r.(*bufio.Reader)
	if ok {
		b, err := buf.ReadByte()
		if err != nil {
			return 0
		}
		return b
	}
	b := make([]byte, 1, 1)
	_, err := r.Read(b)
	if err != nil {
		return 0
	}
	return b[0]
}

func readString(r io.Reader) (string, byte) {
	buf := [512]byte{}
	i := 0
	for ; i < len(buf); i++ {
		b := readOneByte(r)
		if b == ';' {
			return string(buf[:i]), b
		}
		if b == 0 || b == '\n' {
			return "", b
		}
		buf[i] = b
	}
	return "", 0
}

func readFloat(r io.Reader) (float64, byte) {
	val := 0
	decimal := 0

	sign := 1.0
	b := readOneByte(r)
	if b == 0 {
		return 0.0, 0
	}
	if b == '\n' {
		return 0.0, '\n'
	}
	if b == '-' {
		sign = -1.0
	} else {
		val = int(b - '0')
	}
	for _ = range 100 {
		b := readOneByte(r)
		if b == '.' {
			decimal = 1
			continue
		}
		if b == 0 || b == '\n' {
			return sign * float64(val) / float64(decimal), b
		}
		val = val*10 + int(b-'0')
		decimal = decimal * 10
	}
	return sign * float64(val) / float64(decimal), 0
}

func readLine(r io.Reader) (string, float64, byte) {
	str, b := readString(r)
	if b != ';' {
		return "", 0.0, b
	}
	val, b := readFloat(r)
	return str, val, b
}

func R2(in io.Reader, out io.Writer) error {
	type Station struct {
		Min, Sum, Max float64
		Count         int
	}

	data := map[string]*Station{}

	for {
		str, temp, b := readLine(in)

		d, ok := data[str]
		if !ok {
			d = &Station{
				Min:   temp,
				Sum:   temp,
				Max:   temp,
				Count: 1,
			}
			if str != "" {
				data[str] = d
			}
		} else {
			if d.Min > temp {
				d.Min = temp
			}
			if d.Max < temp {
				d.Max = temp
			}
			d.Sum = d.Sum + temp
			d.Count = d.Count + 1
		}
		if b == 0 {
			break
		}
	}

	keys := make([]string, 0, len(data))
	for name := range data {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	fmt.Fprint(out, "{")
	for i, k := range keys {
		d := data[k]
		fmt.Fprintf(out, "%s=%.1f/%.1f/%.1f", k, d.Min, d.Sum/float64(d.Count), d.Max)
		if i < len(keys)-1 {
			fmt.Fprint(out, ", ")
		}
	}
	fmt.Fprint(out, "}\n")

	return nil
}

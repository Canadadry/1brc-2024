package reader

import (
	"fmt"
	"io"
	"sort"
)

func crc32(s []rune) uint32 {
	var crc uint32 = 0xFFFFFFFF

	for i := 0; i < len(s); i++ {
		ch := uint32(s[i])
		for j := 0; j < 8; j++ {
			b := (ch ^ crc) & 1
			crc >>= 1
			if b != 0 {
				crc ^= 0xEDB88320
			}
			ch >>= 1
		}
	}

	return ^crc
}

func R3(in io.Reader, out io.Writer) error {
	type Station struct {
		Name          string
		Min, Sum, Max float64
		Count         int
	}

	data := map[uint32]*Station{}

	for {
		str, temp, b := readLine(in)
		hash := crc32([]rune(str))
		d, ok := data[hash]
		if !ok {
			d = &Station{
				Name:  str,
				Min:   temp,
				Sum:   temp,
				Max:   temp,
				Count: 1,
			}
			if str != "" {
				data[hash] = d
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
	type pair struct {
		hash uint32
		name string
	}
	keys := make([]pair, 0, len(data))
	for h, d := range data {
		keys = append(keys, pair{h, d.Name})
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].name < keys[j].name
	})
	fmt.Fprint(out, "{")
	for i, k := range keys {
		d := data[k.hash]
		fmt.Fprintf(out, "%s=%.1f/%.1f/%.1f", k.name, d.Min, d.Sum/float64(d.Count), d.Max)
		if i < len(keys)-1 {
			fmt.Fprint(out, ", ")
		}
	}
	fmt.Fprint(out, "}\n")

	return nil
}

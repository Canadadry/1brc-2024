package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strconv"
)

type rsFunction func(io.Reader, io.Writer) error

func Read(version string, in io.Reader, out io.Writer) error {
	rsFunctions := map[string]rsFunction{
		"R1": R1,
	}
	rsFunc, ok := rsFunctions[version]
	if !ok {
		return fmt.Errorf("%s invalid function selection Please choose from R1 to R7", version)
	}
	return rsFunc(in, out)
}

func R1(in io.Reader, out io.Writer) error {
	type Station struct {
		Min, Sum, Max float64
		Count         int
	}
	r := csv.NewReader(in)
	r.Comma = ';'

	data := map[string]Station{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if len(record) != 2 {
			continue
		}
		temp, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			continue
		}
		d, ok := data[record[0]]
		if !ok {
			d.Min = temp
			d.Sum = temp
			d.Max = temp
			d.Count = 1
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
		data[record[0]] = d
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

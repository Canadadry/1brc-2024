package reader

import (
	"testing"
)

func TestCrc32(t *testing.T) {
	cases := []struct {
		in   string
		want uint32
	}{
		{
			in:   "Hello, World!",
			want: 0xEC4AC3D0,
		},
		{
			in:   "",
			want: 0x00000000,
		},
		{
			in:   "123456789",
			want: 0xCBF43926,
		},
		{
			in:   "The quick brown fox jumps over the lazy dog",
			want: 0x414FA339,
		},
	}

	for _, c := range cases {
		got := crc32([]rune(c.in))
		if got != c.want {
			t.Errorf("crc32(%q) == %X, want %X", string(c.in), got, c.want)
		}
	}
}

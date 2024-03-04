package reader

import ()

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

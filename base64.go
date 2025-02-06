package base64zero

import (
	"strings"
)

var base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Encode(source []byte) []byte {
	var result []byte

	for i := 0; i < len(source); i += 3 {
		block := uint32(source[i]) << 16

		for j := 0; j < 2; j++ {
			if i+j+1 >= len(source) {
				continue
			}

			block |= uint32(source[i+j+1]) << (8 - 8*j)
		}

		for j := 0; j < 4; j++ {
			var charByte byte = '='
			if i+j <= len(source) {
				index := (block >> uint(18-6*j)) & 0x3F
				charByte = base64Table[index]
			}

			result = append(result, charByte)
		}
	}

	return result
}

func Decode(source []byte) []byte {
	var result []byte

	for i := 0; i < len(source); i += 4 {
		if source[i] == '=' {
			break
		}

		block := uint32(0)
		for j := 0; j < 4; j++ {
			if source[i+j] != '=' {
				index := strings.Index(base64Table, string(source[i+j]))
				if index == -1 {
					break
				}

				block |= uint32(index) << (18 - 6*j)
			}
		}

		for j := 0; j < 3; j++ {
			charByte := byte((block >> uint(16-8*j)) & 0xFF)
			if charByte == 0 {
				break
			}

			result = append(result, charByte)
		}
	}

	return result
}

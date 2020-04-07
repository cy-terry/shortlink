package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

type URLConverter interface {
	ToShortURL(langURL string) []string
}

type HashURLConverter struct{}

func (hc *HashURLConverter) ToShortURL(langURL string) []string {
	data := md5.Sum([]byte(langURL))
	md5code := hex.EncodeToString(data[:])
	shortArray := make([]string, 4)

	for i := 0; i < 4; i++ {
		var bytes []byte
		hexInt, _ := strconv.ParseInt(md5code[i<<3:(i+1)<<3], 16, 64)
		hexInt = hexInt & int64(0x3FFFFFFF)
		for i := 0; i < 6; i++ {
			bytes = append(bytes, chars[hexInt&0x0000003D])
			hexInt = hexInt >> 5
		}
		shortArray[i] = string(bytes)
	}
	return shortArray
}

var chars = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3',
	'4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D',
	'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
	'Y', 'Z',
}

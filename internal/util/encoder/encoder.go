package encoder

import (
	"strings"
)

// alphabet TODO shake me
const (
	alphabet  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	abLength  = uint64(len(alphabet))
	urlLength = 10
)

// EncodeUrl converting an id uint64 from decimal to 63base
func EncodeUrl(id uint64) string {
	res := make([]byte, 10)
	for i := urlLength - 1; i >= 0; i-- {
		if id > 0 {
			res[i] = alphabet[id%abLength]
			id /= abLength
		} else {
			res[i] = alphabet[0]
		}
	}

	return string(res)
}

// DecryptUrl converting a shortened url from 63base to decimal
func DecryptUrl(shortUrl string) uint64 {
	byteArr := []byte(shortUrl)
	res := uint64(0)
	pow := uint64(1)

	for i := 0; i < urlLength; i++ {
		index := strings.Index(alphabet, string(byteArr[urlLength-1-i]))
		res += uint64(index) * pow
		pow *= abLength
	}

	return res
}

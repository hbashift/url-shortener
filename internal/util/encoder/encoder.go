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

func DecryptUrl(shortUrl string) uint64 {
	byteArr := []byte(shortUrl)
	var res uint64 = 0

	for i := 0; i < urlLength; i++ {
		index := strings.Index(alphabet, string(byteArr[urlLength-1-i]))
		res += uint64(index) * uint64Pow(abLength, uint64(i))
	}

	return res
}

// uint64Pow TODO optimize me
func uint64Pow(n, m uint64) uint64 {
	if m == 0 {
		return 1
	}

	result := n

	var i uint64 = 2
	for ; i <= m; i++ {
		result *= n
	}

	return result
}

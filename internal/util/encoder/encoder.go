package encoder

import (
	"math/rand"
)

var (
	alphabet  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	abLength  = uint64(len(alphabet))
	urlLength = 10
)

/*
EncodeUrl converting an id uint64 from decimal to 63base
shuffles alphabet when shuffle flag is true
*/
func EncodeUrl(id uint64, shuffle bool) string {
	if shuffle {
		alphabetArr := []byte(alphabet)
		rand.Shuffle(len(alphabetArr), func(i, j int) {
			alphabetArr[i], alphabetArr[j] = alphabetArr[j], alphabetArr[i]
		})
		alphabet = string(alphabetArr)
	}

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

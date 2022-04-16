package utils

import "crypto/sha512"

func HashIP(ip string) string {
	h := sha512.New384()
	h.Write([]byte(ip))
	return FunnyEncoding(h.Sum(nil))
}

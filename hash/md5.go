package hash

import "crypto/md5"

var hasher = md5.New()

func md5Sum(b []byte) [16]byte {
	return md5.Sum(b)
}

func hasherSum(b []byte) []byte {
	hasher.Write(b)
	return hasher.Sum(nil)
}

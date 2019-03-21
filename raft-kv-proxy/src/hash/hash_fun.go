package hash

import "hash/crc32"

// key拿到hash值

func Crc32IEEE(key []byte, mod uint32) uint32 {
	return crc32.ChecksumIEEE(key) % mod
}
package hash

import (
	"fmt"
	"hash/crc32"
	"testing"
)

func TestHashCRC32(t *testing.T) {

	data := []byte("hello turing")
	fmt.Printf("%d", crc32.ChecksumIEEE(data) % 1024)

}
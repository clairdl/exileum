package main

import (
	"fmt"
	"os"

	"github.com/IncSW/go-bencode"
)

// strings are immutable in golang, whereas byte arrays ([]byte) are not,
// and since we only need to read the filepath, it should be typed as a string...
// however, the bencode library i chose wants a byte array, so until i write my own parser, i'll have to convert!
func Decode(fp string) (*interface{}, error) {
	f, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	data, err := bencode.Unmarshal(f)
	if err != nil {
		return nil, err
	}
	fmt.Println("hi world", data)
	return &data, nil
}

package main

import (
	"fmt"
	"log"
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
	logger("logs.txt", data)

	return &data, nil
}

func logger(outpath string, dataToLog interface{}) {
	out, err := os.OpenFile(outpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer out.Close()

	byteKey := []byte(fmt.Sprintf("%v", dataToLog.(interface{})))
	out.Write(byteKey)
	log.SetOutput(out)
}

package main

// https://golangcode.com/url-encode-value/

type torrentDecoded struct {
	name     string // filename
	announce string // tracker
	pieces   string // number of pieces to download
	pieceLen int    // size of each piece except last
}

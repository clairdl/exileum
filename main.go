package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"

	bencode "github.com/anacrolix/torrent/bencode"
)

type bdecodedInfo struct {
	Name        string `bencode:"name"`
	Length      int    `bencode:"length"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

type bdecodedFile struct {
	Announce     string       `bencode:"announce"`
	PublisherUrl string       `bencode:"publisher-url"`
	Info         bdecodedInfo `bencode:"info"`
	InfoHash     [20]byte
	PeerId       [20]byte
}

// USAGE: exileum [path to torrent file] [output path]
func main() {

	// Setup
	args := os.Args[1:]
	var torStruct bdecodedFile
	var pts = &torStruct

	// rawTorrentfile := make(map[string]interface{}) //t

	if !isInputValid() {
		log.Fatal("Input invalid, check if input and output filepaths are valid")
	}

	// build torrent filepath
	torrentFilePath := getFilePath(args[0])

	// read torrent file into memory
	openedFilepath, err := os.ReadFile(torrentFilePath)
	ErrCheck(err)

	// unmarshal into rawTorrentfile empty interface basically
	// if err := bencode.Unmarshal(openedFilepath, &rawTorrentfile); err != nil {
	// 	panic(err)
	// }

	// unmarshal into main struct
	if err1 := bencode.Unmarshal(openedFilepath, &torStruct); err1 != nil {
		panic(err1)
	}
	pts.computeInfoHash()
	pts.computePeerId()
	fmt.Println("\n\n INFOHASH \n\n", torStruct.InfoHash)

	prettyprint, _ := json.MarshalIndent(torStruct, "", "    ")
	fmt.Println(string(prettyprint))

	// fmt.Println(hex.EncodeToString(torStruct.InfoHash[:]))

	pts.getPeerlistRequestURL()
}

func isInputValid() bool {
	if len(os.Args) < 2 {
		return false
	} else {
		return true
	}
}

func getFilePath(file string) (filepath string) {
	_, fullWorkingDir, _, _ := runtime.Caller(0)
	workingDir, _ := path.Split(fullWorkingDir)
	return workingDir + file
}

// adds the InfoHash key to the bdecodedFile struct
func (i *bdecodedFile) computeInfoHash() {
	data, err := bencode.Marshal(i.Info)
	if err != nil {
		panic(err) // panic v fatal
	}
	i.InfoHash = sha1.Sum(data)
}

func (i *bdecodedFile) computePeerId() {
	r := make([]byte, 1)
	rand.Read(r)
	i.PeerId = sha1.Sum(r)
}

// returns the tracker URL with params included
func (i *bdecodedFile) getPeerlistRequestURL() {

	base, err := url.Parse(i.Announce)
	fmt.Println(base)
	ErrCheck(err)

	v := url.Values{}
	v.Set("info_hash", string(i.InfoHash[:]))
	v.Set("peer_id", string(i.PeerId[:]))

	base.RawQuery = v.Encode()
	fmt.Println(base.String())

	fmt.Println(base.RawQuery)

	req, _ := http.NewRequest("GET", base.String(), nil)
	res, err := http.DefaultClient.Do(req)
	ErrCheck(err)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))

	res.Body.Close()

	// url := "http://bt3.t-ru.org/ann?info_hash=ef8c5b770da5119f584eacb61c024251fc112d03&left=5854908&peer_id=fdcb4c5d2e905e1526c373e34663204d65d0777e"

	// http://bt3.t-ru.org/ann
	// ?info_hash=ef8c5b770da5119f584eacb61c024251fc112d03
	// &peer_id=fdcb4c5d2e905e1526c373e34663204d65d0777e
}

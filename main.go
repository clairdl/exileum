/*
1. skeleton core functionality
2. write tests (TDD paradigm)
3. implement and iterate

general flow:
- user runs binary and passes path to a .torrent file (url and magnet support is next)

- start go routine to parse .torrent file, send data back through channel incrementally
	blocking ops/pieces of data:

		- `annouce`: required to get peer list from tracker
		- `info`: required for validating downloaded pieces

- when i get `announce` from channel, send GET to tracker for peer list, the response should be:
		{
			`interval`: number of seconds until we should GET the updated peer list
			`peers`: byte array, list of ip/port addresses, chunked into 6 bytes
		}
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// USAGE: exileum  [output path]

	args := os.Args[1:]
	// todo: support magnet links and hosted torrent files

	fmt.Printf("Downloading:\n %s into ——> %s\n", args[0], args[1])

	f, err := Decode(args[0])
	if err != nil {
		// log.fatal's in main, because other errors should bubble up to here
		log.Fatal(err)
	}
	fmt.Println(f)
}

/*
bad:
main ––> Client.init (now removed) ––> Decode ––> Client

good:
main ––> Decode ––> Client

why?
Decode is an arbitrary operation as far as Client is concerned because it precedes Client's instantiation; thus it should not be causally coupled and or managed by Client itself
*/

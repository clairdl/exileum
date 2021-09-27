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
	"os"
)

/*
 */
func main() {
	// USAGE: exileum  [output path]

	args := os.Args[1:]
	// todo: support magnet links and hosted torrent files

	fmt.Printf("Downloading:\n %s into ——> %s\n", args[0], args[1])
}
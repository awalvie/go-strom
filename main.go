// command parsing and execution

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	// usage message
	about = `go-strom is a cli bittorrent client written in Go by awalvie.

Usage: go-strom <torrent | magnet> <destination> [-h]
Example: go-strom -t ~/torrent.file -d ~/Downloads

Options:
`

	// more info message when the program is called incorrectly
	moreInfo = "Try 'go-strom --help' for more information.\n"
)

func _error(err string) {
	fmt.Fprintf(flag.CommandLine.Output(), "ERROR: "+err+"\n"+moreInfo)
	os.Exit(1)
}

func main() {
	// options for the cli
	torrentPtr := flag.String("t", "", "Path to torrent file.")
	magnetPtr := flag.String("m", "", "Magnet hash.")
	locationPtr := flag.String("d", "./", "Download location for the torrent.")

	// usage message that is rendered when -h|--help is used
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), about)
		flag.PrintDefaults()
	}

	// parse our flags
	flag.Parse()

	// both flags can't be empty
	if *torrentPtr == "" && *magnetPtr == "" {
		_error("A torrent file or magnet link is required.")
	}

	// one option needs to be chosen out of torrent or magnet
	if *torrentPtr != "" && *magnetPtr != "" {
		_error("Choose either torrent or magnet, not both.")
	}

	// download torrent using torrent file
	if *torrentPtr != "" {
		log.Printf("Downloading using torrent file")
		downloadTorrent(*torrentPtr, *locationPtr)
	}

	// download torrent using magnet infohash
	if *magnetPtr != "" {
		downloadMagnet(*magnetPtr)
	}
}

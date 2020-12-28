// command parsing and execution

package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	// usage message
	about = `go-strom is a cli bittorrent client written in Go by awalvie.

Usage: go-strom <torrent | magnet> <destination> [-h]

Options:
`

	// more info message when the program is called incorrectly
	moreInfo = `Try 'go-strom --help' for more information.`
)

func err(err string) {
	fmt.Fprintln(flag.CommandLine.Output(), "ERROR:", err)
	fmt.Fprintln(flag.CommandLine.Output(), moreInfo)
}

func main() {
	// options for the cli
	torrentPtr := flag.String("t", "", "Path to torrent file.")
	magnetPtr := flag.String("m", "", "Magnet hash.")
	// locationPtr := flag.String("d", "./", "Download location for the torrent.")

	// usage message that is rendered when -h|--help is used
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), about)
		flag.PrintDefaults()
	}

	// parse our flags
	flag.Parse()

	// both flags can't be empty
	if *torrentPtr == "" && *magnetPtr == "" {
		err("A torrent file or magnet link are required.")
		os.Exit(1)
	}

	// one option needs to be chosen out of torrent or magnet
	if *torrentPtr != "" && *magnetPtr != "" {
		err("Choose either torrent or magnet, not both.")
		os.Exit(1)
	}
}

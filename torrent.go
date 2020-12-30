package main

import (
	"fmt"
	"log"
)

func downloadTorrent(torrentFile string, location string) {
	log.Println("Checking if options provided are correct")

	// Check if torrent file exists
	if result, _ := exists(torrentFile); result == false {
		_error("Torrent file does not exist.")
	}

	// Check if location exists if it's different than ./
	if location != "./" {
		if result, _ := exists(location); result == false {
			_error("Download location provided does not exist.")

		}
	}

	torrent, err := parseTorrent(torrentFile)
	if err != nil {
		_error("Problem parsing torrent file, ensure your torrent file is not corrupt")
	}

	fmt.Println(torrent)
}

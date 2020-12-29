package main

import (
	"fmt"
	"os"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func downloadTorrent(torrentFile string, location string) {
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

	fmt.Println("Downloading torrent...")
	parseTorrent(torrentFile)
}

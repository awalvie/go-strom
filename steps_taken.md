# Steps Taken

This is how I approached building the project.

#### Figure out a reasonable directory structure, if there exists one.

Keep it simple, start from a `main.go` and build on top of it.

#### What the hell is `GOPATH` and why does it matter?

Don't need to care about it, we can use `go mod init` to make our repository into a module and it werks.

#### How to take input?

Use the `flag` library.

#### What's going to be the input?

It can either be a magnet hash or a torrent file.

#### How to use the program.
You do,

`go_strom <magnet | torrent> <destination> [-h]`

#### Pick one to work on, torrent or magnet

Torrent.

#### Input a torrent file

Take input as a torrent file.

#### Parse torrent file

Write a parser for the torrent file and stick the data into a structure.

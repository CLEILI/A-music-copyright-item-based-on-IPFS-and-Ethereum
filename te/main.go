package main

import (
	"fmt"
	"item/ipfs"
)

func main() {
	cid := "QmSi9JbT3fn4rsgPgg3Xf4LjGQXs6WSGRt5biwum1usyFd"
	dpath := "./music.mp3"
	path := ipfs.Download(cid, dpath)
	fmt.Println(path)
}

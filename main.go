package main

import (
	"Deepslate/nbt"
	"compress/gzip"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("H:\\minecraft\\1.17.1\\world\\level.dat")

	if err != nil {
		log.Panic(err)
	}

	reader, err := gzip.NewReader(f)
	if err != nil {
		log.Panic(err)
	}

	tag, err := nbt.Read(reader)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(tag)
}
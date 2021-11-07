package main

import (
	"Deepslate/anvil"
	"os"
)

func main() {
	f, err := os.Open("H:\\Test\\world\\region\\r.0.0.mca")
	if err != nil {
		panic(err)
	}
	storage, err := anvil.NewStorage(f)
	if err != nil {
		panic(err)
	}

	println(storage)

	chunk, err := storage.GetChunk(0, 0)
	if err != nil {
		panic(err)
	}
	println(chunk)
}
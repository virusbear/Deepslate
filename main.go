package main

import (
	"Deepslate/nbt"
	"log"
	"os"
)

type test struct {
	Name string `nbt:"name"`
}

func main() {
	f, err := os.Open("H:\\minecraft\\1.17.1\\world\\region\\r.0.0.mca")

	if err != nil {
		log.Panic(err)
	}

	nbt.Unmarshal(f)
}
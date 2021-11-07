package anvil

import (
	"Deepslate/chunk"
	"errors"
	"fmt"
	"io"
	"unsafe"
)

var (
	InvalidHeader = errors.New("invalid anvil header")
)
type ReadWriteSeekerCloser interface {
	io.ReadWriteSeeker
	io.Closer
}

//Storage takes ownership of storage
type Storage struct {
	storage ReadWriteSeekerCloser
	header Header
}

type Location uint32
func (loc Location) Get() (uint32, uint8) {
	return uint32(loc >> 1 & 0xFFFFFF), uint8(loc)
}

type Timestamp uint32

type Header struct {
	locations [1024]Location
	timestamps [1024]Timestamp
}

func NewStorage(storage ReadWriteSeekerCloser) (*Storage, error) {
	sto := Storage{
		storage: storage,
	}

	err := sto.loadHeader()

	return &sto, err
}

func (sto *Storage) GetChunk(x int, z int) (*chunk.Chunk, error) {
	idx := ((x % 32) + (z % 32)) * 32
	if idx > len(sto.header.locations) {
		return nil, fmt.Errorf("chunk (%d, %d) not contained in this region file", x, z)
	}

	offset, sectorCount := sto.header.locations[idx].Get()

	if _, err := sto.storage.Seek(int64(offset), 0); err != nil {
		return nil, err
	}

	return chunk.Read(sto.storage, int(sectorCount))
}

func (sto *Storage) loadHeader() error {
	buf := make([]byte, 8192)
	if _, err := sto.storage.Seek(0, 0); err != nil {
		return err
	}

	if _, err := sto.storage.Read(buf); err != nil {
		return err
	}

	sto.header = *(*Header)(unsafe.Pointer(&buf))

	return nil
}


func (sto *Storage) Close() error {
	err := sto.storage.Close()
	if err != nil {
		return fmt.Errorf("an error occured trying to close Storage: %v", err)
	}

	return nil
}
package chunk

import (
	"errors"
	"io"
)

var (
	SectorSize = 4096
)

type Chunk struct {

}

func Read(reader io.Reader, sectorCount int) (*Chunk, error) {
	buf := make([]byte, SectorSize * sectorCount)
	if _, err := reader.Read(buf); err != nil {
		return nil, err
	}

	return nil, errors.New("not implemented")
}

package log

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	encoding = binary.BigEndian
)

const (
	lenWidth = 8
)

type store struct {
	*os.File
	mu sync.Mutex
	buf *bufio.Writer
	size uint64
}

func NewStore(f *os.File) (*store, error) {
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}

	size := uint64(fi.Size())

	return &store {
		File: f,
		size: size,
		buf: bufio.NewWriter(f),
	}, nil
}

func (s *store) Append(p []byte) (uint64, uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	position := s.size
	if err := binary.Write(s.buf, encoding, uint64(len(p))); err != nil {
		return 0,0,err
	}

	w, err := s.buf.Write(p)
	if err != nil {
		return 0, 0, err
	}

	w += lenWidth
	s.size += uint64(w)
	return uint64(w), position, nil
}

func (s *store) Read(position uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	//Flush the Writer Buffer, in case the record has not been flushed yet
	if err := s.buf.Flush(); err != nil {
		return nil, err
	}

	//Figure out the size of bytes we need to get the record
	size := make([]byte, lenWidth)
	if _, err := s.File.ReadAt(size, int64(position)); err != nil {
		return nil, err
	}

	//Fetch the record
	b := make([]byte, encoding.Uint64(size))
	if _, err := s.File.ReadAt(b, int64(position+lenWidth)); err != nil {
		return nil, err
	}

	return b, nil
}

func (s *store) ReadAt(p []byte, offset int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
		return 0, err
	}
	return s.File.ReadAt(p, offset)
}

func (s *store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.buf.Flush()
	if err != nil {
		return err
	}

	return s.File.Close()
}
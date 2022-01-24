package server

import (
	"fmt"
	"sync"
)

var ErrorOffsetNotFound = fmt.Errorf("Offest not found")

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	//Use the length of the records as index
	//Then append it the end of slice
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	//Check if offset exists, if not return a new Record
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrorOffsetNotFound
	}
	return c.records[offset], nil
}

type Log struct {
	mu sync.Mutex
	records []Record
}

type Record struct {
	Value []byte
	Offset uint64
}
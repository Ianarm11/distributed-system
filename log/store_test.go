package log

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

var (
	write = []byte("HELLLLLLLLLLLOOOOOO LASSSS VEGAS!")
	width = uint64(len(write)) + lenWidth
)

func TestStoreAppendRead(t *testing.T) {
	f, err := ioutil.TempFile("", "store_append_read_test")
	require.NoError(t, err)

	defer os.Remove(f.Name())

	s, err := NewStore(f)
	require.NoError(t, err)

	//testAppend(t, s)
	//testRead(t, s)
	//testReadAt(t, s)

	s, err = NewStore(f)
	require.NoError(t, err)
	//testRead(t, s)
}
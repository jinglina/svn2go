package svn

//#include <svn_io.h>
import "C"
import (
	"io"
	"unsafe"
)

// TODO svn stream supports seek, maybe we can use it
type SvnStream struct {
	stream *C.svn_stream_t
}

// Read bytes from svn stream
func (s *SvnStream) Read(dest []byte) (n int, err error) {
	c := C.apr_size_t(len(dest))

	if err := C.svn_stream_read(s.stream, (*C.char)(unsafe.Pointer(&dest[0])), &c); err != nil {
		return int(c), makeError(err)
	}

	if c == 0 {
		return 0, io.EOF
	} else {
		return int(c), nil
	}
}

// Closes svn stream
func (s *SvnStream) Close() error {
	if err := C.svn_stream_close(s.stream); err != nil {
		return makeError(err)
	}

	return nil
}
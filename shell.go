package interactive

import "io"

// shell implements the ReadWriter interface.
type shell struct {
	io.Reader
	io.Writer
}

func (s *shell) read(data []byte) (n int, err error) {
	return s.Read(data)
}

func (s *shell) write(data []byte) (n int, err error) {
	return s.Write(data)
}

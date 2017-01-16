package interactive

import "io"

// Shell implements the ReadWriter interface.
type Shell struct {
	r io.Reader
	w io.Writer
}

func (s *Shell) Read(data []byte) (n int, err error) {
	return s.r.Read(data)
}
func (s *Shell) Write(data []byte) (n int, err error) {
	return s.w.Write(data)
}

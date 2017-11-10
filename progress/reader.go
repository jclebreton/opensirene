package progress

import "io"

// It's proxy reader, implement io.Reader
type Reader struct {
	io.Reader
	Progress *Progress
}

func NewProgressReader(rc io.Reader, name, step string, total uint64) *Reader {
	reader := &Reader{
		Reader: rc,
		Progress: &Progress{
			Name:         name,
			Step:         step,
			Total:        total,
			ProgressChan: DefaultChan,
		},
	}
	return reader
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Progress.Add(int64(n))
	return
}

// Close the reader when it implements io.Closer
func (r *Reader) Close() (err error) {
	if closer, ok := r.Reader.(io.Closer); ok {
		return closer.Close()
	}
	return
}

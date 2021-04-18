package httpx

// Dummy is a dummy that implement io.ReadCloser interface to mock response body.
type Dummy struct {
	ReadErr  error
	CloseErr error
}

// NewDummy returns a dummy http.
func NewDummy(err error) *Dummy {
	return &Dummy{
		ReadErr:  err,
		CloseErr: err,
	}
}

// Read return response based on ReadErr.
func (d *Dummy) Read([]byte) (n int, err error) {
	if d.ReadErr != nil {
		return 0, d.ReadErr
	}

	return 1, nil
}

// Close return response based on CloseErr.
func (d *Dummy) Close() error {
	return d.CloseErr
}

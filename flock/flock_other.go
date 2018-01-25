package flock

// Acquire is not implemented on non-unix systems.
func Acquire(path string) error {
	return nil
}

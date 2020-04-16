package std

import (
	"bytes"
	"io"
	"os"
	"sync"
)

// Bind binds given writers to standard output and error streams
func Bind(out, err io.Writer, f func()) {
	// Replacing os.Std*
	prevOut := os.Stdout
	prevErr := os.Stderr
	readOut, writeOut, _ := os.Pipe()
	readErr, writeErr, _ := os.Pipe()
	os.Stdout = writeOut
	os.Stderr = writeErr
	wg := sync.WaitGroup{}
	wg.Add(2)
	wg2 := sync.WaitGroup{}
	wg2.Add(2)
	go func() {
		wg.Done()
		_, _ = io.Copy(out, readOut)
		wg2.Done()
	}()
	go func() {
		wg.Done()
		_, _ = io.Copy(err, readErr)
		wg2.Done()
	}()

	// Waiting for goroutines to start
	wg.Wait()

	// Invoking callback
	f()

	// Rollback
	os.Stdout = prevOut
	os.Stderr = prevErr

	// Sync
	_ = readOut.Sync()
	_ = readErr.Sync()
	_ = writeOut.Sync()
	_ = writeErr.Sync()

	// Closing
	_ = readOut.Close()
	_ = readErr.Close()
	_ = writeOut.Close()
	_ = writeErr.Close()
	wg2.Wait()
}

// BindStrings replaces stdout and stderr while callback func is invoked
// Produced output is returned as strings
func BindStrings(f func()) (sout, serr string) {
	out := bytes.NewBuffer(nil)
	err := bytes.NewBuffer(nil)

	Bind(out, err, f)
	sout = string(out.Bytes())
	serr = string(err.Bytes())
	return
}

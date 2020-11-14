package paasio

import (
	"io"
	"sync"
)

type counter struct {
	count int64
	ops   int
	lock  sync.Mutex
}

func (c *counter) add(amount int64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count += amount
	c.ops++
}

func (c *counter) amount() (count int64, ops int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	count, ops = c.count, c.ops
	return c.count, c.ops
}

type readCounter struct {
	r io.Reader
	counter
}

func (r *readCounter) Read(p []byte) (i int, err error) {
	i, err = r.r.Read(p)
	if err == nil {
		r.add(int64(i))
	}
	return i, err
}

func (r *readCounter) ReadCount() (int64, int) {
	return r.amount()
}

type writeCounter struct {
	w io.Writer
	counter
}

func (w *writeCounter) Write(p []byte) (i int, err error) {
	i, err = w.w.Write(p)
	if err == nil {
		w.add(int64(i))
	}
	return i, err
}

func (w *writeCounter) WriteCount() (int64, int) {
	return w.amount()
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		r: reader,
	}
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
		w: writer,
	}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		NewReadCounter(rw),
		NewWriteCounter(rw),
	}
}

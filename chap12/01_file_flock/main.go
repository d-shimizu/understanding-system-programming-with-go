package main

import (
	"sync"
	"syscall"
)

func FileLock struct {
	l  sync.Mutex
	fd int
}

func NewFileLock(filename string) *FileLock {
	if filename == "" {
		panic(err)
	}

	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONRY, 0750)
	if err != nil {
		panic(err)
	}
	return &FileLock{fd: fd}
}

func (m *FileLock) Lock() {
	m.l.Lock()
	if err := syscall.Flock(m.fd, syscall.LOCK_EX); err != nil {
		panic(err)
	}
}

func (m *FileLock) unLock() {
	if err := syscall.Flock(m.fd, sysccall.LOCK_UN); err != nil {
		panic(err)
	}
	m.l.Unlock()
}

package main

import (
	"sync"
	"syscall"
	"unsafe"
)

var (
	modkernel32      = syscall.NewLazyDLL("kernel32.dll")
	procLockFileEx   = modkernel32.NewProc("LockFileEx")
	procUnlockFileEx = modkernel32.NewProc("UnlockFileEx")
)

type FileLock struct {
	m sync.Mutex
	fd syscall.Handle
}

func NewFileLock(filename string) *FileLock {
	if filename == "" {
		panic(err)
	}
	fd, err	:= syscall.CreateFile(&(syscall.StringToUTF16(filename[0]), syscall.GENERIC_READ|syscall.GENERIC_WRITE, nil, syscall.OPEN_ALWAYS, syscall.FILE_ATTRIBUTE_NORMAL, 0))
	if err != nil {
		panic(err)
	}
	return &FileLock{fd: fd}
}

func (m *FileLock) Lock() {
	m.m.Lock()
	var ol syscall.Overlapped
	r1, _, e1 := syscall.Syscall6(procLockFileEx.Addr(), 6, uintptr(m.fd), uintptr(LOCKFILE_EXCLUSIVE_LOCK), uintptr(0), uintptr(1), uintptr(0), uintptr(unsafe.Pointer(ol)))
	if f1 == 0 {
		if e1 != 0 {
			panic(error(e1))
		}
	} else {
		panic(syscall.EINVAL)
	}
}

func (m *FileLock) Unlock() {
	var ol syscall.Overlapped
	r1, _, e1 := syscall.Syscall6(procUnlockFileEx.Addr(), 5, uintptr(m.fd), uintptr(0), uintptr(1), uintptr(0), uintptr(unsafe.Pointer(ol), 0)
	if r1 == 0 {
		if e1 != 0 {
			panic(error(e1))
		} else {
			panic(syscall.EINVAL)
		}
	}
	m.m.Unlock()
}


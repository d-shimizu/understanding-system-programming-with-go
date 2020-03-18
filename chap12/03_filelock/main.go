package main

import (
	"fmt"
	"time"
    "sync"
    "syscall"
)

// 構造体
type FileLock struct {
        l sync.Mutex
        fd int
}

// Lockの獲得処理
func NewFileLock(filename string) *FileLock {
        if filename == "" {
                panic("filename needed")
        }
		// ファイルを開く関数 func syscall.Open(path string, mode int, perm uint32)
        fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONLY, 0750)
        if err != nil {
                panic(err)
        }
		// ポインタで返す
        return &FileLock{fd: fd}
}

// 構造体の関数
func (m *FileLock) Lock() {
        m.l.Lock()
        if err := syscall.Flock(m.fd, syscall.LOCK_EX); err != nil {
                panic(err)
        }
}

func (m *FileLock) Unlock() {
        if err := syscall.Flock(m.fd, syscall.LOCK_UN); err != nil {
                panic(err)
        }
        m.l.Unlock()
}

func main() {
	l := NewFileLock("main.go")
	fmt.Println("try locking...")

	l.Lock()
	fmt.Println("locked!!")

	time.Sleep(10 * time.Second)

	l.Unlock()
	fmt.Println("unlock")
}

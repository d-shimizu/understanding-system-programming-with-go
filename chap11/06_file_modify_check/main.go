package main

import (
	"log"
	"gopkg.in/fsnotify.v1"
	// https://godoc.org/gopkg.in/fsnotify.v1#Op.String
)

func main() {
	counter := 0

	// 監視用のインスタンスを作成する
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func(){
		for {

			select { // ≒ switch 
			// watcher.Events 変更イベントが入るチャネル
			// ブロッキングしてイベントを待つ
			case event := <-watcher.Events:
				log.Println("events:", event)
				if event.Op & fsnotify.Create == fsnotify.Create {
					log.Println("event.Op:", event.Op) // debug
					log.Println("fsnotify.Create:", fsnotify.Create) // debug
					log.Println("create file:", event.Name)
					counter++
				} else if event.Op & fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					counter++
				} else if event.Op & fsnotify.Remove == fsnotify.Remove {
					log.Println("renamed file:", event.Name)
					counter++
				} else if event.Op & fsnotify.Rename == fsnotify.Rename {
					log.Println("chmod file:", event.Name)
					counter++
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}

			if counter > 3 {
				done<-true
			}
		}
	}()

	// 監視対象フォルダを指定する
	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}
	// 受信のみを検知する
	// log.Println("goroutine:", <-done)
	<-done
}


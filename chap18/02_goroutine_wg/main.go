package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 計算用の関数: 元金均等
func calc(id, price int, interestRate float64, year int) {
	months := year * 12
	interest := 0
	for i := 0; i < months; i++ {
		balance := price * (months - i) / months
		interest += int(float64(balance) * interestRate / 12)
	}
	fmt.Println("year=%d total=%d interest=%d id=%d\n", year, price + interest, interest, id)
}

// goroutineで呼び出されるworker
func worker(id, price int, interestRate float64, years chan int, wg *sync.WaitGroup) {
	// タスクがなくなってタスクのチャネルがcloseされるまで無限ループ
	for year := range years {
		calc(id, price, interestRate, year)
		wg.Done()
	}
}

func main() {
	// 借入額
	price := 40000000
	// 利子 1.1%固定金利
	interestRate := 0.011

	// タスクをchannelに格納する
	// バッファ数35のチャネルを生成する
	years := make(chan int, 35)
	for i := 1; i < 36; i++ {
		// channelへデータ(i)を送信する
		years <- i
	}

	// Wait Groupを宣言する
	var wg sync.WaitGroup

	wg.Add(35)

	// CPUコア数分のgoroutine起動
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, price, interestRate, years, &wg)
	}
	// 全てのワーカーが終了する
	close(years)
	wg.Wait()
}

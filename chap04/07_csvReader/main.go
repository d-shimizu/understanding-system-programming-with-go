package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `13101,"100  ","1000003","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾄﾂﾊﾞｼ(1ﾁｮｳﾒ)","東京都","千代田区","一ツ橋（１丁目）",1,0,1,0,0,0
13101,"101  ","1010003","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾄﾂﾊﾞｼ(2ﾁｮｳﾒ)","東京都","千代田区","一ツ橋（２丁目）",1,0,1,0,0,0
13101,"100  ","1000012","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾋﾞﾔｺｳｴﾝ","東京都","千代田区","日比谷公園",0,0,0,0,0,0
13101,"102  ","1020093","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾋﾗｶﾜﾁｮｳ","東京都","千代田区","平河町",0,0,1,0,0,0
13101,"102  ","1020071","ﾄｳｷｮｳﾄ","ﾁﾖﾀﾞｸ","ﾌｼﾞﾐ","東京都","千代田区","富士見",0,0,1,0,0,0
`

func main() {
	// 終端まで読み込み
	reader := strings.NewReader(csvSource)
	// debug
	// fmt.Println(reader)
	// csvで読み込み
	csvReader := csv.NewReader(reader)
	// debug
	fmt.Println(csvReader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}
}


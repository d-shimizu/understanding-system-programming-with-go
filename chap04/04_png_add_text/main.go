package main

import (
        "bytes"
        "encoding/binary"
        "io"
        "os"
		"hash/crc32"
)

func readChunks(file *os.File) []io.Reader {
	// チャンクを格納する配列を定義する
	var chunks []io.Reader

	// 最初の8バイトを探す
	file.Seek(8, 0)
	var offset int64 = 0

	for {
		var length int32 //  = 0
		// binaryファイルをint32(32ビットの整数)バイト列分の長さを確保してビッグエンディアンに変換して読み込む
		// err := binary.Read(file, binary.BigEndian, &length)
		err := binary.Read(file, binary.BigEndian, 12)
		// EOFだったら終了
		if err == io.EOF {
			break
		}
		// offsetの位置からint64(length)+12のデータを読み出す
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		// 次のチャンクの先頭に移動する
		// 現在位置はlengthを読み終わった箇所なので、
		// チャンク名(4バイト) + データ長 + CRC(4バイト)先に移動する
		offset, _ = file.Seek(int64(length + 8), 1)
	}
	return chunks
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)

	// CRCを計算して追加する
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())

	return &buffer
}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create("Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)

	// シグネチャ書き込み
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// 先頭に必要なIHDRチャンクを書き込み
	io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))
	// 残りのチャンクを追加
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}
}

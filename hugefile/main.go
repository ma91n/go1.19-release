package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("c03.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	indexMap := make(map[string]int64)
	r := csv.NewReader(f)

	var prefectureCD string
	for {
		offset := r.InputOffset()
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if record[0] != prefectureCD {
			prefectureCD = record[0]
			indexMap[record[1]] = offset
		}

		if err != nil {
			log.Fatal(err) // 何かしらのエラーハンドリング
		}
	}

	// オフセット位置を表示
	fmt.Println("北海道", indexMap["北海道"])
	fmt.Println("神奈川県", indexMap["東京都"])
	fmt.Println("沖縄県", indexMap["沖縄県"])

	// 該当の最初の1行を取得
	line, err := fetchFirstLine(indexMap, "東京都")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", line)
}

func fetchFirstLine(indexMap map[string]int64, key string) ([]string, error) {
	f, err := os.Open("c03.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = f.Seek(indexMap[key], 0) //特定の位置から読む
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	return r.Read()
}

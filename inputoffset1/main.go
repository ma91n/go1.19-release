package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	s := `aaa,"b

bb",ccc
ddd,ee"e,fff
zzz,yyy,xxx
`
	r := csv.NewReader(strings.NewReader(s))
	for {
		fmt.Printf("input offset:%d: ", r.InputOffset())

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err) // 何かしらのエラーハンドリング
		}

		fmt.Printf("%#v\n", record)
	}

}

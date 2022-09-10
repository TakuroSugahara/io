package read

import (
	"fmt"
	"log"
	"os"
)

func ReadSample() {
	// os.FileオブジェクトをOpen関数か何かで得てるものとする
	// os.FileはREAD専用のfile descriptorだったり、I/Oできるものだったりする

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("get wd error", err.Error())
	}

	// file読み込みのためのfdのオブジェクトを返却
	file, err := os.Open(fmt.Sprintf("%s/text.txt", wd))
	if err != nil {
		log.Fatalln("get wd error", err.Error())
	}
	// 読み込みだけだからerror握りつぶしてもいい
	defer file.Close()

	data := make([]byte, 1024)
	// file読み込みを実行して読み込めた分のbyte数を返却する
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("fail to read file", err)
	}

	fmt.Printf("read %d bytes: \n", count)
	// 確保したbyteに対して返却されたbyte数分を出力する
	fmt.Println(string(data[:count]))
}

package write

import (
	"fmt"
	"os"
)

func Write() {
	file, _ := os.Create("write.txt")
	// TODO: 書き込みなどファイルに影響を与えるので本来はちゃんとエラーハンドリングするべき
	defer file.Close()

	str := "write this file by Golang!"
	data := []byte(str)
	count, _ := file.Write(data)

	fmt.Printf("write %d bytes\n", count)
}

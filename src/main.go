package main

import (
	"fmt"
	"io"
	"net/http"
)

// 爬取top250
func main() {
	url := "https://movie.douban.com/top250"
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	buf := make([]byte, 4096)
	var result string
	for {
		n, err := response.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp body err", err)
			return
		}
		result += string(buf[:n])
		//打印读取的网页
		fmt.Println(result)
	}
}

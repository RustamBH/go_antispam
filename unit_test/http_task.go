// http_task.go
package main

import (
	"fmt"
)

// MaskURL маскирует часть URL после протокола
func MaskURL(url []byte) []byte {
	if len(url) <= 7 || string(url[0:7]) != "http://" {
		return string(url)
	}

	rs := []rune(string(url))
	for i := 7; i < len(url); i++ {
		rs[i] = '*'
	}
	url = []byte(string(rs))
	return url
}

func main() {
	str := "http://hehefouls.netHAHAHA"
	//str := "hTTp://youth-elixir.com"
	url := []byte(str)
	masked := MaskURL(url)
	fmt.Println(string(masked))
}


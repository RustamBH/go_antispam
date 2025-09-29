// http_task.go
package main

import (
	"fmt"
)

// MaskURL маскирует часть URL после протокола
func MaskURL(url []byte) string {
	if string(url[0:7]) != "http://" {
		return string(url)
	}

	for i := 7; i < len(url); i++ {
		url[i] = '*'
	}
	return string(url)
}

func main() {
	str := "http://hehefouls.netHAHAHA"
	//str := "hTTp://youth-elixir.com"
	url := []byte(str)
	masked := MaskURL(url)
	fmt.Println(masked)
}






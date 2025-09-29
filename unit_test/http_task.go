// http_task.go
package main

import "fmt"

// MaskURL маскирует часть URL после протокола
func MaskURL(url string) string {
	if len(url) <= 7 {
		return url
	}

	out := []rune(url)
	for i := 7; i < len(out); i++ {
		out[i] = '*'
	}
	return string(out)
}

func main() {
	url := "http://hehefouls.netHAHAHA"
	masked := MaskURL(url)
	fmt.Println(masked)
}

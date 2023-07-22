package utils

import (
	"bufio"
	"log"
	"os"
)

func CheckError(err error, msg string) {
	if err != nil {
		log.Panic(msg)
	}
}

func GetInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inp := scanner.Text()
	return inp
}
func StarCount(msg []byte) int {
	count := 0
	for _, v := range msg {
		if string(v) == "*" {
			count++
		}
	}
	return count
}

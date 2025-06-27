package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

// exercise 2
// go run grepfiles.go content txtfiles/one.txt txtfiles/two.txt txtfiles/three.txt
func main() {
	slog.Info("Chapter 2 Exercise 2")

	// exercise 2
	target := os.Args[1]
	args := os.Args[2:]
	for i := 0; i < len(args); i++ {
		go grepfiles(args[i], target)
	}
	time.Sleep(2 * time.Second)
}

func grepfiles(filename, target string) {
	fmt.Printf("Work %s started at %s\n", filename, time.Now().Format("15:04:05"))
	data, err := os.ReadFile(filename)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if strings.Contains(string(data), target) {
		msg := fmt.Sprintf("file %s has %s", filename, target)
		slog.Info(msg)
	}
	fmt.Printf("Work %s finished at %s\n", filename, time.Now().Format("15:04:05"))
}

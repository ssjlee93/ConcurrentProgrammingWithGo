package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

// exercise 3
// go run grepdir.go content ./txtfiles
func main() {
	slog.Info("Chapter 2 Exercise 3")

	// exercise 2
	target := os.Args[1]
	args := os.Args[2]
	go grepdir(args, target)
	time.Sleep(2 * time.Second)
}

func grepdir(dir, target string) {
	fmt.Printf("Work %s started at %s\n", dir, time.Now().Format("15:04:05"))
	files, err := os.ReadDir(dir)
	for _, entry := range files {
		fmt.Println(entry)
		go scanfiles(dir+"/"+entry.Name(), target)
	}
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Printf("Work %s finished at %s\n", dir, time.Now().Format("15:04:05"))
}

func scanfiles(filename, target string) {
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

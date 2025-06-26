package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

func main() {
	slog.Info("Chapter 2 Exercises")

	// exercise 1
	//args := os.Args[1:]
	//for i := 0; i < len(args); i++ {
	//	go catfiles(args[i])
	//}
	//time.Sleep(2 * time.Second)

	// exercise 2
	target := os.Args[1]
	args := os.Args[2:]
	for i := 0; i < len(args); i++ {
		go grepfiles(args[i], target)
	}
	time.Sleep(2 * time.Second)
}

// exercise 1
// go run exercise.go one.txt two.txt three.txt
func catfiles(filename string) {
	fmt.Printf("Work %s started at %s\n", filename, time.Now().Format("15:04:05"))
	data, err := os.ReadFile(filename)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(string(data))
	fmt.Printf("Work %s finished at %s\n", filename, time.Now().Format("15:04:05"))
}

// exercise 2
// go run exercise.go content one.txt two.txt three.txt
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

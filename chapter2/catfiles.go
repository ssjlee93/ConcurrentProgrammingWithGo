package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

// exercise 1
// go run catfiles.go txtfiles/one.txt txtfiles/two.txt txtfiles/three.txt
func main() {
	slog.Info("Chapter 2 Exercise 1")

	// exercise 1
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		go catfiles(args[i])
	}
	time.Sleep(2 * time.Second)
}

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

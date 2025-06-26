package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

func main() {
	slog.Info("Exercise 2")
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		go one(args[i])
	}
	time.Sleep(2 * time.Second)
}

func one(filename string) {
	fmt.Printf("Work %s started at %s\n", filename, time.Now().Format("15:04:05"))
	data, err := os.ReadFile(filename)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(string(data))
	fmt.Printf("Work %s finished at %s\n", filename, time.Now().Format("15:04:05"))
}

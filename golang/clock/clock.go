package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func clock(stop chan bool) {
	secondTicker := time.NewTicker(time.Duration(time.Second))
	for {
		select {
		case <-secondTicker.C:
			fmt.Print("\033[H\033[2J\033[32m")
			fmt.Printf("%v\033[0m\033[?25l", time.Now().Format("Mon 15:04:05, Jan 2 OTC"))
		case <-stop:
			fmt.Println("\nByebye\033[?25h")
			return
		}
	}
}

func main() {
	fmt.Println("# 1764Snippets")
	fmt.Println("## Golang Clock")

	stop := make(chan bool)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n')
		stop <- true
	}()
	clock(stop)
}

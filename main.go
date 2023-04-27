package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Message struct {
	Content string
}

func main() {
	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	// Create a channel of Message structs
	messageQueue := make(chan Message)

	// Start a goroutine to receive messages from the queue
	go func() {
		defer wg.Done()
		for message := range messageQueue {
			fmt.Printf("Received message: %s\n", message.Content)
		}
	}()

	// Start a goroutine to enter messages into the queue
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Enter message: \n")
			content, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			content = strings.TrimSpace(content)
			if content == "" {
				fmt.Println("Empty message")
				continue
			}

			message := Message{Content: content}
			fmt.Printf("Adding message: %s", message.Content)
			messageQueue <- message
		}
	}()

	// Wait for goroutines to complete
	wg.Wait()
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Message represents a user message
type Message struct {
	ID      int
	Sender  string
	Content string
}

// worker processes messages from the queue
func worker(id int, messageQueue <-chan Message, processedMessages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range messageQueue {
		fmt.Printf("📩 Worker %d processing message from %s: %s\n", id, msg.Sender, msg.Content)
		time.Sleep(time.Duration(rand.Intn(2)+1) * time.Second) // Simulate processing time
		processedMessages <- fmt.Sprintf("✅ Message %d processed by Worker %d", msg.ID, id)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // To generate random delays

	const numWorkers = 3
	messageQueue := make(chan Message, 10)     // Buffered channel for incoming messages
	processedMessages := make(chan string, 10) // Buffered channel for processed messages
	var wg sync.WaitGroup

	// Create Workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, messageQueue, processedMessages, &wg)
	}

	// Simulating user messages
	go func() {
		users := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
		for i := 1; i <= 10; i++ {
			user := users[rand.Intn(len(users))]
			message := Message{ID: i, Sender: user, Content: fmt.Sprintf("Hello, I'm %s!", user)}
			fmt.Printf("📤 New message received from %s\n", user)
			messageQueue <- message
			time.Sleep(time.Millisecond * 500) // Simulate message arrival
		}
		close(messageQueue) // Close the channel after all messages are sent
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(processedMessages) // Close processed messages channel after all workers finish
	}()

	// Display processed messages
	for msg := range processedMessages {
		fmt.Println(msg)
	}
}

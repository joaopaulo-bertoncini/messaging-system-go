## 📌 Practical: Concurrent Messaging System in Go
Scenario:

Create a concurrent messaging server that receives user messages and processes each one in parallel using workers.

- Users send messages to an input channel (messageQueue).\
- A pool of workers reads and processes the messages.\
- The processed messages are sent to an output channel (processedMessages).\
- The main goroutine prints the processed messages.\

### How Does It Work?

1 - Receives user messages in a buffered channel.\
2 - Workers process messages concurrently.\
3 - The program waits for all workers to finish before exiting.\
4 - Processed messages are printed to the console.\

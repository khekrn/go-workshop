# **Golang Concurrency and Channels**

- Understand why Go is designed for concurrency and its advantages.
- Learn communication between Goroutines using Channels.

#### **1. Difference between Concurrency and Parallelism**

- **Concurrency:** The ability to handle multiple tasks at once by switching between them efficiently.
- **Parallelism:** Executing multiple tasks simultaneously on multiple processors.
- Go supports **Concurrency** as a core design principle, while **Parallelism** depends on the hardware.

#### **2. Goroutines: The Building Blocks**

- **Goroutines** are lightweight threads managed by the Go runtime.
- They enable functions to run independently and concurrently.

  ```go
  package main

  import (
      "fmt"
      "time"
  )

  func printNumbers() {
      for i := 1; i <= 5; i++ {
          fmt.Println(i)
          time.Sleep(500 * time.Millisecond)
      }
  }

  func main() {
      go printNumbers() // Launching Goroutine
      fmt.Println("Main function continues execution")
      time.Sleep(3 * time.Second) // Give Goroutine time to finish
  }
  ```

#### **3. Basic Syntax and Usage of Goroutines**

- Use the `go` keyword followed by a function call.
- Anonymous Goroutines can also be created:
  ```go
  go func() {
      fmt.Println("Anonymous Goroutine")
  }()
  ```

#### **4. What are Channels?**

- **Channels** provide a way for Goroutines to communicate and synchronize.
- They allow safe data exchange between Goroutines.

  ```go
  package main

  import "fmt"

  func main() {
      ch := make(chan string) // Create a channel
      go func() {
          ch <- "Hello from Goroutine"
      }()

      msg := <-ch // Receive data from channel
      fmt.Println(msg)
  }
  ```

#### **5. Unbuffered vs Buffered Channels**

- **Unbuffered Channels:** Block until data is received.
- **Buffered Channels:** Have capacity; sender doesn't block until the buffer is full.
  ```go
  ch := make(chan int, 2) // Buffered channel with capacity 2
  ch <- 1
  ch <- 2
  fmt.Println(<-ch)
  fmt.Println(<-ch)
  ```

#### **6. Closing Channels and Best Practices**

- Use `close(channel)` to close a channel.
- Avoid writing to closed channels.

  ```go
  ch := make(chan int)
  go func() {
      defer close(ch)
      ch <- 42
  }()

  val, ok := <-ch
  fmt.Println(val, ok) // ok will be false if channel is closed
  ```

#### **7. Select Statement for Multiplexing**

- The `select` statement in Go is used for **concurrent communication** when working with **channels**. It’s similar to a `switch` statement but specifically designed for handling **multiple channel operations simultaneously**.
- It **waits** for one of its multiple channel operations to become **ready**.
- If multiple channels are ready, it **selects one randomly**.
- If no channels are ready, it **blocks** until one becomes ready (unless there’s a `default` case).

#### 📦 **Basic Syntax of `select`:**

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
default:
    fmt.Println("No messages ready, moving on...")
}
```

- `case msg := <-ch1`: Executes when there’s a value available from `ch1`.
- `default`: Executes if no channel is ready (optional case).

#### 🚀 **Why is `select` Essential for Reading from Channels?**

1. **Handles Multiple Channels Efficiently:**

   - Instead of blocking on a single channel, `select` allows listening on multiple channels **simultaneously**.

2. **Avoid Deadlocks:**

   - Without `select`, if you try reading from an unready channel, the program could **deadlock**.

3. **Non-blocking Operations with `default`:**

   - Using `default`, you can perform non-blocking reads or writes on channels.

4. **Graceful Shutdowns:**

   - Allows monitoring a **signal channel** (like `os.Signal`) alongside data channels for smooth application termination.

5. **Fairness in Selection:**
   - When multiple channels are ready, `select` chooses **randomly**, preventing starvation of any channel.

#### ⚖️ **Comparison: Blocking Read vs. `select` Read**

#### ❌ **Blocking Read (Single Channel Read):**

```go
msg := <-ch1
fmt.Println(msg)
```

- This blocks indefinitely if `ch1` has no message.
- Cannot handle multiple channels.

#### ✅ **Using `select`:**

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
}
```

- The program listens to **both channels simultaneously**.
- Executes as soon as any one channel has data.

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine for Channel 1
	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch1 <- "Message from Channel 1"
		}
	}()

	// Goroutine for Channel 2
	go func() {
		for {
			time.Sleep(2 * time.Second)
			ch2 <- "Message from Channel 2"
		}
	}()

	// Handle Ctrl+C for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Continuous select without explicit loop
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-stop:
			fmt.Println("Received interrupt signal. Exiting program.")
			return
		}
	}
}
```

- Both goroutines continuously send messages into their respective channels.
- Signal Handling (`os.Signal`). A `stop` channel listens for `os.Interrupt` (`Ctrl+C`) and `syscall.SIGTERM`.
- - The `select` listens on:
    - `ch1` for messages from Channel 1.
    - `ch2` for messages from Channel 2.
    - `stop` for interrupt signals.
- When `Ctrl+C` is pressed, the program prints a shutdown message and exits cleanly.

---

### **8. sync.WaitGroup**

- `sync.WaitGroup` is used to wait for a collection of Goroutines to complete.
- It helps ensure that the main Goroutine waits until all spawned Goroutines finish execution.
- Key Methods:
  - `Add(int)` - Increments the counter.
  - `Done()` - Decrements the counter.
  - `Wait()` - Blocks until the counter becomes zero.

**Example with sync.WaitGroup:**

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers completed")
}
```

### **9. sync.Mutex and sync.RWMutex**

- In concurrent programming, multiple goroutines may attempt to modify shared data simultaneously, leading to race conditions.
- **sync.Mutex:** ensures exclusive access to shared resources.
- **sync.RWMutex:** Allows multiple readers but only one writer.

```go
package main

import (
    "fmt"
    "sync"
)

type SafeCounter struct {
    mu sync.Mutex
    value int
}

func (sc *SafeCounter) Increment() {
    sc.mu.Lock()
    sc.value++
    sc.mu.Unlock()
}

func main() {
    sc := SafeCounter{}
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            sc.Increment()
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Final Value:", sc.value)
}
```

---

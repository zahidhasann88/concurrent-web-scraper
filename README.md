
```markdown
# Concurrent Web Scraper

This is a simple yet powerful demonstration of Go's concurrency model using goroutines and channels. It is designed to fetch data from multiple URLs concurrently, showcasing the advantages of using Go for high-performance, scalable applications.

## Features

- **Concurrency with Goroutines**: Utilizes Go's lightweight goroutines to perform multiple tasks simultaneously, making it ideal for handling many concurrent operations.
- **Channel Communication**: Uses Go's channels to safely communicate between goroutines, avoiding shared memory issues and race conditions.
- **Error Handling**: Implements proper error handling to ensure that the application fails gracefully and provides meaningful error messages.
- **Modular Design**: The code is organized into packages (`scraper`) and is easy to extend and maintain.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher installed on your machine.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/concurrent-web-scraper.git
    cd concurrent-web-scraper
    ```

2. Install the dependencies:

    ```bash
    go mod tidy
    ```

### Running the Application

To run the project:

```bash
go run main.go
```

This will start the application and fetch data from a list of URLs concurrently.

### Sample Output

```plaintext
Fetched https://golang.org in 0.37 seconds (12434 bytes)
Fetched https://godoc.org in 0.41 seconds (8973 bytes)
Fetched https://gophercises.com in 0.32 seconds (4523 bytes)
```

## Project Files

- `main.go`: The main entry point of the application. It initializes the URLs, starts the concurrent fetching process, and handles the results.
- `scraper/scraper.go`: Contains the logic for fetching URLs concurrently and sending results through channels.
- `scraper/types.go`: Defines the `Result` struct used to encapsulate the details of each URL fetch operation.
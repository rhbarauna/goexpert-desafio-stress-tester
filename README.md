## Stress Tester

This is a tool developed in Go designed for conducting stress tests on a specific web application.

### Detailed Explanation

The **Stress Tester** is designed to enable users to assess the performance and robustness of a web application by simulating a significant traffic load.
It is implemented in Go and leverages concurrency to send multiple requests simultaneously.

The command-line interface of the **Stress Tester** is configured through three main flags:

- `--url` (mandatory): Sets the URL of the application to be tested. It is the only mandatory flag and indicates where the requests will be sent.

- `--requests` (optional): Specifies the total number of requests the stress test will perform. If not provided, the default is 1.

- `--concurrency` (optional): Defines the concurrency level, i.e., how many requests will be executed simultaneously. If not provided, the default is 1.

### Usage Example

bash

```go
go run main.go --url https://example.com --requests 1000 --concurrency 50
```

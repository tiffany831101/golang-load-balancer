# Simple Go Load Balancer with Round Robin and Proxy

A basic Go load balancer using round-robin for even request distribution among servers. Includes a proxy server to showcase request forwarding within the load balancer group.

## How It Works

- **Round Robin Load Balancing:** Distributes incoming requests evenly among servers in a cyclic fashion.
  
- **Proxy Server for Forwarding:** Demonstrates forwarding of requests within the load balancer group using a reverse proxy.

## Usage

1. Start the load balancer:

    ```bash
    go run main.go
    ```

2. Access the load balancer at `http://localhost:8080`.

3. Observe round-robin distribution of requests among servers.

## Configuration

Customize servers and their addresses in the `main` function. The load balancer evenly distributes requests among specified servers.

```go
servers := []Server{
    newSimpleServer("http://server1.example.com/"),
    newSimpleServer("http://server2.example.com/"),
    // Add more servers as needed
}
```

![image](https://github.com/tiffany831101/golang-load-balancer/assets/39373272/7ef94907-5b13-42e9-9b19-723b1c495cbb)

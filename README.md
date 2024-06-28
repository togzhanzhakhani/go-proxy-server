# Go Proxy Server

## Introduction

This project is an HTTP proxy server built in Go. It receives requests from clients, forwards them to specified external services, and returns the response in JSON format.

## Installation

1. Clone the repository:

```sh
git clone https://github.com/togzhanzhakhani/go-proxy-server.git
cd go-proxy-server
```

2. Build the project:

```sh
make build
```

3. Run the server:

```sh
make run
```

## Deployed on RENDER:

### Base URL
https://go-proxy-server.onrender.com

### Endpoint
URL: /proxy
Method: POST
Content-Type: application/json

### Request Body:
```sh
{
  "method": "GET",           // The HTTP method (GET, POST, etc.) to use for the proxied request.
  "url": "http://example.com", // The URL of the third-party service to which the request will be forwarded.
  "headers": {                // (Optional) Any headers to include in the proxied request.
    "Authorization": "Bearer token",
    "User-Agent": "YourCustomUserAgent"
  }
}
```
### Response:
```sh
{
  "id": "requestId",         // A unique identifier for the request.
  "status": 200,             // The HTTP status code returned by the third-party service.
  "headers": {               // The headers returned by the third-party service.
    "Content-Type": "application/json",
    "Content-Length": "1256"
  },
  "length": 1256             // The length of the response body from the third-party service.
}
```
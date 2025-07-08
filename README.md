# Go URL Shortener

This is a simple project written in Go that implements a **minimal URL shortener**, similar to bit.ly or tinyurl.

## Features

- Generates random short IDs for submitted URLs
- Stores the mapping between the short ID and the original URL in memory
- Automatically redirects visitors to the original URL when accessing the short link
- Simple JSON API for creating short URLs

## How to Use

### 1. Start the Application

Make sure you have Go installed (Go 1.18+ recommended).

Clone the repository and run the application:

```bash
go run main.go
```

The server will listen on port 8080.

### 2. Create a Short URL

Send an HTTP POST request to /shorten with a JSON payload containing the original URL.

Example request:

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.example.com"}'
```

Example response:

```bas
{
  "short_url": "http://localhost:8080/Ab3xYz"
}
```

### 3. Access the Short URL

Open the short URL in your browser or use curl. The server will redirect you to the original URL.

## Project Structure

- `main.go`: The complete application code
  - `URLStore`: The in-memory storage for URL mappings
  - `generateShortID`: The function generating random short IDs
  - Two HTTP handlers:
    - `/shorten` for creating short URLs
    - `/` for redirection

## Limitations

- URLs are stored only in memory and will be lost when the server restarts.
- No extensive URL validation.
- No expiration mechanism for links.
- No persistence to files or databases.

## Possible Improvements

You can extend this project with:

- Persistence in Redis or a local file
- TTL (time-to-live) for expiring links
- Deterministic or user-defined short IDs
- HTML pages for creating short URLs
- Access statistics

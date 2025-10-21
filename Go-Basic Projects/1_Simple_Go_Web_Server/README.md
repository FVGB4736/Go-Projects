# Simple Go Web Server

A basic web server implemented in Go that demonstrates handling static files, form submissions, and basic routing.

## Features

- Static file serving
- Form handling with POST requests
- Basic routing with multiple endpoints
- Simple HTML pages

## Project Structure

```
├── main.go          # Main server implementation
├── static/          # Static files directory
│   ├── index.html   # Homepage
│   └── form.html    # Form page
```

## Endpoints

- `/` - Serves static files
- `/hello` - Returns a welcome message (GET only)
- `/form` - Handles form submissions (POST)

## Getting Started

### Prerequisites

- Go 1.x installed on your system

### Running the Server

1. Clone this repository
2. Navigate to the project directory
3. Run the server:
```bash
go run main.go
```
4. The server will start at `http://localhost:8080`

## Usage

- Visit `http://localhost:8080` to see the static homepage
- Visit `http://localhost:8080/hello` for a welcome message
- Visit `http://localhost:8080/form.html` to access the form page

## API Reference

### GET /hello
Returns a welcome message. Only accepts GET requests.

### POST /form
Handles form submissions with name and address fields.

Request body parameters:
- `name`: String
- `address`: String

By FVGB4736
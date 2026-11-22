# Go E-Commerce Backend API

A simple, lightweight E-Commerce REST API built using Go's standard library (`net/http`, `encoding/json`). This project serves as the foundation for a backend store service, currently handling product data retrieval from local JSON storage.

## Features
- **RESTful Endpoints:** Fetch product catalogs via HTTP.
- **JSON Data Handling:** Structured parsing and serialization using Go structs.
- **Clean Architecture:** Organized into modular packages (`controller`, `data`).

## Project Structure
```text
├── controller/
│   └── product_controller.go   # HTTP handlers and JSON parsing logic
├── data/
│   └── products.json           # Mock product catalog
└── main.go                     # Application entry point & routing
```

## Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/Ali-Haidar-159/e-commerce-golang.git
   cd e-commerce-golang
   ```

2. **Run the server**
   ```bash
   go run main.go
   ```

## Roadmap
- [x] Read and serve all products (`GET /products`)
- [ ] Get single product by ID (`GET /products/{id}`)
- [ ] Add, update, and delete products (`POST`, `PUT`, `DELETE`)
- [ ] Migrate from JSON file storage to PostgreSQL/MySQL
- [ ] Add user authentication & shopping cart endpoints

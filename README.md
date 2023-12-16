# Project Title: Cautious Memory

## Introduction
Cautious Memory is a Go-based web application designed to manage profiles and their associated fields and subfields. It provides a RESTful API to perform CRUD operations on these entities, leveraging Golang's robustness and efficiency.

## Features
- CRUD operations for profiles.
- Management of associated fields and subfields.
- RESTful API endpoints.
- PostgreSQL database integration.
- Use of GORM for ORM and Gin for HTTP routing.

## Getting Started

### Prerequisites
- Go (Version 1.16 or later)
- PostgreSQL
- Set up environment variables for database connection:
  - `POSTGRES_USER`
  - `POSTGRES_PASSWORD`
  - `POSTGRES_DB`
  - `POSTGRES_HOSTNAME`
  - `POSTGRES_PORT`

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/joe5saia/cautious-memory.git
   ```
2. Navigate to the project directory:
   ```sh
   cd cautious-memory
   ```
3. Install Go dependencies:
   ```sh
   go mod tidy
   ```

### Running the Application
Run the application using the following command:
```sh
go run cmd/main.go
```
The server will start, typically on `http://localhost:8080`.

## Swagger Documentation for API Endpoints
Swagger documentation is available at `http://localhost:8080/swagger/index.html`.

To generate the Swagger documentation, run the following command:
```sh
swag fmt
swag init -g cmd/main.go --parseInternal  --parseDependency 
```

Swagger documentation is generated using comments in the code.
For more details see the [Swaggo Documentation](https://github.com/swaggo/swag).


## Contributing
Contributions to Cautious Memory are welcome! Please follow these steps:
1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a pull request.

## License
Distributed under the MIT License. See `LICENSE` for more information.

## Contact
Your Name - [Joe Saia](mailto:joe5saia@gmail.com)
Project Link: https://github.com/joe5saia/cautious-memory
```

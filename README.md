# Job Board API

![Job Board API Banner](https://via.placeholder.com/1200x300.png?text=Job+Board+API+%7C+Golang+%2B+MongoDB+%2B+JWT)

[![Go](https://img.shields.io/badge/Golang-v1.24-blue?logo=go)](https://go.dev/)  
[![Gin](https://img.shields.io/badge/Gin-Gonic-green?logo=go)](https://github.com/gin-gonic/gin)  
[![MongoDB](https://img.shields.io/badge/Database-MongoDB-green?logo=mongodb)](https://www.mongodb.com/)  
[![JWT](https://img.shields.io/badge/Auth-JWT-red?logo=jsonwebtokens)](https://jwt.io/)  
[![Swagger](https://img.shields.io/badge/API-Docs-yellow?logo=swagger)](https://swagger.io/)

A **RESTful API** for a *Job Board* platform developed in **Golang** with the **Gin** framework, **JWT**-based authentication, and **MongoDB** persistence. API endpoint documentation is available through **Swagger UI**.

This project contains only the **API (backend)**, while the **frontend** is located in a separate repository.

---

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Authentication](#authentication)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [License](#license)

---

## Features
- CRUD operations for users, companies, and job postings
- JWT-based authentication and authorization
- Input validation for sensitive data
- Swagger UI for API endpoint exploration
- Clean Architecture principles and modular separation
- RESTful API design patterns

---

## Installation
```bash
# Clone repository
git clone https://github.com/bonjourrog/jb-back.git

cd jb-back
```

---

## Configuration
Create a `.env` file in the root directory with the following variables:

```env
MONGODB_URI=mongodb://localhost:27017
PORT=:8080
DATABASE=jb
SigningKey=<secret_key>
ALLOWED_ORIGINS=<local_origins>
APP_ENV=local
```

---

## Running the Application
```bash
# Download dependencies
go mod tidy

# Run server
go run main.go
```

The server will start on `http://localhost:8080`

---

## Authentication
- **JWT (JSON Web Tokens)** based authentication
- Main flows:
  - User registration
  - User login (JWT token generation)
  - Protected endpoint access with `Authorization: Bearer <token>`

---

## Project Structure
```bash
job-back/
├── config/         # Configuration and environment variables
├── controller/     # Request handlers and business logic
├── db/             # Dabadase configuration
├── docs/           # Swagger documentation
├── entity/         # MongoDB models and schemas
├── middleware/     # Authentication and validation middleware
├── repository      #  Data access layer: interfaces and implementations to interact with the database
├── routes/         # API endpoint definitions
├── service/        # Supporting services
└── main.go         # Application entry point
```

---

## Technologies Used
- [Golang](https://go.dev/) (main backend language)
- [Gin Gonic](https://gin-gonic.com/) (web framework)
- [MongoDB](https://mongodb.com/) (NoSQL database)
- [Swagger](https://swagger.io/) (API documentation)
- [JWT](https://jwt.io/) (secure authentication)

---

## Future Improvements
- Integration with external services (LinkedIn, APIs)
- Unit and integration testing with Go testing framework
- CI/CD pipeline with GitHub Actions
- Docker containerization

---

## License
This project is licensed under the **MIT License**.

---

## Contact
**Developer**: Rogelio Beltran
**Email**: rbv.rogelio.beltran@gmail.com  
**LinkedIn**: https://www.linkedin.com/in/rogeliobeltran/
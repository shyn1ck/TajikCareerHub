# TajikCareerHub

![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Required-brightgreen)
![License](https://img.shields.io/badge/license-MIT-green)

TajikCareerHub is a job listing website for Tajikistan, developed using Go with the Gin framework and PostgreSQL as the database. This project is designed for managing job vacancies and users.

## Table of Contents
- [Functionality](#functionality)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API](#api)
- [Technologies](#technologies)
- [License](#license)

## Functionality

### User Management
- **User Registration**: Create a new user with a username, email, and password.
- **Get Users**: Retrieve a list of all registered users.
- **Get User by ID**: Get information about a user by their unique ID.
- **Update User**: Update user details, including username and password.
- **Delete User**: Soft delete a user (mark as deleted).

### Job Management
- **Create Job**: Add a new job listing with title, description, location, and company.
- **Get Jobs**: Retrieve a list of all available job listings.
- **Get Job by ID**: Get information about a job by its unique ID.
- **Update Job**: Modify job details.
- **Delete Job**: Soft delete a job (mark as deleted).

### Search and Filtering
- **Search Jobs**: Search for jobs by keywords, location, and other criteria.
- **Filter Jobs**: Filter job listings by status, publication date, and other criteria.
- **Sort Jobs**: Sort job listings by publication date, priority, and status.

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/shyn1ck/tajik-career-hub.git
   cd tajik-career-hub
   ```

2. **Install Dependencies**
   Ensure you have Go and PostgreSQL installed. Then, install the project dependencies:
   ```bash
   go mod tidy
   ```

3. **Configuration**
   Create a configuration file `.env` in the root directory of the project and specify your database connection parameters:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=2003
   DB_NAME=tajik_career_hub
   ```

## Running the Application

1. **Start the Application**
   ```bash
   go run main.go
   ```
   The application will be available at [http://localhost:8080](http://localhost:8080).

2. **Database Migrations**
   If there are any database migrations, run them to ensure the database schema is up-to-date:
   ```bash
   go run migrations/migrate.go
   ```

## API

### Users
- **POST /users**: Register a new user.
- **GET /users**: Retrieve a list of users.
- **GET /users/{id}**: Retrieve a user by ID.
- **PUT /users/{id}**: Update a user.
- **DELETE /users/{id}**: Delete a user.

### Jobs
- **POST /jobs**: Create a new job listing.
- **GET /jobs**: Retrieve a list of job listings.
- **GET /jobs/{id}**: Retrieve a job by ID.
- **PUT /jobs/{id}**: Update a job listing.
- **DELETE /jobs/{id}**: Delete a job listing.

### Search and Filtering
- **GET /jobs/search**: Search for jobs.
- **GET /jobs/filter**: Filter job listings.
- **GET /jobs/sort**: Sort job listings.

## Technologies
- **Go**: Programming language for the backend.
- **Gin**: HTTP framework for Go.
- **PostgreSQL**: Database management system.
- **GORM**: ORM for database interactions in Go.
- **JWT**: Authentication and authorization.

## License
This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

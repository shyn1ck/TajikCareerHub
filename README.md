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
- **Update User Password**: Update the password for a user.
- **Check User Exists**: Check if a user exists by their ID.

### Job Management
- **Create Job**: Add a new job listing with title, description, location, and company.
- **Get Jobs**: Retrieve a list of all available job listings.
- **Get Job by ID**: Get information about a job by its unique ID.
- **Update Job**: Modify job details.
- **Delete Job**: Soft delete a job (mark as deleted).
- **Filter Jobs**: Filter job listings by various criteria.
- **Get Jobs by Salary Range**: Retrieve jobs based on a specified salary range.

### Application Management
- **Create Application**: Add a new application for a job.
- **Get Applications**: Retrieve a list of all applications.
- **Get Application by ID**: Get information about an application by its unique ID.
- **Get Applications by User ID**: Retrieve all applications submitted by a specific user.
- **Get Applications by Job ID**: Retrieve all applications submitted for a specific job.
- **Update Application**: Modify application details.
- **Delete Application**: Soft delete an application (mark as deleted).

### Company Management
- **Create Company**: Add a new company.
- **Get Companies**: Retrieve a list of all companies.
- **Get Company by ID**: Get information about a company by its unique ID.
- **Update Company**: Modify company details.
- **Delete Company**: Soft delete a company (mark as deleted).

### Favorite Jobs
- **Add Favorite**: Add a job to a user's list of favorites.
- **Get Favorites by User ID**: Retrieve all favorite jobs for a specific user.
- **Check Favorite Exists**: Check if a job is marked as a favorite for a specific user.
- **Remove Favorite**: Remove a job from a user's list of favorites.

### Job Categories
- **Create Job Category**: Add a new job category.
- **Get Job Categories**: Retrieve a list of all job categories.
- **Get Job Category by ID**: Get information about a job category by its unique ID.
- **Update Job Category**: Modify job category details.
- **Delete Job Category**: Soft delete a job category (mark as deleted).

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
   DB_NAME=tajik_career_hub_db
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

### Authentication
- **POST /auth/sign-up**: Register a new user.
- **POST /auth/sign-in**: Authenticate a user and generate a JWT token.

### Users
- **POST /users**: Register a new user.
- **GET /users**: Retrieve a list of users.
- **GET /users/{id}**: Retrieve a user by ID.
- **PUT /users/{id}**: Update a user.
- **DELETE /users/{id}**: Soft delete a user.
- **PUT /users/{id}/password**: Update user password.
- **GET /users/existence**: Check if a user exists.

### Jobs
- **POST /jobs**: Create a new job listing.
- **GET /jobs**: Retrieve a list of job listings.
- **GET /jobs/{id}**: Retrieve a job by ID.
- **PUT /jobs/{id}**: Update a job listing.
- **DELETE /jobs/{id}**: Soft delete a job listing.
- **GET /jobs/filter**: Filter job listings.
- **GET /jobs/salary-range**: Get jobs by salary range.

### Applications
- **POST /applications**: Create a new application.
- **GET /applications**: Retrieve a list of applications.
- **GET /applications/{id}**: Retrieve an application by ID.
- **GET /applications/user/{userID}**: Retrieve applications by user ID.
- **GET /applications/job/{jobID}**: Retrieve applications by job ID.
- **PUT /applications/{id}**: Update an application.
- **DELETE /applications/{id}**: Soft delete an application.

### Companies
- **POST /companies**: Create a new company.
- **GET /companies**: Retrieve a list of companies.
- **GET /companies/{id}**: Retrieve a company by ID.
- **PUT /companies/{id}**: Update a company.
- **DELETE /companies/{id}**: Soft delete a company.

### Favorites
- **POST /favorites**: Add a job to favorites.
- **GET /favorites/user/{userID}**: Retrieve all favorites for a user.
- **GET /favorites/user/{userID}/job/{jobID}**: Check if a job is a favorite for a user.
- **DELETE /favorites/user/{userID}/job/{jobID}**: Remove a job from favorites.

### Job Categories
- **POST /jobcategories**: Create a new job category.
- **GET /jobcategories**: Retrieve a list of job categories.
- **GET /jobcategories/{id}**: Retrieve a job category by ID.
- **PUT /jobcategories/{id}**: Update a job category.
- **DELETE /jobcategories/{id}**: Soft delete a job category.


## License
This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

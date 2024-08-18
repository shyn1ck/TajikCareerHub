## Technical Specification for TajikCareerHub

**Project Name:** TajikCareerHub

**Description:**
TajikCareerHub is a job listing website tailored for Tajikistan. The website allows users to manage job vacancies and user profiles. It provides functionalities for job and user management, as well as search, filtering, and sorting job listings.

### Functional Requirements:

#### User Management:
- **User Registration:**
    - Create a new user with the following fields:
        - Username (unique)
        - Email (unique)
        - Password
        - Full Name

- **Get Users:**
    - Retrieve a list of all registered users.

- **Get User by ID:**
    - Retrieve information about a user by their unique ID.

- **Update User:**
    - Update user details including username, email, and password.

- **Delete User:**
    - Soft delete a user by marking them as deleted.

#### Job Management:
- **Create Job:**
    - Add a new job listing with:
        - Title
        - Description
        - Location
        - Company

- **Get Jobs:**
    - Retrieve a list of all available job listings.

- **Get Job by ID:**
    - Retrieve information about a job by its unique ID.

- **Update Job:**
    - Modify job details including title, description, location, and company.

- **Delete Job:**
    - Soft delete a job by marking it as deleted.

#### Search and Filtering:
- **Search Jobs:**
    - Search for jobs by keywords, location, and other criteria.

- **Filter Jobs:**
    - Filter job listings by status, publication date, and other criteria.

- **Sort Jobs:**
    - Sort job listings by publication date, priority, and status.

### Technical Requirements:
- **Programming Language:** Go
- **Web Framework:** Gin
- **Database:** PostgreSQL
- **ORM:** GORM
- **Authentication:** JWT for secure user management
- **Data Format:** JSON for API requests and responses

### Additional Features (Optional):
- **Job Alerts:**
    - Notify users of new job postings matching their criteria.

- **User Dashboard:**
    - Provide a dashboard with user statistics and job application tracking.


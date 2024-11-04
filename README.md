# Go API Project Template

I've been reading a lot of best practices articles on Go recently and so I thought I'd put the major ideas I found into one repo. This Go API project was structured following best practices in Go development that I'd come across. It's set up to easily implement basic CRUD operations and demonstrates the flow of a typical Go API, without yet implementing more complex / opinionated features like database migrations, authentication, etc. The project was configured to connect to an existing database for testing out these practices, and kept to a single external dependency (pgx) to keep things simple and maintainable.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before running this project, you need to have the following installed:

- Go (version 1.22+)
- Air for live reloading in development
- A PostgreSQL database with a users table
- A .env file in the root directory with the following environment variables:

```bash
DB_HOST=[probably localhost]
DB_PORT=[probably 5432]
DB_USER=[your username]
DB_PASSWORD=[your password]
DB_NAME=[your database name]
PORT=[your port]
```

### Installing

1. Clone the repository to your local machine:

```bash
git clone https://github.com/Cotter45/go-api-template.git
```

2. Navigate into the project directory:

```bash
cd go-api-template
```

3. To initialize the project, run:

```bash
air init
```

### Development

- To start the development server with live reloading, run:

```bash
make dev
```

- This uses air under the hood to automatically reload your server on file changes.

### Running the Tests

- Run the automated tests with:

```bash
make test
```

- To check the test coverage, use:

```bash
make coverage
```

- This command runs the tests and generates a coverage report.
- Test coverage is sitting at 99% (not 100 because of loadEnv func 🫠), excluding the main.go, and db_pool.go files.

### Building the Project

To build the project, use:

```bash
make build
```

- This generates an executable file named template (or the name you set in your Makefile).

### Running the Application

- To run the built application, use:

```bash
make run
```

### Project Structure

- Here is a brief overview of the project structure:

```
.
├── Makefile
├── api
│ ├── deps_test.go
│ ├── router.go
│ ├── router_test.go
│ ├── server.go
│ └── server_test.go
├── cover.out
├── domain
│ ├── users.go
│ └── users_test.go
├── go.mod
├── go.sum
├── main.go
├── pkg
│ ├── entities
│ │ └── users.go
│ ├── handlers
│ │ ├── deps_test.go
│ │ ├── users.go
│ │ └── users_test.go
│ ├── middleware
│ ├── repositories
│ │ ├── users.go
│ │ └── users_test.go
│ ├── services
│ │ ├── users.go
│ │ └── users_test.go
│ └── utils
│ ├── config.go
│ ├── config_test.go
│ ├── db.go
│ ├── db_pool.go
│ ├── http.go
│ ├── http_test.go
│ ├── logger.go
│ └── logger_test.go
└── main
```

- /api: Contains the router and server setup along with their tests.
- /domain: Sets up dependency injection for the domain packages.
- /pkg: Includes various packages like entities, handlers, repositories, services, and utilities.
  - /entities: Contains the domain entities.
  - /handlers: Contains the HTTP handlers for the API, handling the request and response.
  - /middleware: Contains the middleware for the API (none to show here but you get the idea).
  - /repositories: Contains the database operations for the entities.
  - /services: Contains the business logic for the entities (no real business logic included here in the template though).
  - /utils: Contains the utility functions for the API that are shared across the packages.

### License

This project is licensed under the MIT License - see the LICENSE.md file for details.

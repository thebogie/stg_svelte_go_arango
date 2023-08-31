# GoTodos: GraphQL API with JWT Authentication & Clean Code Architecture

Welcome to GoTodos, a powerful GraphQL API for efficient task management, featuring JWT authentication and a clean code architecture in Golang.

## Features

- **GraphQL API:** Seamlessly manage tasks through a GraphQL interface.
- **JWT Authentication:** Securely authenticate users using JWT tokens.
- **Clean Code Architecture:** Follows a clean and maintainable code structure.
- **Efficient Task Handling:** Easily organize tasks using GraphQL queries and mutations.

## Getting Started

### Prerequisites

- Install Golang: [https://go.dev/dl/](https://go.dev/dl/)
- Install PostgreSQL: [https://www.postgresql.org/download/](https://www.postgresql.org/download/)

### Setting Up

1. Clone this repository to your local machine.
2. Install required packages:
   ```sh
   go mod tidy
3. Set up your PostgreSQL database and configure environment variables:
   ```sh
   export db_user="YOUR_DB_USERNAME"
   export db_pass="USER_PASSWORD"
   export db_name="DATABASE_NAME"
   export db_host="localhost"

### Running the API

1. Start the server:
   ```sh
   go run server.go
2. Access the GraphQL Playground:
   
   Open your browser and navigate to  
   http://localhost:8080.


## Usage

1. Explore the GraphQL Playground to interact with the API.
2. Use GraphQL queries and mutations to manage tasks efficiently.
3. Authenticate securely using JWT tokens.

## Contributing

Contributions are welcome! If you want to enhance features, fix bugs, or improve documentation, please fork this repository, create a new branch, and submit a pull request. For significant changes, kindly open an issue first to discuss the proposed changes.

## License

This project is licensed under the [MIT License](LICENSE).

---

Enhance your task management experience with GoTodos. Explore the power of GraphQL, secure authentication, and clean code architecture in Golang.

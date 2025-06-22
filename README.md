# Monthly Payment Calculator

Monthly Payment Calculator is a simple API built with Go that calculates the monthly payment for a loan based on the provided amount and number of years.

## Features

- **Calculate interest**: Determines the interest based on inputted amount and number of years.
- **Monthly payment calculation**: Calculates the monthly payment for the loan.
- **RESTful API**: Simple endpoints for calculating monthly payments.

## Installation

1. **Clone the repository**
    ```sh
    git clone https://github.com/walonCode/monthly-payment
    cd monthly-payment
    ```

2. **Set up your environment**
    - Ensure you have Go 1.24.4 installed.
    - Use `go mod tidy` to install dependencies.

3. **Run the server**
    ```sh
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

## Configuration and Env

No configuration or environment files are required for this project.

## Usage

### **GET/** - Health Check

**Description**: Confirm the server is running.

**Endpoint**: `GET /`

**Response**:
   ```json
   {
     "ok": true,
     "message": "Welcome to my server"
   }
   ```

### **POST /monthly-payment** - Calculate Monthly Payment

**Description**: Calculate the monthly payment based on the loan amount and number of years.

**Request**:
```json
{
    "amount": 1000,
    "year": 5
}
```

**Response**:
   ```json
   {
     "ok": true,
     "message": "monthly payment",
     "data": {
       "value": 120.68
     }
   }
   ```

## Technologies

- <img src="https://golang.org/doc/gopher/frontpage.png" width="30" height="30" alt="Go">

## Folder Structure

```
.
├── .gitignore               # Git ignore file.
├── go.mod                   # Go module definitions.
├── main.go                  # Main entry file of the app.
└── tmp                      # Directory for temporary use.
```

## Authors

- **walonCode** - [GitHub @walonCode](https://github.com/walonCode)

---

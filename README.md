# Fetch Backend Take-Home Exercise
This repository contains the files for the coding challenge for the Backend Engineer role at Fetch. The app was built according to the requirements listed in this repo: https://github.com/fetch-rewards/receipt-processor-challenge

## Running App
With this being a Go project, assuming you have Go installed on your machine you can clone the repository and run it directly from your terminal:
```
git clone https://github.com/hewitson-j/fetch-backend-challenge.git
cd fetch-backend-challenge
go run .
```
The app will be running on `localhost:8080`. Once you have the app running, you can open Postman/Bruno to test the endpoints. You can also use your terminal to test the endpoints:

- Powershell:
  - ex. `Invoke-RestMethod -Uri "http://localhost:8080/receipts/<receipt_id>/points" -Method Get`
- bash/zsh:
  - ex. `curl "http://localhost:8080/receipts/<receipt_id>/points"`

## Routes/Endpoints
`receipts/process`

- Takes a request with a content application/json header and a JSON body. Body must include the following (all are string params except where indicated otherwise):
  - retailer
  - purchaseDate
  - purchaseTime
  - items (arr, each item will have the following):
    - shortDescription
    - price
  - total
- Returns an ID string for the receipt as JSON in the response.
 
`receipts/{id}/points`

- Takes a request with an ID string as a URL param.
- It returns the calculated points as a JSON in the response.

## Notes
- This service uses in-memory storage for data.
- Restarting the app will erase all stored receipts.
- API follows RESTful principles and returns JSON responses.

## Other Considerations
- Before this project, I was unfamiliar with `Golang/Go`. However, I was able to quickly learn its syntax and build a functional API.
- I chose to use the `Gin` framework to build the API. While I considered using Go’s native `net/http` package and the `mux` router, Gin provided the best balance of performance, simplicity, scalability, and built-in middleware. Given that this project involves handling structured JSON data and frequent API requests, Gin allowed me to keep my code clean and efficient without unnecessary complexity.  
- Additionally, Gin’s high-performance request routing, built-in middleware, and optimized request handling make it a great choice for scalable, production-ready APIs which would mirror a real-world app.
- With more time and if the project requirements allowed it, I would implement persistent storage using a database (e.g., PostgreSQL, MongoDB, or DynamoDB) instead of storing data in memory, ensuring that receipts are not lost on server restart.
- In a real-world application, I would implement authentication and authorization to restrict access to the API.  
- `JWT (JSON Web Tokens)` would be an ideal choice for stateless authentication, ensuring only authenticated users can process and retrieve receipts.  
- If third-party logins were required, `OAuth 2.0` (e.g., Google, GitHub authentication) could be implemented.  
- For internal services, `API keys` could be used as a lightweight authentication mechanism.
- I would also include extensive unit tests and integration tests to validate business logic and API reliability.
- Rate limiting could be implemented to prevent abuse and improve API security.

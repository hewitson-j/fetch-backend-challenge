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
- This service is stateless, meaning data is stored in-memory.
- Restarting the app will erase all stored receipts.
- API follows RESTful principles and returns JSON responses.

# Exchange Rate Service


A backend api written in **GO** that provides exchange rate data between currencies.
Support currency exchange between  
○ United States Dollar (USD) 
○ Indian Rupee (INR) 
○ Euro (EUR) 
○ Japanese Yen (JPY) 
○ British Pound Sterling (GBP)
Up to 90 days from the current date 
Support in-memory caching using memcache
Have exposed "/metrics" using Prometheus

## ENDPOINTS
- Get latest exchange rates at ("/api/latest")
  - query params 'from', 'to'.
- Convert amount between currencies ("/api/exchange")
  - query params 'from', 'to', 'amount'
- Get historical exchange rates up to past 90 days ("/api/convert")
  - query params 'from', 'to', 'date'.

## 🛠️ Requirements
- [Go 1.22+](https://go.dev/dl/)  
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerized runs)


## 🚀 Running Locally
### 1. Clone the repo
git clone https://github.com/Black-tag/exchange-rate-service.git
cd exchange-rate-service



### 2. Run with Go 
go run cmd/server/main.go
      Server starts at: http://localhost:8080

## Running with Docker 
docker run -p 8080:8080 exchange-rate-service:latest
      Server starts at: http://localhost:8080


## examples with Postman
Download the collection: [exchange-rate-service.postman_collection.json](./postman_collection.json)

# to use :
1. Open Postman
2. Click "import"
3. Select the `postman_collection.json` file
4. Run the requests





# totality-api-gateway

This server acts as the Api gateway for the grpc server - Which converts the http request to grpc and communicates with grpc user-server

add .env file in the root directory and include your desired port eg: PORT=:8080
and include a port for the grpc server eg USERCLIENTPORT=localhost:8081

used gin developer mode so you are able to see the routes - "userID" is the variable name used in json parsing 

{{baseurl}}/user/details - json input =  {"userID": 2}

{{baseurl}}/user/list/details - json input = {"userID":[1,2,10]}

Added sample user details for user Ids ranging from 1 to 5.

run the command - go run cmd/main.go

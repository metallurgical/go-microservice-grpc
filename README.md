# Introduction
Implement gRPC(Remote Procedure Call) to enable microservice communicate to each other which using `protobuf` as a service definition without relying on `json`, `xml` to interchange the data between client and server.

## How to install & run
To test in your local:

**[Option 1]** Just clone this repo into your local directory and run :
- Run server - `go run server/main.go`
- Run client - `go run client/main.go`

**[Option 2]** Make sure you install `go` and set  `GOPATH`. Then do following:
 - `go get github.com/metallurgical/go-microservice-grpc/server`
 - `go get github.com/metallurgical/go-microservice-grpc/client`
 - This will compile and install the client and server into `GOBIN` directory.
 - Run `server` and `client` in terminal
 
 **[Option 3]** Just clone this repo into your local directory and run :
 - Go to server directory, run `go build && ./server`
 - Go to client directory, run `go build && ./client`
 - This will compile the server and client and placing its executable files on current directory
 
 **[Option 4]** Just clone this repo into your local directory and run :
  - Go to server directory, run `go install && server`
  - Go to client directory, run `go install && client`
  - This will compile and install both client and server and placing its executable files into `GOBIN` directory
  
  
  For any options you followed above, you should see the results :
  
  ```
2019/04/12 12:25:43 Successfull create new contact with ID 1
2019/04/12 12:25:43 Found : name:"Norlihazmey Ghazali" mobile_no:"60123456789" office_no:"034543456" email:"norlihazmey.ghazali@gmail.com" id:1 address:<line_1:"Lot 123, Block A" line_2:"Jalan 123" postcode:"68100" country:"Selangor" city:"Batu Caves" > address:<line_1:"Lot 345, Block B" line_2:"Jalan 345" postcode:"765678" country:"Selangor" city:"Batu Caves" >
```


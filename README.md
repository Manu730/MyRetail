 MyRetail is a small web server which gets you the product with the given id and also updates it's price.

prerequisites :

To run this we need to install  Golang and MongoDB (optional) if not present.

To Install Golang Follow the link : https://golang.org/doc/install

To Install MongoDB Follow the link : https://docs.mongodb.com/manual/installation/



Directory Structure :

Code is divided into three packages.

common : contains code related to objects and constants

handler  : this is where all the fun happens, all hosted api actions happen here

db      :  mongodb database related functions

start.go : this is the entry point or main package



External FrameWorks :

This code uses two external libraries. GorillaMux and MongoDB driver, To get these two libraries in your local please follow the below commands

GorillaMux : go get -u github.com/gorilla/mux (Library Link : https://github.com/gorilla/mux)

MongoDb  : go get go.mongodb.org/mongo-driver (Source Link : https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial)


To Run :

We can directly run using command     go run start.go (Hosts on port 8088)

To build binary:

Linux : GOOS=linux go build 
Windows : GOOS=windows go build
Mac     : GOOS=darwin go build

Test :

handler package has a test case file, which can be executed using go test -v

 

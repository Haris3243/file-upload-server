# File Upload Server in Golang
1. clone this repository and run the following commad to build executable
```go build ./main.go```
1. now run the command to conatinerize(build image) you application
 ```docker run -t build server:latest .\```
1. now run the docker conatiner to start the server.
  ```docker run -d --name goserver -p 8080:8083 server:latest```
> this will start the container, here goserver is the name of container and -p flag maps host's port 8080 to container port 8083. as we configured our application to 8083 and exposed it to container so we have to bind it with any host port in order to access. we can't access container port directly 

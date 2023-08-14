# mini blog

### setup mongo
1. pull mongo 4.4
`docker pull mongo:4.4`

2. run mongo on your machine 
`docker run -d -p 27017:27017 --name my-mongo mongo:4.4`

### how to run
1. install go version 1.19
2. run command `go run main.go`
3. open swagger to see api documentation (optional) `http://localhost:8080/swagger/index.html`
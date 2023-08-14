# mini blog

### how to set up mongo
1. pull mongo 4.4
`docker pull mongo:4.4`

2. run mongo on your machine 
`docker run -d -p 27017:27017 --name my-mongo mongo:4.4`

### prerequisite
- set up mongo
- install go version `1.19`

### how to run
1. run command `go run main.go`
2. open swagger to see api documentation (optional) `http://localhost:8080/swagger/index.html`
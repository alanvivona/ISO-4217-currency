test:
	go test -covermode=count -coverprofile=cov.out ./... && \
	go tool cover -func=cov.out
	
cov: test
	go tool cover -html=cov.out

build:
	go build ./...
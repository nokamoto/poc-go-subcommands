
all: fmt
	dep ensure
	go test .
	go install .

fmt:
	gofmt -d .
	gofmt -w .

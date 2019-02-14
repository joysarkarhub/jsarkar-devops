GO-FLAGS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	${GO-FLAGS} go build -o jsarkar-devops
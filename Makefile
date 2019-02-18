APP=jsarkar-devops
GO-FLAGS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

clean:
	rm -rf ${APP}
build:
	${GO-FLAGS} go build -o jsarkar-devops
image: build
	docker build -t ${APP} .

FROM docker.io/golang:1.23
LABEL authors="hendisantika"
COPY . ./

WORKDIR test/integration/test-server
RUN ls

CMD go run server.go

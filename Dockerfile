FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

# this will perform unit test
RUN go test -v 
# we are team-container
RUN go build -o /main.go

EXPOSE 8080

CMD [ "/main.go" ]

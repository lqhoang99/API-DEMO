FROM golang:1.14
WORKDIR /API-DEMO-master
COPY go.mod go.sum ./
RUN go mod download 
COPY . .

RUN go build main.go
EXPOSE 3000
CMD ["./main"]


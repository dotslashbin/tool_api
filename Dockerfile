FROM golang:latest

LABEL maintainer="Joshua Fuentes <joshuarpf@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o tool_api . 

EXPOSE 5000 

# CMD ["./tool_api"]

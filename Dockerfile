FROM golang:1.12

# Set the Current Working Directory inside the container
WORKDIR  /app
COPY ./go.mod .
COPY ./go.sum .
COPY ./main.go .
RUN go mod download
COPY . .
RUN  CGO_ENABLED=0 go build -o ./fiyatbot 

FROM alpine
RUN apk add ca-certificates && update-ca-certificates
COPY --from=0 /app/fiyatbot /bin/fiyatbot
FROM golang:alpine


# golang code
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# air for hot reload
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./blog_build ./cmd/server/main.go

# for http server
EXPOSE 80

CMD ["./blog_build"]

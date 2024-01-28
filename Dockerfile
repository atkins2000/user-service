#use golang official image
FROM golang

WORKDIR /app

# Download all dependencies. 
COPY go.mod go.sum ./
RUN go mod download

# Copying source code
COPY . .

# Build
RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/api/main.go

FROM alpine:latest

WORKDIR /app

# Create a non-root user
RUN adduser --disabled-password --gecos '' rootless

# Change onwership of the app dir
RUN chown -R rootless /app

# Switch to the new user
USER rootless


COPY --from=builder /app/app .
COPY resourses/migration.txt ./migration.txt

EXPOSE 8080

# Swithc Gin to release mod
ENV GIN_MODE release

# Command to run the application
CMD ["./app", "--migration", "migration.txt"]

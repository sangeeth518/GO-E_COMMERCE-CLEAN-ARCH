FROM golang:1.26-alpine AS builder


WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# # Stage 2 — Runtime

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata


WORKDIR /app

COPY --from=builder /app/main .


EXPOSE 3000

CMD ["./main"]





# # Stage 1 — Builder
# Use official Go image with Alpine Linux (smaller than full Ubuntu)
# golang:1.22-alpine has Go compiler installed
# FROM golang:1.22-alpine AS builder
# # AS builder = give this stage a name so Stage 2 can reference it

# # Install git — needed by go mod download for some packages
# RUN apk add --no-cache git
# # apk = Alpine's package manager (like apt for Ubuntu)
# # --no-cache = don't store package cache (keeps image smaller)

# # Set working directory inside container
# # all following commands run from /app
# WORKDIR /app

# # Copy ONLY go.mod and go.sum first
# # Why? Docker caches each line as a layer
# # If only your code changes (not dependencies)
# # Docker reuses the cached go mod download layer
# # much faster rebuilds
# COPY go.mod go.sum ./

# # Download all dependencies
# RUN go mod download

# # Now copy all source code
# # This is after mod download so dependency layer is cached separately
# COPY . .

# # Build the binary
# # CGO_ENABLED=0 = no C dependencies (pure Go binary)
# # GOOS=linux = build for Linux (EC2 runs Linux)
# # -o main = output file named "main"
# # ./cmd/api = your main package location
# RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# # Stage 2 — Runtime
# # Alpine Linux only — no Go compiler
# # much smaller final image
# FROM alpine:latest

# # needed for HTTPS calls (S3, Razorpay) and timezone data
# RUN apk add --no-cache ca-certificates tzdata

# WORKDIR /app

# # Copy ONLY the compiled binary from builder stage
# # everything else (Go compiler, source code) is discarded
# COPY --from=builder /app/main .

# # Your app listens on port 3000
# EXPOSE 3000
# # EXPOSE is documentation only — does not actually open the port
# # port mapping happens in docker-compose

# # Command to run when container starts
# CMD ["./main"]
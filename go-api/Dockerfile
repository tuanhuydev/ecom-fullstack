FROM golang:1.24 AS builder

WORKDIR /app
# Install dependencies
COPY ./go.mod ./go.sum  ./
RUN go mod download
# Copy source
COPY . .

# Map variables from build args
ARG DB_HOST
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME

# Set as environment variables
ENV DB_HOST=${DB_HOST}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}


# Build code
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/api ./cmd

FROM golang:alpine3.21 AS runner
WORKDIR /app

# Install certificates
RUN apk --no-cache add ca-certificates

# Copy built code
COPY --from=builder /app/bin/api /app/api

# Set executable permissions
RUN chmod +x /app/api

# Create a non-root user to run the application
RUN adduser -D -u 1000 appuser
USER appuser

# Expose the application port
EXPOSE 8080

# Command to run
CMD ["/app/api"]



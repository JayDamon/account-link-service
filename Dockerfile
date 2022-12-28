# FROM golang:1.18-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o plaidIntegration ./cmd/main

# RUN chmod +x /app/plaidIntegration

FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/plaidIntegration /app
COPY plaidIntegration /app

CMD ["/app/plaidIntegration"]
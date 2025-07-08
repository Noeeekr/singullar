FROM golang:1.24.4-alpine AS build
WORKDIR /singullar
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api ./internal/api/

FROM alpine:latest
WORKDIR /singullar
COPY --from=build /singullar/api .
CMD ["./api", "start", "--enable-migrations", "--ignore-existing"]
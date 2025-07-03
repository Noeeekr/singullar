# Composer context
FROM golang:1.24.4-alpine AS build
WORKDIR /go/singullar
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o . ./internal/server

FROM alpine:latest
COPY --from=build ./go/singullar/server ./server
COPY ./dist ./dist 
CMD ["./server", "serve", "./dist"]
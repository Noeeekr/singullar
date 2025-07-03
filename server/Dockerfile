FROM golang:1.24.4-alpine
WORKDIR /singullar
COPY go.mod go.sum ./
RUN go mod download
COPY . .
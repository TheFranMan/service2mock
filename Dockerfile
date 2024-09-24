FROM golang:1.23 AS base-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /service2

FROM alpine:3.14 AS release
WORKDIR /
COPY --from=base-stage /service2 /service2
EXPOSE 3001
ENTRYPOINT [ "/service2" ]
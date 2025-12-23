FROM golang:1.25-alpine3.23 AS builder
WORKDIR /app
COPY ./ ./
RUN go build -o /bin/drcjudge main.go

FROM scratch
COPY --from=builder /bin/drcjudge /bin/drcjudge
ENTRYPOINT ["drcjudge"]


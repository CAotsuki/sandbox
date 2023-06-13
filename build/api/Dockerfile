FROM golang:alpine3.18 AS builder
WORKDIR /go/src

COPY ../../go.mod ../../go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd/main.go


FROM golang:alpine3.18 AS runner
WORKDIR /go/bin

COPY --from=builder /go/src/app .
COPY ../../build/.env .
CMD [ "./app" ]
EXPOSE 8080

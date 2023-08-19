FROM golang:1.20-alpine
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o payments

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=0 /src/payments /bin/payments
ENTRYPOINT ["payments"] 
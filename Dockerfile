FROM golang:1.18-alpine as build
WORKDIR /build
RUN apk --update --no-cache add build-base
COPY . .
RUN go mod download && \
    go get && \
    go mod tidy && \
    go build -o seedtopia-service -buildvcs=false

FROM golang:1.18-alpine
WORKDIR /app
COPY --from=build /build/seedtopia-service /app/seedtopia-service
COPY .env .

ARG PORT
EXPOSE $PORT
ENTRYPOINT ["/app/seedtopia-service"]
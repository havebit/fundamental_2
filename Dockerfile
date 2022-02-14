FROM golang:1.17-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./cmd/api/ ./
ENV GOARCH=amd64
RUN go build -o /go/bin/app

## Deploy
FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/app /app
EXPOSE 8081
USER nonroot:nonroot
CMD ["/app"]
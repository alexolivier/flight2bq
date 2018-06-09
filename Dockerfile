
FROM golang:1.10-alpine AS builder
RUN apk add --no-cache git mercurial 
# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/alexolivier/flight2bigquery
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN go build -o /app .

FROM alpine
COPY --from=builder /app ./
ENTRYPOINT ["./app"]

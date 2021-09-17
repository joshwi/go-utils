FROM golang:1.16-alpine
WORKDIR /app
RUN ls -la
COPY go.mod ./
COPY go.sum ./
COPY app /app
RUN ls -la
RUN go mod download
RUN go build -o /nfl_collector
CMD ["/nfl_collector"]
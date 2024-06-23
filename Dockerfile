FROM golang:alpine
WORKDIR /api_server
COPY . .
RUN go mod download
RUN go build -tags netgo -o api_server main.go
EXPOSE 8000
ENV TASKPOOLSIZE 1000
CMD ["./api_server"]
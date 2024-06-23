FROM golang:alpine
WORKDIR /api_server
COPY . .
ENV TASKPOOLSIZE=1000
ENV SERVERPORT=8080
RUN go mod download
RUN go build -tags netgo -o api_server main.go
EXPOSE 8080
CMD ["./api_server"]
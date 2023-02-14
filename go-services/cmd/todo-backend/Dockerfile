FROM golang:1.19 AS build
WORKDIR /go/src
COPY . .

ENV CGO_ENABLED=0
ENV GOARCH=amd64
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o todo-backend ./cmd/todo-backend

FROM alpine:3.15 AS runtime
ENV TODO_BACKEND_API_ENDPOINT=:80
ENV TODO_BACKEND_MAINTENANCE_ENDPOINT=:8000
COPY --from=build /go/src/todo-backend ./
EXPOSE 80/tcp
EXPOSE 8000/tcp
ENTRYPOINT ["./todo-backend"]

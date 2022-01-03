FROM    golang:1.17.0-alpine3.14
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod tidy -go=1.16 && go mod tidy -go=1.17
RUN go build -o bin/main main.go model.go server.go handlers.go
EXPOSE 4006
CMD [ "bin/main" ]
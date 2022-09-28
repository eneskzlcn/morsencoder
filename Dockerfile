FROM    golang:1.18-alpine3.16
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod tidy -go=1.18
RUN go build -o bin/morsencoder ./cmd/morsencoder
EXPOSE 4200
CMD [ "bin/morsencoder" ]
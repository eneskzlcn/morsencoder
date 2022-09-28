## morsencoder

### About Project 

morsencoder is an api that listens for any request
to the endpoint "/encode" with a query parameter text
like "/encode?text=text to morse code encode", encodes
that text and returns it back to the client as string.

## Letter Separation
morsencoder uses letter separator as one whitespace. Each
letter encoded to morse code will be seperated by one whitespace.

Ex: given text "ab" will be encoded as `.- -...` where `.-` is
the `a` letter and `-...` is the `b` letter in morse code alphabet.

## Word Separation

morsencoder uses letter separator as one backslash between two whitespaces(' / '). Each
word encoded to morse code will be seperated by one backslash between two whitespaces(' / ').

Ex: given text `aa bb` will be encoded as `.- .- / -... -...`
where `.-` is the `a` letter and `-...` is the `b` letter and ` ` is the letter separator in morse code alphabet.

### How To Build

You can directly use the make command written for build 
```shell
make build
```
to build the application.

For docker builds, you can use 
```shell
make dockerize
```
to build the image with the name eneskzlcn/morsencoder and 
with tag "latest". If you build with another name and tag 
you can configure the `dockerize` command in Makefile or you 
can directly execute the following command:

```shell
docker build -t <image-name>:<image-tag> .
```

### How To Test

Before you execute the tests, you need to execute the command 
```shell
make generate-mocks
```
to generate all necessary mocks for tests.

After mocks generated, you can simply execute 
```shell
make unit-tests
``` 
to
run all unit tests, or you can use following command:
```shell
go test -v ./...
``` 
for run all tests or 
```shell
go test -v .internal/...
``` 
to just run domain tests.

### How To Run
If you have `go 1.18` installed on your machine, you can simply
build the application with command 
```shell
go build -o bin/morsencoder .cmd/morsencoder
``` 
and then run 
```shell
./bin/morsencoder
```
or
```shell
go run ./cmd/morsencoder/main.go
```

If you have docker installed on your machine, you can also run
application as docker container with the following commands
```shell
docker build -t <image-name>:<image-tag> .
```
to build docker image, and 
```shell
docker run -p <host-port>:4200 <image-name>:<image-tag>
```
to run docker image as docker container.
### Run Linter
I use `golangci-lint` default linter. If you have
`golangci-lint` installed on your machine, you can
simply use 
```shell
golangci-lint run
```
to run linters.

### Test Coverage
I use GoLand IDE which provides me to test coverage of my code.

If you use GoLand IDE you can simply go to the `./internal/morsencoder`
and right click on `morsencoder_test` package. Then select `Run with Coverage` from menu which
brings you a coverage report that contains your coverage as percentage on chosen package (morsencoder).

I have 100% coverage on the domain files `./internal/morsencoder`

### Letter Morse Code Key Map
```
'A': ".-"

'B': "-..."

'C': "-.-."

'D': "-.."

'E': "."

'F': "..-."

'G': "--."

'H': "...."

'I': ".."

'J': ".---"

'K': "-.-"

'L': ".-.."

'M': "--"

'N': "-."

'O': "---"

'P': ".--."

'Q': "--.-"

'R': ".-."

'S': "..."

'T': "-"

'U': "..-"

'V': "...-"

'W': ".--"

'X': "-..-"

'Y': "-.--"

'Z': "--.."

'1': ".----"

'2': "..---"

'3': "...--"

'4': "....-"

'5': "....."

'6': "-...."

'7': "--..."

'8': "---.."

'9': "----."

'0': "-----"

'.': ".-.-.-"

',': "--..--"

'?': "..--.."

'!': "-.-.--"

'-': "-....-"

'/': "-..-."

'@': ".--.-."

'(': "-.--."

')': "-.--.-"

```
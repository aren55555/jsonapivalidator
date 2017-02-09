# jsonapivalidator [![Build Status](https://travis-ci.org/aren55555/jsonapivalidator.svg?branch=master)](https://travis-ci.org/aren55555/jsonapivalidator) [![GoDoc](https://godoc.org/github.com/aren55555/jsonapivalidator?status.svg)](http://godoc.org/github.com/aren55555/jsonapivalidator)

A validator that determines whether arbitrary JSON payloads are in compliance with the [JSON API specification](http://jsonapi.org/). If the JSON is not compliant this validator will indicate the reason why.

This validator has been written in [Go](https://golang.org/), but it is easy to use it in a Javascript environment thanks to [GopherJS](https://gopherjs.github.io/); see [github.com/aren55555/corroborate](https://github.com/aren55555/corroborate) for a working JS implementation.

## Installation

```
go get -u github.com/aren55555/jsonapivalidator
```

## Examples

### Valid JSON API Payload
```go
req, err := http.DefaultClient.Get("https://raw.githubusercontent.com/aren55555/jsonapivalidator/master/test/samples/valid/default.json")

result, err := jsonapivalidator.UnmarshalAndValidate(req.Body)
if err != nil {
  panic(err)
}

if result.IsValid() {
  fmt.Println("The JSON sample was valid!")
}
```

### Invalid JSON API Payload

```go
req, err = http.DefaultClient.Get("https://raw.githubusercontent.com/aren55555/jsonapivalidator/master/test/samples/invalid/default.json")

result, err = jsonapivalidator.UnmarshalAndValidate(req.Body)
if err != nil {
  panic(err)
}

fmt.Println("Errors:")
for i, err := range result.Errors() {
  fmt.Println("\t", i, err)
}

fmt.Println("Warnings:")
for i, err := range result.Warnings() {
  fmt.Println("\t", i, err)
}
```

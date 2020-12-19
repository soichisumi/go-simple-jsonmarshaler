# go-simple-jsonmarshaler

![Test](https://github.com/soichisumi/go-simple-jsonmarshaler/workflows/Test/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/soichisumi/go-simple-jsonmarshaler)](https://goreportcard.com/report/github.com/soichisumi/go-simple-jsonmarshaler) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A library that provides json.Marshal/Unmarshal/Encoder/Decoder that do not base64 encode []byte fields

## Motivation

Go base64 encodes the []byte field of struct when Marshal/Unmarshal it.

This is an appropriate way to safely represent byte data in JSON, but most external systems do not base64 encode the JSON string type, and in effect it is not possible to convert the JSON string field of an external system into a []byte field.

If the struct field is a string, it can be marshal/unmarshal without any problem, but if the type is determined by the external (e.g. IDL such as Protocol Buffers), this method cannot be used.

As a workaround for this problem, I have created a library that converts the []byte field of struct to JSON without base64 encoding.

## Installation

* `go get -u github.com/soichisumi/go-simple-jsonmarshaler`

## Usage

```go
import sjson "github.com/soichisumi/go-simple-jsonmarshaler"

...

err := sjson.Marshal(foo)
```

<!-- ## How it works-->


# fission-go-form-req-transformer
[![Coverage Status](https://coveralls.io/repos/github/OpenIndustryCloud/fission-go-form-req-transformer/badge.svg?branch=master)](https://coveralls.io/github/OpenIndustryCloud/fission-go-form-req-transformer?branch=master)

The `go` runtime uses the [`plugin` package](https://golang.org/pkg/plugin/) to dynamically load an HTTP handler.

## Requirements

First, set up your fission deployment with the go environment.

```
fission env create --name go-env --image fission/go-env:1.8.1
```

To ensure that you build functions using the same version as the
runtime, fission provides a docker image and helper script for
building functions.

## Example Usage

### form-data-transformer.go

`form-data-transformer.go` is an API which accepts JSON payload from TYPE FORM and transform the JSON data which can used by other APIs to query weather info or registering a ticket to ZenDesk.

```bash
# Download the build helper script
$ curl https://raw.githubusercontent.com/fission/fission/master/environments/go/builder/go-function-build > go-function-build
$ chmod +x go-function-build

# Build the function as a plugin. Outputs result to 'function.so'
$ go-function-build form-data-transformer.go

# Upload the function to fission
$ fission function create --name transform-data --env go-env --package function.so

# Map /hello to the hello function
$ fission route create --method POST --url /transform-data --function transform-data

# Run the function
$ curl -d '{--INPUT JSON--}' -H "Content-Type: application/json" -X POST http://$FISSION_ROUTER/transform-data

#sample input

#sample output


```

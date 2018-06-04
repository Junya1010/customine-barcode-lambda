Customine Barcode Lambda
====

Simple QR code image generator.

## Description

This package is simple barcode image generator. This lambda function responses QR code image on the fly via BASE64 img tag (It means HTML text response).

You can use easily this with Gusuku Customine and Cybozu Kintone.

## Requirement

Requirements management depend on dep.

```
go get -u github.com/golang/dep/cmd/dep
```

**used libraries**

* https://github.com/boombuler/barcode

## Build

```
$ GOOS=linux GOARCH=amd64 go build -o barcodeMaker
$ zip handler.zip ./barcodeMaker
```

## Usage

Request parameters:

```
{
    "sentence": "the message you want to set in barcode image",
    "width": 200,
    "height": 200
}
```

Response parameters:

```
{
    "img": "......"
}
```


## Install


## Licence

MIT

## Author

[Koichiro Nishijima](https://github.com/k-nishijima/)

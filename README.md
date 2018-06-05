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

## Build and Deploy with AWS SAM

```
$ GOOS=linux GOARCH=amd64 go build -o build/barcodeMaker

$ aws cloudformation package \
    --profile your_profile_name \
    --template-file template.yml \
    --s3-bucket your_bucket \
    --region your_region \
    --s3-prefix your_bucket_prefix \
    --output-template-file .template.yml
$ aws cloudformation deploy \
    --profile your_profile_name \
    --template-file .template.yml \
    --capabilities CAPABILITY_IAM \
    --stack-name yourStackName

$ aws cloudformation describe-stack-events \
    --profile your_profile_name \
    --stack-name yourStackName
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
    "img": "<img src=......"
}
```

## License

Apache License

## Author


[Koichiro Nishijima](https://github.com/k-nishijima/) / [R3 institute](https://www.r3it.com/)

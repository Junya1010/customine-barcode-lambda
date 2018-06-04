package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/png"

	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Event struct {
	Sentence string `json:"sentence" valid:"required"`
	Width    int    `json:"width" valid:"range(0|800)"`
	Height   int    `json:"height" valid:"range(0|800)"`
}

type Response struct {
	ImgTag string `json:"img"`
}

func barcodeMaker(event Event) (Response, error) {
	if event.Width == 0 {
		event.Width = 200
	}
	if event.Height == 0 {
		event.Height = 200
	}

	_, err := govalidator.ValidateStruct(event)
	if err != nil {
		return Response{ImgTag: errImage}, errors.New(err.Error())
	}

	if qrCode, err := qr.Encode(event.Sentence, qr.M, qr.Auto); err == nil {
		if qrCode, err = barcode.Scale(qrCode, event.Width, event.Height); err == nil {
			out := new(bytes.Buffer)
			png.Encode(out, qrCode)

			return Response{
				ImgTag: fmt.Sprintf(`<img src="data:image/png;base64,%s">`, base64.StdEncoding.EncodeToString(out.Bytes())),
			}, nil
		}
		return Response{ImgTag: errImage}, errors.New(err.Error())
	}
	return Response{ImgTag: errImage}, errors.New(err.Error())
}

func main() {
	lambda.Start(barcodeMaker)
}

const errImage = `<img src="data:image/png;base64,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">`

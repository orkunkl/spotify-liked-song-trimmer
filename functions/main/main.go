package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	trimmer "github.com/orkunkl/spotify-liked-song-trimmer"
)

func main() {
	lambda.Start(trimmer.Handler)
}

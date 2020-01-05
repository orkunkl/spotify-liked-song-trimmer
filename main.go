package main

import (
"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.CloudWatchEvent) {
	fmt.Printf("Detail = %s\n", event.Detail)
}

func main() {
	lambda.Start(handler)
}

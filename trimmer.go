package trimmer

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(ctx context.Context) (int, error) {
	return trim()
}

func Handler(ctx context.Context, event events.CloudWatchEvent) (int, error) {
	return trim()
}

func trim() (int, error) {
	return 1, nil
}

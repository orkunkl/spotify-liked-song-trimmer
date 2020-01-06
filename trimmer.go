package trimmer

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       int `json:"body"`
}

func Handler(ctx context.Context , event events.CloudWatchEvent) (*Response, error) {
	result, err := trim()
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: 200,
		Body:       result,
	}, nil
}
func TestHandler() (Response, error) {
	result, err := trim()
	if err != nil {
		return Response{
			StatusCode: 500,
			Body: 1,
		}, err
	}

	return Response{
		StatusCode: 200,
		Body:       result,
	}, nil
}

func trim() (int, error) {
	return 1, nil
}

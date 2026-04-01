package prodneutral

import (
	"context"
	"fmt"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

func settingHostname(hostname string) error {
	// hostname := "eu-vision.googleapis.com:443"

	ctx := context.Background()
	opt := option.WithEndpoint(hostname)
	client, err := vision.NewImageAnnotatorClient(ctx, opt)
	if err != nil {
		return fmt.Errorf("NewImageAnnotatorClient: %v", err)
	}
	defer client.Close()

	return nil
}

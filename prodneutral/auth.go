package prodneutral

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func withADCCredentials() error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	return nil
}

func withManualCredentials(jsonPath string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	defer client.Close()

	if err != nil {
		return err
	}
	return nil
}

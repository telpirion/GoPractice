package prodneutral

import (
	"context"
	"log"

	"cloud.google.com/go/translate"
	"google.golang.org/api/option"
)

func setConnectionPool(num int) error {
	ctx := context.Background()
	clientOption := option.WithGRPCConnectionPool(num)

	translateClient, err := translate.NewClient(ctx, clientOption)
	if err != nil {
		return err
	}
	defer translateClient.Close()

	languages, _ := translateClient.DetectLanguage(context.Background(), []string{
		"Urbs antiqua fuit (Tyrii tenuere coloni) Karthago, Italiam contra Tiberinaque longe",
	})
	for _, l := range languages {
		log.Println(l)
	}
	return nil
}

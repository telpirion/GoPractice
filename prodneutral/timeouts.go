package prodneutral

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func usingContextTimeout(seconds int64) error {

	translateClient, err := translate.NewClient(context.Background())
	if err != nil {
		return err
	}
	defer translateClient.Close()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(seconds*int64(time.Second)))
	defer cancel()

	translations, err := translateClient.Translate(ctx, []string{
		"guten tag", "guten morgen",
	}, language.English, &translate.Options{
		Source: language.German,
	})

	if err != nil {
		return err
	}

	for _, t := range translations {
		log.Println(t)
	}
	return nil
}

func usingCallOptions(seconds int64) error {
	ctx := context.Background()
	opt := option.WithGRPCDialOption(
		grpc.WithTimeout(time.Duration(seconds * int64(time.Second))),
	)

	translateClient, err := translate.NewClient(ctx, opt)
	defer translateClient.Close()
	if err != nil {
		return err
	}

	languages, _ := translateClient.DetectLanguage(context.Background(), []string{
		"Omnis Gallia est divisa in tres partes",
	})
	for _, l := range languages {
		log.Println(l)
	}
	return nil
}

func settingKeepAlive(seconds int) error {
	ctx := context.Background()

	keepAlive := keepalive.ClientParameters{
		Time:                time.Duration(seconds) * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}

	dialOpts := grpc.WithKeepaliveParams(keepAlive)
	opt := option.WithGRPCDialOption(dialOpts)

	translateClient, err := translate.NewClient(ctx, opt)
	defer translateClient.Close()
	if err != nil {
		return err
	}

	languages, _ := translateClient.DetectLanguage(context.Background(), []string{
		"Arma virumque cano, Troiae qui prima ab oris",
	})
	for _, l := range languages {
		log.Println(l)
	}
	return nil
}

package prodneutral

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/translate"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

func settingRetries(seconds int32) error {
	ctx := context.Background()

	dialOption := grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay: time.Millisecond,
			MaxDelay:  time.Second,
		},
		MinConnectTimeout: time.Duration(seconds * int32(time.Second)),
	})

	opt := option.WithGRPCDialOption(
		dialOption,
	)

	translateClient, err := translate.NewClient(ctx, opt)
	defer translateClient.Close()
	if err != nil {
		log.Printf("error: %v\n", err)
		return err
	}

	languages, _ := translateClient.DetectLanguage(context.Background(), []string{
		"Interea ea legione quam secum habebat, militibusque qui ex provincia convenerant",
	})
	for _, l := range languages {
		log.Println(l)
	}
	return nil
}

func settingBackoffPolicy(minTimeout, baseDelay, maxDelay int, multiplier, jitter float64) error {
	ctx := context.Background()

	dialOption := grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  time.Duration(baseDelay * int(time.Second)),
			MaxDelay:   time.Duration(maxDelay * int(time.Second)),
			Multiplier: multiplier,
			Jitter:     jitter,
		},
		MinConnectTimeout: time.Duration(int32(minTimeout) * int32(time.Second)),
	})

	opt := option.WithGRPCDialOption(
		dialOption,
	)

	translateClient, err := translate.NewClient(ctx, opt)
	defer translateClient.Close()
	if err != nil {
		log.Printf("error: %v\n", err)
		return err
	}

	languages, _ := translateClient.DetectLanguage(context.Background(), []string{
		"Hoc proelio trans Rhenum nuntiato Suebi, qui ad ripas Rheni venerant, domum reverti coeperunt",
	})
	for _, l := range languages {
		log.Println(l)
	}
	return nil
}

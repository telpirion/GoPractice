package prodneutral

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes"

	video "cloud.google.com/go/videointelligence/apiv1"
	videopb "google.golang.org/genproto/googleapis/cloud/videointelligence/v1"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/status"
)

func getLROName() error {
	ctx := context.Background()

	client, err := video.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return err
	}
	defer client.Close()

	op, err := client.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputUri: "gs://cloud-samples-data/video/cat.mp4",
		Features: []videopb.Feature{
			videopb.Feature_LABEL_DETECTION,
		},
	})
	if err != nil {
		log.Fatalf("Failed to start annotation job: %v", err)
		return err
	}

	log.Printf("Operation data: %v\n", op.Name())

	resp, err := op.Wait(ctx)
	if err != nil {
		log.Fatalf("Failed to annotate: %v", err)
	}

	result := resp.GetAnnotationResults()[0]

	for _, annotation := range result.SegmentLabelAnnotations {
		fmt.Printf("Description: %s\n", annotation.Entity.Description)

		for _, category := range annotation.CategoryEntities {
			fmt.Printf("\tCategory: %s\n", category.Description)
		}

		for _, segment := range annotation.Segments {
			start, _ := ptypes.Duration(segment.Segment.StartTimeOffset)
			end, _ := ptypes.Duration(segment.Segment.EndTimeOffset)
			fmt.Printf("\tSegment: %s to %s\n", start, end)
			fmt.Printf("\tConfidence: %v\n", segment.Confidence)
		}
	}
	return nil
}

func getLRODataFromName() error {
	ctx := context.Background()

	client, err := video.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return err
	}
	defer client.Close()

	op, err := client.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputUri: "gs://cloud-samples-data/video/cat.mp4",
		Features: []videopb.Feature{
			videopb.Feature_LABEL_DETECTION,
		},
	})
	if err != nil {
		log.Fatalf("Failed to start annotation job: %v", err)
		return err
	}

	lroName := op.Name()

	req := &longrunning.GetOperationRequest{
		Name: lroName,
	}
	resp, err := client.LROClient.GetOperation(ctx, req)

	if err != nil {
		return err
	}
	log.Println(resp.Metadata.TypeUrl)
	return nil
}

func getErrorCode() error {
	ctx := context.Background()

	client, err := video.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputUri: "gs://cloud-samples-data/video/this-doesnt-exist.mp4",
		Features: []videopb.Feature{
			videopb.Feature_LABEL_DETECTION,
		},
	})
	if err != nil {
		st := status.Convert(err)
		log.Println(st.Code())
		return err
	}
	return nil
}

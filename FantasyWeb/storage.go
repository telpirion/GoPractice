package fantasymapsweb

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
)

func SaveImageToStorage(
	fileBytes []byte,
	fantasyMap *FantasyMap,
	filename string,
	ctx context.Context,
) error {
	// TODO: Move bucket and folder name to config file
	bucketName := "fantasy-maps"
	folderName := "UserSubmitted"
	newObjectName := fmt.Sprintf("%s/%s", folderName, filename)

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// TODO: Handle timeouts, retries
	// Ex. https://github.com/GoogleCloudPlatform/golang-samples/blob/main/storage/objects/stream_file_upload.go

	buf := bytes.NewBuffer(fileBytes)
	writer := client.Bucket(bucketName).Object(newObjectName).NewWriter(ctx)

	if _, err = io.Copy(writer, buf); err != nil {
		log.Println("Cannot write file to Cloud Storage")
		return fmt.Errorf("Cannot write file to Cloud Storage")
	}

	// Store the GCS URI
	newGCSUri := fmt.Sprintf("gs://%s/%s", bucketName, newObjectName)
	fantasyMap.GcsURI = newGCSUri

	if err := writer.Close(); err != nil {
		log.Println("Cannot close Storage writer")
		return fmt.Errorf("Cannot close Storage writer")
	}

	return nil
}

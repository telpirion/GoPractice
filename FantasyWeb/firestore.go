package fantasymapsweb

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"os"

	firestore "cloud.google.com/go/firestore"
)

func StoreMapRecord(fileBytes []byte, fantasyMap *FantasyMap, ctx context.Context) (string, error) {
	projectId := os.Getenv("PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Error creating Firestore client: %v", err.Error())
		return "", err
	}

	documentId := GetDocumentID(fileBytes)
	if isInDB, err := CheckWhetherDocumentExists(documentId, ctx); isInDB {
		log.Printf("Document with ID %v already exists.", documentId)
		return documentId, nil
	} else if err != nil {
		return "", err
	}

	log.Printf("Document ID: %v", documentId)

	// TODO: Store Collection name in config file
	client.Collection("FantasyMaps").Doc(documentId).Set(ctx, map[string]interface{}{
		"computedBBoxes":  fantasyMap.ComputedBBoxes,
		"gcsURI":          fantasyMap.GcsURI,
		"filename":        fantasyMap.Filename,
		"predictedBBoxes": fantasyMap.PredictedBBoxes,
		"source":          fantasyMap.Source,
		"userID":          fantasyMap.UserID,
	})

	return documentId, nil
}

// Check whether the specified document exists
func CheckWhetherDocumentExists(documentId string, ctx context.Context) (bool, error) {

	projectId := os.Getenv("PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Error creating Firestore client: %v", err.Error())
		return false, err
	}
	defer client.Close()

	_, err = client.Collection("FantasyMaps").Doc(documentId).Get(ctx)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Translate the file's SHA1 hash value into a unique string ID
func GetDocumentID(fileBytes []byte) string {
	hashFunc := sha1.New()
	hashFunc.Write(fileBytes)
	documentIdBytes := hashFunc.Sum(nil)
	documentId := hex.EncodeToString(documentIdBytes)
	return documentId
}

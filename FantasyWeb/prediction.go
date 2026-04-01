package fantasymapsweb

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	"google.golang.org/api/option"
	aiplatformpb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1"
	instancepb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1/schema/predict/instance"
	paramspb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1/schema/predict/params"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetOnlinePrediction(fileBytes []byte, fantasyMap *FantasyMap, ctx context.Context) (string, error) {

	// TODO: Move endpoint into a config file
	endpoint := "projects/fantasymaps-334622/locations/us-central1/endpoints/7423137250650619904"
	apiEndpoint := "us-central1-aiplatform.googleapis.com:443"

	imageData := base64.StdEncoding.EncodeToString(fileBytes)
	//log.Println(imageData)

	clientOption := option.WithEndpoint(apiEndpoint)

	client, err := aiplatform.NewPredictionClient(ctx, clientOption)
	if err != nil {
		log.Printf("Error creating PredictionClient: %v\n", err.Error())
		return "", err
	}
	defer client.Close()

	// Compose the prediction instance message
	instanceType := &instancepb.ImageObjectDetectionPredictionInstance{
		Content: imageData,
	}

	instanceJSON, err := protojson.Marshal(instanceType)
	if err != nil {
		log.Print("Error converting prediction instance to JSON")
		return "", err
	}

	var instanceMap map[string]interface{}
	if err := json.Unmarshal([]byte(instanceJSON), &instanceMap); err != nil {
		log.Print("Error translating instance JSON string to map")
		return "", err
	}

	instance, err := structpb.NewValue(instanceMap)
	if err != nil {
		log.Printf("Error creating prediction instance: %v\n", err.Error())
		return "", err
	}

	// Compose the prediction parameters
	parametersType := &paramspb.ImageObjectDetectionPredictionParams{
		ConfidenceThreshold: 0.5,
		MaxPredictions:      5,
	}

	parametersJSON, err := protojson.Marshal(parametersType)
	if err != nil {
		log.Print("Error converting parameters type into JSON")
		return "", err
	}

	var parametersMap map[string]interface{}
	if err := json.Unmarshal([]byte(parametersJSON), &parametersMap); err != nil {
		log.Print("Error converting parameterJSON to map")
		return "", err
	}

	parameters, err := structpb.NewValue(parametersMap)
	if err != nil {
		log.Printf("Error creating prediction parameter: %v\n", err.Error())
		return "", err
	}

	var instances = []*structpb.Value{instance}

	// Create a prediction request
	req := &aiplatformpb.PredictRequest{
		Endpoint:   endpoint,
		Parameters: parameters,
		Instances:  instances,
	}

	b, err := protojson.Marshal(req)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
	}
	log.Printf("%s", b)

	resp, err := client.Predict(ctx, req)
	if err != nil {
		log.Printf("Error sending prediction: %v\n", err.Error())

		if s, ok := status.FromError(err); ok {
			log.Println(s.Message())
			for _, d := range s.Proto().Details {
				log.Println(d)
			}
		}

		return "", err
	}

	log.Println(resp)

	return imageData, nil
}

func GetPredictionFromFunction(fileBytes []byte, fantasyMap *FantasyMap, ctx context.Context) (string, error) {
	url := "https://us-central1-fantasymaps-334622.cloudfunctions.net/test-1"
	contentType := "application/json; charset=utf-8"

	imageData := base64.StdEncoding.EncodeToString(fileBytes)
	postBody, err := json.Marshal(map[string]string{
		"content": imageData,
	})

	if err != nil {
		log.Printf("JSON body of POST message error: %v\n", err.Error())
		return "", err
	}
	body := bytes.NewBuffer(postBody)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Printf("Unable to build request to function: %v", err.Error())
		return "", err
	}

	req.Header.Set("Content-Type", contentType)

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Unable to complete POST request: %v", err.Error())
		return "", err
	}
	defer res.Body.Close()

	predictions, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Unable to read response data: %v", err.Error())
		return "", err
	}

	fantasyMap.PredictedBBoxes = string(predictions)

	return string(predictions), nil
}

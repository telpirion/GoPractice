package fantasymapsweb

import (
	"context"
	"io/ioutil"
	"testing"
)

func Test_GetOnlinePrediction(t *testing.T) {

	// Open the test file and read the contents
	buf, err := ioutil.ReadFile("../tests/resources/test_map.jpg")
	if err != nil {
		t.Error("Couldn't open test resource")
	}

	ctx := context.Background()
	f := FantasyMap{
		Filename:        "my-faked-file",
		Source:          "User-Submitted",
		GcsURI:          "gs://some-bucket",
		PredictedBBoxes: "{}",
		ComputedBBoxes:  "{}",
		UserID:          "fantasy.maps@gmail.com",
	}

	_, err = GetOnlinePrediction(buf, &f, ctx)
	if err != nil {
		t.Skip("Skipping direct Vertex AI online prediction")
	}
	//t.Log("It should return the imageData from the model")
	//t.Log(imageData)
}

package fantasymapsweb

import (
	"context"
	"io/ioutil"
	"testing"
)

func Test_SaveImageToStorage(t *testing.T) {

	t.Log("It should save a test image to Google Cloud Storage.")

	// Open the test file and read the contents
	fileBytes, err := ioutil.ReadFile("../tests/resources/test_map.jpg")
	if err != nil {
		t.Error("Couldn't open test resource")
	}
	ctx := context.Background()
	filename := "my-test-file.jpg"
	f := FantasyMap{
		Filename:        filename,
		Source:          "User-Submitted",
		GcsURI:          "",
		PredictedBBoxes: "{}",
		ComputedBBoxes:  "{}",
		UserID:          "fantasy.maps@gmail.com",
	}

	err = SaveImageToStorage(fileBytes, &f, filename, ctx)
	if err != nil {
		t.Error("Couldn't save test resource to GCS")
	}

	// TODO: Test that  f.GcsURI is populated
}

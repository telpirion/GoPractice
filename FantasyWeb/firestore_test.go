package fantasymapsweb

import (
	"context"
	"io/ioutil"
	"testing"
)

func Test_StoreMapRecord(t *testing.T) {
	t.Log("It should store information about the document in Firestore")

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

	imageID, err := StoreMapRecord(buf, &f, ctx)
	if err != nil {
		t.Error("Couldn't store the map")
	}
	t.Log("It should return the ID of the new DB document")
	t.Log(imageID)
}

package fantasymapsweb

// FantasyMap stores the data to be stored in the DB document.
type FantasyMap struct {
	DocumentID      string
	ComputedBBoxes  string // Bounding boxes computed from predicted bboxes
	Filename        string // Filename as submitted by the user
	GcsURI          string // Cloud Storage location of the image
	PredictedBBoxes string
	Source          string
	UserID          string
}

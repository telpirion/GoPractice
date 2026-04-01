package fantasymapsweb

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

/**
 */
func postSubmitMaps(c *gin.Context) {
	log.Println("Submit map postback!")

	filename := c.PostForm("file-name") // This is how we get the other form fields out
	log.Println(filename)

	fileHeader, err := c.FormFile("file-upload")
	if err != nil {
		log.Println(err)
	}

	if filename == "" {
		// TODO: Replace call to .Filename with something else, maybe timestamp
		filename = fileHeader.Filename
	}
	log.Println(filename)

	// Saves the file to a local directory
	// TODO: Delete the following since will be stored in GCS?
	filePath := fmt.Sprintf("./output/%v", filename)
	if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	// Extract and read the file contents
	file, err := fileHeader.Open()
	if err != nil {
		c.String(http.StatusBadRequest, "file open err: %s", err.Error())
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %v\n", err.Error())
		c.String(http.StatusBadRequest, "file read err: %s", err.Error())
		return
	}

	// TODO: Check file size; if larger than 1.2 MB, skip prediction

	// Check whether we already have this file in the database.
	documentId := GetDocumentID(buf)

	if isInDB, err := CheckWhetherDocumentExists(documentId, c); isInDB {
		c.String(http.StatusOK, "File already in database")
		return
	} else if err != nil {
		log.Printf("Error checking database: %v\n", err.Error())
		c.String(http.StatusBadRequest, "Cannot read from database: %s", err.Error())
		return
	}

	fantasyMapRecord := FantasyMap{
		DocumentID:      documentId,
		Filename:        filename,
		Source:          "User-Submitted",
		GcsURI:          "",
		PredictedBBoxes: "",
		ComputedBBoxes:  "",
		UserID:          "",
	}

	err = SaveImageToStorage(buf, &fantasyMapRecord, filename, c)
	if err != nil {
		log.Printf("Error saving file to GCS: %v\n", err.Error())
		c.String(http.StatusBadRequest, "Couldn't save file to storage %s", err.Error())
		return
	}

	imageData, err := GetPredictionFromFunction(buf, &fantasyMapRecord, c)
	if err != nil {
		c.String(http.StatusBadRequest, "Prediction error: %s", err.Error())
		return
	}

	// Store the map record in Firestore
	mapID, err := StoreMapRecord(buf, &fantasyMapRecord, c)
	if err != nil {
		log.Printf("Error recording %v map into database: %v", filename, err.Error())
		c.String(http.StatusBadRequest, "Record error: %s", err.Error())
		return
	}

	log.Printf("Stored new map: %v", mapID)

	// TODO: Replace this with redirect to another page
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded! Predictions: %s", filename, imageData))
}

func main() {
	router := gin.Default()

	// Serve static JS and CSS files
	router.Static("/styles", "./styles")
	router.Static("/scripts", "./scripts")
	router.Static("/images", "./images")

	// Load HTML files for templating
	router.LoadHTMLGlob("html/*")

	router.GET("/terms-of-service", func(c *gin.Context) {
		c.HTML(http.StatusOK, "terms-of-service.html", gin.H{})
	})

	router.GET("/privacy-policy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "privacy-policy.html", gin.H{})
	})

	router.GET("/submit-map", func(c *gin.Context) {
		c.HTML(http.StatusOK, "submit-map.html", gin.H{})
	})

	// Handle the form submission postback
	router.POST("/submit-map", postSubmitMaps)

	// Landing page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Set env vars
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "something-something-something.json")
	os.Setenv("PROJECT_ID", "my-project-id")

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

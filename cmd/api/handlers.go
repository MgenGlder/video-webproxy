package main

import (
	"fmt"
	"io/ioutil"
	"bytes"
	"time"
	"net/http"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"example.com/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) testStream(w http.ResponseWriter, r *http.Request) {
	k := "Video-20220208_000143-Meeting Recording.mp4"

	config := &aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("shared_config.txt", "default"),
	}



	sess := session.New(config)

	s3c := s3.New(sess, config)

	output, err := s3c.GetObject(&s3.GetObjectInput{Bucket: aws.String("kunle-dance-videos"), Key: aws.String(k)})
	if err != nil {
		fmt.Println(err)
		return
	}

	buff, buffErr := ioutil.ReadAll(output.Body)
	if buffErr != nil {
		fmt.Println(buffErr)
		return
	}

	reader := bytes.NewReader(buff)
	fmt.Println("Successfully read from s3")

	http.ServeContent(w, r, k, time.Now(), reader)
	return
}


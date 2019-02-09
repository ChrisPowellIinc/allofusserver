package tests

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func PostFile(filename, uploadURL, formFileName, AuthToken string) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile(formFileName, filename)
	if err != nil {
		log.Println("body writer : ", err)
		return nil, err
	}
	// open file handle
	fh, err := os.Open("../../testdata/" + filename)
	if err != nil {
		log.Println("file open : ", err)
		return nil, err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		log.Println("file copy : ", err)
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", uploadURL, bodyBuf)
	if err != nil {
		log.Println("Request :", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+AuthToken)
	req.Header.Add("Content-Type", contentType)
	rr, err := client.Do(req)
	// rr, err := http.Post(uploadURL, contentType, bodyBuf)
	if err != nil {
		log.Println("client DO: ", err)
		return nil, err
	}
	return rr, nil
}

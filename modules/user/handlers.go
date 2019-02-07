package user

import (
	"mime/multipart"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"path/filepath"
	
	"github.com/ChrisPowellIinc/allofusserver/models"
	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/render"
	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

// Get : Shows that the app is working
func (handler *Handler) Get(w http.ResponseWriter, r *http.Request) {

	type resStruct struct {
		Message string `json:"msg"`
	}

	res := resStruct{
		Message: "It works!",
	}

	render.JSON(w, r, res)

	return
}

func uploadFileToS3(s *session.Session, file multipart.File, fileName string, size int64, con *config.Config) error {
    // get the file size and read
    // the file content into a buffer
	buffer := make([]byte, size)
	file.Read(buffer)
	

    // config settings: this is where you choose the bucket,
    // filename, content-type and storage class of the file
    // you're uploading
    e, s3err := s3.New(s).PutObject(&s3.PutObjectInput{
        Bucket:               aws.String(con.Constants.S3Bucket),
        Key:                  aws.String(fileName),
        ACL:                  aws.String("public-read"),
        Body:                 bytes.NewReader(buffer),
        ContentLength:        aws.Int64(int64(size)),
        ContentType:          aws.String(http.DetectContentType(buffer)),
        ContentDisposition:   aws.String("attachment"),
        ServerSideEncryption: aws.String("AES256"),
        StorageClass:         aws.String("INTELLIGENT_TIERING"),
    })
	log.Println(e)
	log.Println(s3err)
    return s3err
}

// UploadProfilePic uploads a user's profile picture
func (handler *Handler) UploadProfilePic(w http.ResponseWriter, r *http.Request) {
	maxSize := int64(2048000) // allow only 2MB of file size

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		log.Println(err)
		models.HandleResponse(w, r, "Image too large", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("profile_picture")
	if err != nil {
		log.Println(err)
		models.HandleResponse(w, r, "Image not supplied", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// TODO:: Check for file type...
	tempFileName := "profile_pics/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)
	err = uploadFileToS3(handler.config.AwsSession, file, tempFileName, fileHeader.Size, handler.config)
	if err != nil {
		log.Println(err)
		models.HandleResponse(w, r, "An Error occured while uploading the image", http.StatusInternalServerError)
		return
	}

	models.HandleResponse(w, r, "Successfully Created File", http.StatusOK)
}

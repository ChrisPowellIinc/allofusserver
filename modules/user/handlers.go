package user

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	"github.com/ChrisPowellIinc/allofusserver/internal/jwt"
	"github.com/ChrisPowellIinc/allofusserver/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/render"
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

func uploadFileToS3(file multipart.File, fileName string, size int64, con *config.Config) error {
	// get the file size and read
	// the file content into a buffer
	buffer := make([]byte, size)
	file.Read(buffer)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(con.AwsSession).PutObject(&s3.PutObjectInput{
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
	return err
}

// UploadProfilePic uploads a user's profile picture
func (handler *Handler) UploadProfilePic(w http.ResponseWriter, r *http.Request) {
	userEmail, err := jwt.GetLoggedInUserEmail(r.Context())
	if err != nil {
		log.Println(err)
		models.HandleResponse(w, r, "Unable to retrieve authenticated user.", http.StatusUnauthorized)
		return
	}

	maxSize := int64(2048000) // allow only 2MB of file size

	err = r.ParseMultipartForm(maxSize)
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
	supportedFileTypes := map[string]bool{
		".png": true,
		".jpeg": true,
		".jpg": true,
	}
	filetype := filepath.Ext(fileHeader.Filename)
	if !supportedFileTypes[filetype] {
		log.Println(filetype)
		models.HandleResponse(w, r, "This image file type is not supported", http.StatusBadRequest)
		return
	}
	tempFileName := "profile_pics/" + bson.NewObjectId().Hex() + filetype
	err = uploadFileToS3(file, tempFileName, fileHeader.Size, handler.config)
	if err != nil {
		log.Println(err)
		models.HandleResponse(w, r, "An Error occured while uploading the image", http.StatusInternalServerError)
		return
	}

	imageURL := "https://s3.us-east-2.amazonaws.com/www.all-of.us/" + tempFileName

	err = handler.config.DB.C("user").Update(bson.M{"email": userEmail}, bson.M{"$set": bson.M{"image": imageURL}})
	if err != nil {
		log.Println(err)
		models.HandleResponse(w, r, "Unable to update user's email.", http.StatusInternalServerError)
		return
	}

	res := models.Response{}
	res.Message = "Successfully Created File"
	res.Status = http.StatusOK
	res.Data = map[string]interface{}{
		"imageurl": imageURL,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

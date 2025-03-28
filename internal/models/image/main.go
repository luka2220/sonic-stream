package image_model

import "mime/multipart"

// Model of a valid client request
type FileMetaData struct {
	Base             *multipart.FileHeader
	BaseExtention    string
	ConvertExtension string
}

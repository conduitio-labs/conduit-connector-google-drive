package googledrive

// Config contains shared config parameters, common to the source and
// destination. If you don't need shared parameters you can entirely remove this
// file.
type Config struct {
	// Google Project ID
	DriveProjectID string `json:"drive.projectId" validate:"required"`
	// Google Private Key ID
	DrivePrivateKeyID string `json:"drive.privateKeyId" validate:"required"`
	// Google Private Key
	DrivePrivateKey string `json:"drive.privateKey" validate:"required"`
	// Google Client Email
	DriveClientEmail string `json:"drive.clientEmail" validate:"required"`
	// Google Client ID
	DriveClientID string `json:"drive.clientId" validate:"required"`
	// Google Client Cert URL
	DriveClientCertUrl string `json:"drive.clientCertUrl" validate:"required"`
	// Folder to connect to
	DriveFolderID string `json:"drive.folderId" validate:"required"`
}

package googledrive

import (
	"encoding/json"
	"fmt"
)

// Config contains shared config parameters, common to the source and
// destination. If you don't need shared parameters you can entirely remove this
// file.
type Config struct {
	// The Google Cloud project ID associated with the service account.
	DriveProjectID string `json:"drive.projectId" validate:"required"`

	// The ID of the private key used to authenticate the service account.
	DrivePrivateKeyID string `json:"drive.privateKeyId" validate:"required"`

	// The private key (PEM-encoded) used to sign service account requests.
	DrivePrivateKey string `json:"drive.privateKey" validate:"required"`

	// The email address of the service account (e.g. my-service-account@project.iam.gserviceaccount.com).
	DriveClientEmail string `json:"drive.clientEmail" validate:"required"`

	// The OAuth2 client ID associated with the service account.
	DriveClientID string `json:"drive.clientId" validate:"required"`

	// The URL to the X.509 certificate for the service account, used to verify its identity.
	DriveClientCertURL string `json:"drive.clientCertUrl" validate:"required"`

	// The ID of the Google Drive folder where records will be uploaded.
	// This can be found in the folder's URL: https://drive.google.com/drive/folders/<folderId>
	DriveFolderID string `json:"drive.folderId" validate:"required"`
}

func (c *Config) BuildCredentials() ([]byte, error) {
	serviceAccountCredentials := map[string]interface{}{
		"type":                        "service_account",
		"project_id":                  c.DriveProjectID,
		"private_key_id":              c.DrivePrivateKeyID,
		"private_key":                 c.DrivePrivateKey,
		"client_email":                c.DriveClientEmail,
		"client_id":                   c.DriveClientID,
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":        c.DriveClientCertURL,
	}

	bytes, err := json.Marshal(serviceAccountCredentials)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal credentials: %w", err)
	}

	return bytes, nil
}

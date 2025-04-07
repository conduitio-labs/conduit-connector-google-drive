package googledrive

import (
	"encoding/json"
	"fmt"
)

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
	DriveClientCertURL string `json:"drive.clientCertUrl" validate:"required"`
	// Folder to connect to
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

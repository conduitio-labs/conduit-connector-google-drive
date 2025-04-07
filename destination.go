package googledrive

import (
	"bytes"
	"context"
	"fmt"

	"github.com/conduitio/conduit-commons/opencdc"
	sdk "github.com/conduitio/conduit-connector-sdk"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type Destination struct {
	sdk.UnimplementedDestination

	config  DestinationConfig
	service *drive.Service
}

type DestinationConfig struct {
	sdk.DefaultDestinationMiddleware
	// Config includes parameters that are the same in the source and destination.
	Config
}

func (s *DestinationConfig) Validate(context.Context) error {
	// Custom validation or parsing should be implemented here.
	return nil
}

func NewDestination() sdk.Destination {
	// Create Destination and wrap it in the default middleware.
	return sdk.DestinationWithMiddleware(&Destination{})
}

func (d *Destination) Config() sdk.DestinationConfig {
	return &d.config
}

func (d *Destination) Open(ctx context.Context) error {
	// Open is called after Configure to signal the plugin it can prepare to
	// start writing records. If needed, the plugin should open connections in
	// this function.
	creds, err := d.config.BuildCredentials()
	if err != nil {
		return fmt.Errorf("failed to build credentials: %w", err)
	}

	driveService, err := drive.NewService(ctx, option.WithCredentialsJSON(creds))
	if err != nil {
		return fmt.Errorf("failed to open connection: %w", err)
	}
	d.service = driveService

	return nil
}

func (d *Destination) Write(ctx context.Context, r []opencdc.Record) (int, error) {
	// Log the number of records
	sdk.Logger(ctx).Info().Int("records", len(r)).Msg("Starting file uploads...")

	// Initialize a counter to track the number of successfully uploaded records
	successfulUploads := 0

	// Loop through each record and upload it as a separate file
	for _, record := range r {
		fileData := record.Payload.After.Bytes()

		// Create a bytes buffer to hold the record data
		fileBuffer := bytes.NewBuffer(fileData)

		// Prepare the file metadata (include the folder ID in the Parents field)
		fileMetadata := &drive.File{
			Name:    fmt.Sprintf("%s.txt", record.Key.Bytes()), // Set the file name
			Parents: []string{d.config.DriveFolderID},          // Set the shared folder ID
		}

		// Upload the file directly from the bytes buffer
		uploadedFile, err := d.service.Files.Create(fileMetadata).Media(fileBuffer).Do()
		if err != nil {
			return successfulUploads, fmt.Errorf("unable to upload file: %w", err)
		}

		// Log the uploaded file's ID
		sdk.Logger(ctx).Info().Msgf("File uploaded successfully! File ID: %s\n", uploadedFile.Id)

		// Increment the successful uploads counter
		successfulUploads++
	}

	return successfulUploads, nil
}

func (d *Destination) Teardown(_ context.Context) error {
	// Teardown signals to the plugin that all records were written and there
	// will be no more calls to any other function. After Teardown returns, the
	// plugin should be ready for a graceful shutdown.
	return nil
}

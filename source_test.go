package googledrive_test

import (
	"context"
	"testing"

	googledrive "github.com/conduitio-labs/conduit-connector-google-drive"
	"github.com/matryer/is"
)

func TestTeardownSource_NoOpen(t *testing.T) {
	is := is.New(t)
	con := googledrive.NewSource()
	err := con.Teardown(context.Background())
	is.NoErr(err)
}

package googledrive_test

import (
	"context"
	"testing"

	"github.com/matryer/is"
	googledrive "github.com/repository/conduit-connector-google-drive"
)

func TestTeardownSource_NoOpen(t *testing.T) {
	is := is.New(t)
	con := googledrive.NewSource()
	err := con.Teardown(context.Background())
	is.NoErr(err)
}

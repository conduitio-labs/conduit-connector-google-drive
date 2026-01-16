package googledrive_test

import (
	"context"
	"testing"

	googledrive "github.com/conduitio-labs/conduit-connector-google-drive"
	"github.com/matryer/is"
)

func TestTeardown_NoOpen(t *testing.T) {
	is := is.New(t)
	con := googledrive.NewDestination()
	err := con.Teardown(context.Background())
	is.NoErr(err)
}

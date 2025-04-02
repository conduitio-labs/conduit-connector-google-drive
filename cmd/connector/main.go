package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
	googledrive "github.com/repository/conduit-connector-google-drive"
)

func main() {
	sdk.Serve(googledrive.Connector)
}

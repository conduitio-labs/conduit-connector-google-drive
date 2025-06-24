package main

import (
	googledrive "github.com/conduitio-labs/conduit-connector-google-drive"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func main() {
	sdk.Serve(googledrive.Connector)
}

package main

import (
	"flag"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	rsMgr "google.golang.org/api/cloudresourcemanager/v2"
)

var (
	folderID = flag.String("folder_id", "", "your folder ID")
)

// Copied from https://github.com/google/google-api-go-client
func main() {
	flag.Parse()

	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.Background()

	client, err := google.DefaultClient(ctx, rsMgr.CloudPlatformScope)
	if err != nil {
		log.Fatalf("error getting client, err: %v", err)
	}
	rsMgrSvc, err := rsMgr.New(client)
	if err != nil {
		log.Fatalf("error creating service, err: %v", err)
	}

	fs := rsMgr.NewFoldersService(rsMgrSvc)

	folder, err := fs.Get("folders/" + *folderID).Do()
	if err != nil {
		log.Fatalf("error listing folders, err: %v", err)
	}

	log.Println(folder.DisplayName)
}

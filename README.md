# gcp-res-mgr

## Usage

1. Install go packages.

```
glide install
```

1. Enable Cloud Resource Manager API in your project: https://console.developers.google.com/apis/api/cloudresourcemanager.googleapis.com/overview

1. Create a service account. Download its JSON key, and then export the following environment variable.

```
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/key
```

1. Grant your service account "Resource Manager > Folder Viewer" access.

1. Run!

```
go run main/main.go -folder_id=<your_folder_ID>
```

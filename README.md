# gcp-res-mgr

## Usage

- Install go packages.
```
glide install
```

- Enable Cloud Resource Manager API in your project: https://console.developers.google.com/apis/api/cloudresourcemanager.googleapis.com/overview

- Create a service account. Download its JSON key, and then export the following environment variable.
```
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/key
```

- Grant your service account "Resource Manager > Folder Viewer" access.

- Run!
```
go run main/main.go -folder_id=<your_folder_ID>
```

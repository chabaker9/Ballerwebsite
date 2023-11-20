## Local testing
podman build -t website .
podman run -d -p 8080:8080 --name website website
podman logs website

podman stop website
podman rm website

podman run -it --rm -p 8080:8080 --name website website bash
podman exec -it website bash

## delete all stopped containers
podman rm $(podman ps -a -q -f status=exited)

## deploy to gcp

### Login
gcloud auth login

# update github 
git add .;git commit -m 'update';git push

### Config
$project = "extended-method-292015"
$appName = "ballerwebsite"
$image = "gcr.io/$project/${appName}:0.4.0"
gcloud config set project $project

### Build and deploy
gcloud builds submit --tag $image
gcloud run deploy $appName --image $image --platform managed --region us-central1 --allow-unauthenticated

## clean up
gcloud container images delete $image --force-delete-tags
gcloud run services delete $appName --platform managed --region us-central1


## go
go run src/main.go
go run main.go

go mod init src/main.go
go mod tidy
go build src/main.go

dep ensure
CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main main.go
sudo docker build . -t clip-go-image
sudo docker save clip-go-image -o docker_image
sudo rm docker-image

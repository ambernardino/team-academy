dep ensure
CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main main.go
sudo /root/bin/docker-compose down
sudo docker build . -t clip-go-image
sudo docker save clip-go-image -o docker_image
sudo /root/bin/docker-compose up -d
sudo rm docker_image


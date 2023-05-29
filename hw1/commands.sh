go build -o http-server ./service/main.go
docker build -t hw1:v1 .
docker build -f dockerfile -t hw1:v1 .
docker build -f dockerfile_build -t hw1:v2 .
docker build -f dockerfile_3 -t hw1:v3 .

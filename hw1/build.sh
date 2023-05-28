go build -o srv ./service/main.go
docker build -t hw1:v1 .
docker build -f dockerfile -t hw1:v1 .
docker build -f dockerfilebuild -t hw1:v2 .

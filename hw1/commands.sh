go build -o http-server ./service/main.go
docker build -t hw1:v1 .
docker build -f dockerfile -t hw1:v1 .
docker build -f dockerfile_build -t hw1:v2 .
docker build -f dockerfile_3 -t hw1:v3 .

docker run hw1:v4 -d -p 8000:8000

docker build -f dockerfile_1 -t hw1:v4 .

docker tag local-image:tagname new-repo:tagname

docker push alekseybondarev/otus:latest

docker tag hw1:v4 alekseybondarev/otus
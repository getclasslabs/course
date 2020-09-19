#!/bin/bash

echo "Compiling the API"
docker run -it --rm -v "$(pwd)":/go -e GOPATH= golang:1.14 sh -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/course/"

rm ./docker/course
mv ./course ./docker/
cp ./docker-config.yaml ./docker/config.yaml

docker build -t getclass/course:"$1" docker/

docker push getclass/course:"$1"

if [[ ! $(docker service ls | grep gcl_course) = "" ]]; then
  docker service update gcl_course --image getclass/course:$1
else
  docker stack deploy -c docker-compose.yaml gcl
fi
#!/bin/bash
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)

docker rmi aloxc/gappweb
cd /gapp/gappweb
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gappweb
docker build -t aloxc/gappweb /gapp/gappweb/

docker rmi aloxc/gappuser
cd /gapp/gappuser
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gappuser
docker build -t aloxc/gappuser /gapp/gappuser/

docker rmi aloxc/gapporder
cd /gapp/gapporder
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gapporder
docker build -t aloxc/gapporder /gapp/gapporder/
#!/usr/bin/env bash

docker pull nexus:9820/rpmbuilder:10.00.00
docker pull nexus:9820/apkbuilder:3.17.00
docker tag nexus:9820/rpmbuilder:10.00.00 rpmbuilder
docker tag nexus:9820/apkbuilder:3.17.00 apkbuilder
docker rmi nexus:9820/apkbuilder:3.17.00

docker images


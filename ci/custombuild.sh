#!/bin/bash
NEW_IMAGE_TAG=latest
echo "Custom build script for cyderes continueous intergeration"

cd ..

echo "Syntax check"
go vet *.go

echo "Building new application image"
docker image build . -t datatx:latest

echo "Creating new tag from latest application"
docker tag datatx karmegamp/datatx:${NEW_IMAGE_TAG}
echo "Build and tag creation done"

echo "Uploading new container image to docker hub"
docker push karmegamp/datatx:$NEW_IMAGE_TAG
echo "Uploaded new container image datatx:${NEW_IMAGE_TAG} to docker hub"
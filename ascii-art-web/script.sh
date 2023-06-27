#!/bin/bash

# Nom de l'image
IMAGE_NAME="ascii-art-web"

# Tag de l'image
TAG="v1.0"

# Construire l'image à partir du Dockerfile
docker build -t "${IMAGE_NAME}:${TAG}" .

#Relancer le container
docker run -dp 8080:8080 "${IMAGE_NAME}:${TAG}"

# Vérifier si la construction de l'image a réussi
if [ $? -eq 0 ]; then
    echo "Construction de l'image ${IMAGE_NAME}:${TAG} terminée avec succès."
else
    echo "La construction de l'image a échoué."
fi

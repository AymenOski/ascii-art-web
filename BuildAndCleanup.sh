#!/bin/bash

# building the image
docker build -t my_go_image .

# running the container from the image
docker run -d --name my_go_app -p 8080:8080 my_go_image

# Print out the status of the system
echo "Cleaning up Docker environment..."

# Remove all unused volumes
docker volume prune -f

# Remove all dangling images (images without tags)
docker image prune -af

# Cleanup cache (remove all intermediate layers that are no longer needed)
docker builder prune -af

echo "Docker cleanup completed."

echo "The container is running at--> http://localhost:8080"
# From base image
FROM nginx:latest

# Update apt-get
RUN apt-get update

# Install vim and say yes to all questions
RUN apt-get install -y vim

# Copy files from host to container
## First argument: path to file or folder on host relative to the current directory
## Second argument: path to file or folder in the container relative to root
COPY /modules/docker/html /usr/share/nginx/html

# Specify working directory (base directory when you access the container)
WORKDIR /app
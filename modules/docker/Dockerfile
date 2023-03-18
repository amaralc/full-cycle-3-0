# From base image
FROM nginx:latest

# Update apt-get
RUN apt-get update

# Install vim and say yes to all questions
RUN apt-get install -y vim

# Copy files from host to container
## First argument: path to file on host relative to the dockerfile
## Second argument: path to file on container relative to root
COPY html /usr/share/nginx/html

# Specify working directory (base directory when you access the container)
WORKDIR /app
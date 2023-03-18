# From base image
FROM ubuntu:latest

# Specify which executable should be called once the container starts (first parameter) and what arguments to pass to that executable (following parameters) ["executable", ..."parameters"]
## This command can be substituted once you run a new container, by another executable and other parameters
## Ex.: docker run --rm username/nginx-with-vim:latest echo "Hello Full Cycle!"
CMD [ "echo", "Hello World" ]
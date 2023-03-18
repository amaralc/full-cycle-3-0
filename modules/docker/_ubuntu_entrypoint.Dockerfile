# From base image
FROM ubuntu:latest

# Specify which executable should be called once the container starts (first parameter) and what arguments to pass to that executable (following parameters) ["executable", ..."parameters"]
## This command will always be executed when you run a new container
ENTRYPOINT [ "echo", "Hello" ]

# Specify a parameter that can be overwritten
CMD [ "World" ]
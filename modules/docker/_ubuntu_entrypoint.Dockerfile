# From base image
FROM ubuntu:latest

# Copy entrypoint to root folder in the container
COPY /modules/docker/_ubuntu_entrypoint.sh /_ubuntu_entrypoint.sh

# Add executable permission to the entrypoint script
## If you don't add this line you will see the following error message:
### docker: Error response from daemon: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: exec: "/_ubuntu_entrypoint.sh": permission denied: unknown.
RUN chmod +x /_ubuntu_entrypoint.sh

# Specify which executable should be called once the container starts (first parameter) and what arguments to pass to that executable (following parameters) ["executable", ..."parameters"]
## This command will always be executed when you run a new container
ENTRYPOINT [ "/_ubuntu_entrypoint.sh" ]

# # Specify a parameter that can be overwritten
CMD [ "echo", "World" ]
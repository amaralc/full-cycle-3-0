# Working with Images

[<- Index](../../README.md)

</br>

### Key notes

-

## Understanding Images and DockerHub

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6705

- Docker official images are hosted in https://hub.docker.com/
- It is common to upload two tags once a product is released: "latest" and a fixed version (ex.: 3.9.0);
- The latest version is usually overwritten overtime;
- (terminal) Pull an ubuntu image from the container registry: `docker pull ubuntu`
- (terminal) List docker images in my host machine: `docker images`
- (terminal) Remove an image from my host machine: `docker rmi ubuntu:latest`

## Creating first image with Dockerfile

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6706

- The Dockerfile is a recipe for creating a new image;
- Install the Docker VS Code extension to view code highlights in your Dockerfile;
- Create a Dockerfile adding vim to a nginx:latest image;

```Dockerfile
# Dockerfile

# From base image
FROM nginx:latest

# Update apt-get
RUN apt-get update

# Install vim and say yes to all questions
RUN apt-get install -y vim
```

- Build an image from that file:

```bash
docker build -t username/nginx-with-vim:latest path/to/folder/with/dockerfile
```

Example:

```bash
docker build -t username/nginx-with-vim:latest "$(pwd)"/modules/docker
```

- If the name of the file is not "Dockerfile", then specify it as:

```bash
docker build . -t username/nginx-with-vim:latest -f "$(pwd)"/modules/docker/my-file-name.Dockerfile
```

- (terminal) List all images in your computer: `docker images`
- Verify that there a new image "username/nginx-with-vim:latest" was created;

```bash
REPOSITORY                                  TAG                  IMAGE ID       CREATED         SIZE
username/nginx-with-vim                     latest               46b8018b5700   3 minutes ago   196MB
```

- (terminal) Create a new container from that image: `docker run -it username/nginx-with-vim:latest bash`
- (terminal) Inside the container verify if vim is installed: `vim --version`

## Going Further with Dockerfile

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6707

- Enhance your dockerfile with a custom index.html file:

```Dockerfile
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
```

- Notice that during the image build, previous steps were reused from cache;

## Entrypoint vs Command

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6708

The **CMD** specifies an executable and its parameters that should execute once the container starts but that can be overwritten once the container is run;

The **ENTRYPOINT** specifies a command and parameters that are always executed once the container starts;

---

</br>

### CMD

- (terminal) Create a new dockerfile: `touch _ubuntu_cmd.Dockerfile`
- (\_ubuntu_cmd.Dockerfile) Add content to the file:

```Dockerfile
# From base image
FROM ubuntu:latest

# Specify which executable should be called once the container starts (first parameter) and what arguments to pass to that executable (following parameters) ["executable", ..."parameters"]
CMD [ "echo", "Hello World" ]
```

- (terminal) Build the image from that file:

```bash
docker build . -t username/nginx-with-vim:latest -f "$(pwd)"/modules/docker/_ubuntu_cmd.Dockerfile
```

- (terminal) Run a container from that image: `docker run --rm username/nginx-with-vim-cmd:latest`

- (terminal) Verify that the container was run and that "Hello World" was logged to the terminal;
- (terminal) Now run another container but substitute the executable and parameters:

```bash
docker run --rm username/nginx-with-vim-cmd:latest echo "Hello Full Cycle!"
```

### ENTRYPOINT

- (terminal) Create a new dockerfile: `touch _ubuntu_entrypoint.Dockerfile`
- (\_ubuntu_entrypoint.Dockerfile) Add content to the file:

```Dockerfile
# From base image
FROM ubuntu:latest

# Specify which executable should be called once the container starts (first parameter) and what arguments to pass to that executable (following parameters) ["executable", ..."parameters"]
## This command will always be executed when you run a new container
ENTRYPOINT [ "echo", "Hello" ]

# Specify a parameter that can be overwritten
CMD [ "World" ]
```

- (terminal) Build the image from that file:

```bash
docker build . -t username/nginx-with-vim-entrypoint:latest -f "$(pwd)"/modules/docker/_ubuntu_entrypoint.Dockerfile
```

- (terminal) Run a container from that image: `docker run --rm username/nginx-with-vim-entrypoint:latest`

- (terminal) Verify that the container was run and that "Hello World" was logged to the terminal;
- (terminal) Now run another container but add parameter to overwrite the "World" string:

```bash
docker run --rm username/nginx-with-vim-entrypoint:latest "Full Cycle!"
```

- (terminal) Verify that the message logged to the terminal was "Hello Full Cycle!"

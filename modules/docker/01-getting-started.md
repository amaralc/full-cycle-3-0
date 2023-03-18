# Getting Started

### Key notes

-

## Hello World

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6698

- (terminal) List running containers: `docker ps`
- (terminal) Run first container: `docker run hello-world:latest`

  - "hello world": image name in container registry;
  - "latest: image tag;

  - Images have an entrypoint that is an executable;
  - This "hello-world" image have an entrypoint that only prints a message to the screen:

  ```

  Unable to find image 'hello-world:latest' locally
  latest: Pulling from library/hello-world
  2db29710123e: Pull complete
  Digest: sha256:ffb13da98453e0f04d33a6eee5bb8e46ee50d08ebe17735fc0779d0349e889e9
  Status: Downloaded newer image for hello-world:latest

  Hello from Docker!
  This message shows that your installation appears to be working correctly.

  To generate this message, Docker took the following steps:

  1.  The Docker client contacted the Docker daemon.

  2.  The Docker daemon pulled the "hello-world" image from the Docker Hub.
      (amd64)

  3.  The Docker daemon created a new container from that image which runs the
      executable that produces the output you are currently reading.
  4.  The Docker daemon streamed that output to the Docker client, which sent it
      to your terminal.

  To try something more ambitious, you can run an Ubuntu container with:
  $ docker run -it ubuntu bash

  Share images, automate workflows, and more with a free Docker ID:
  https://hub.docker.com/

  For more examples and ideas, visit:
  https://docs.docker.com/get-started/

  ```

  - After printing the message the container died (stopped running), since the entrypoint stopped running;
  - If we need the container to keep running, than that process that is called by the entrypoint must be kept running;

  </br>

- (terminal) List all containers (even those that are not running): `docker ps -a`

```

CONTAINER ID   IMAGE                COMMAND    CREATED         STATUS                     PORTS     NAMES
6965660f0323   hello-world:latest   "/hello"   6 seconds ago   Exited (0) 5 seconds ago             quizzical_banach

```

- The "COMMAND" column shows what command was executed once the container was instantiated;
- The "NAMES" column show the automatically generated name for the container;

## Executing Ubuntu

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6699

- (terminal) Run ubuntu container: `docker run -it ubuntu bash` or `docker run -i -t ubuntu bash`

```
  - The "-i" parameter means interactive mode, attaching your terminal to the container;
  - The "-t" parameter means "tty" which stands for "teletypewriter" which lets you run remote commands from your terminal to that container;
```

- (terminal) In other terminal list the containers that are executing: `docker ps`
- (terminal) Verify that a container with ubuntu image is running and that "bash" is the process that is keeping it alive;

```
A container is a process in the host machine. If the inner process that keeps it running dies, the container die.
```

- (terminal) Run a container and remove it after you exit: `docker run -it --rm ubuntu:latest bash`

## Publishing ports with Nginx

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6700

- (terminal) Run nginx image: `docker run nginx`

```
Exposing a port is not enough to access the container from the host machine, which is in other network
```

- (terminal) Run nginx image and redirect port 80 from host machine to port 80 from that container: `docker run -p 80:80 nginx`
- (terminal) Exit the nginx process;
- (terminal) Run nginx image in detached mode: `docker run -d -p 80:80 nginx`

## Removing containers

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6701

- (terminal) List running containers: `docker ps`

```
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS                               NAMES
daa3bd8bc96b   nginx     "/docker-entrypoint.…"   6 minutes ago   Up 6 minutes   0.0.0.0:80->80/tcp, :::80->80/tcp   stupefied_bohr
```

- (terminal) Stop running container by id: `docker stop daa3bd8bc96b`
- (terminal) Remove container by id after stopping it: `docker rm daa3bd8bc96b`
- (terminal) Force removal of a running container by id: `docker rm daa3bd8bc96b -f`
- (terminal) Force removal of a running container by name: `docker rm mystifying_thompson -f`


# Getting Started

[<- Index](../../README.md)

</br>

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

If you want to stop all containers you can use the following command:

- (terminal) Stop all containers: `docker stop $(docker ps -q)`;
- (terminal) To remove all containers that are not running: `docker rm $(docker ps -a -q)`
- (terminal) To remove all containers (running or not): `docker rm $(docker ps -a -q) -f`

## Accessing and changing files in a container

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6702

- (terminal) Run a named container: `docker run -d -p 8080:80 --name my-container nginx`
- (terminal) Execute the "ls" bash command in a running container and get out from it: `docker exec my-container ls`
- (terminal) Execute the "ls" bash command in a running container and stay connected in tty: `docker exec -it my-container bash`
- (terminal) Update apt-get: `apt-get update`
- (terminal) Install vim: `apt-get install vim`
- (terminal) Navigate to folder within the container: `cd /usr/share/nginx/html/`
- (terminal) Open index.html with vim: `vim index.html`;
- (terminal) Enter edit mode by pressing "i";
- (terminal) Navigate to "Welcome to Nginx!" and edit that sentence to "Welcome to Full Cycle!";
- (terminal) Exit edit mode by pressing Ctrl + C;
- (terminal) Save and exit by pressing ":wq" + Enter;
- (terminal) If you choose to exit without saving: ":q!" + Enter;
- (terminal) Exit the container: `exit`
- (browser) Access localhost:8080;
- (browser) Verify that your changes are live;
- (terminal) Remove the container: `docker rm my-container -f`

## Getting started with bind mounts

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6703

- (terminal) Create a local folder that you want to use as a volume for your container: `mkdir modules/docker/html`;
- (terminal) Add an index.html file to it: `touch modules/docker/html.index.html`;
- (VS Code) Add content to that html page: `<h1>Full Cycle<h1>`;
- (terminal) Run a container mounting a volume from your host machine in a container: `docker run -d --name my-container -p 8080:80 -v $(pwd)/modules/docker/html:/usr/share/nginx/html nginx:latest`

```
The -v option specifies what path in my machine should be bound to what path in the container:

"-v path/in/my/machine:path/in/the/container"
```

- (browser) Navigate to locahost:8080;
- (browser) Verify that the content on you local volume is displayed in the nginx page;

```
Note:

This is how you can work in your development environment using Docker. You can mount your src folder content as a volume of your container and then run it. By doing that, you can change your source code, and then access the changes through the container.

```

Now a days, instead of using -v, the most popular command is the "mount command. It lets it explicit the type, source and target paths:

- (terminal) Run a container, binding a volume using the "mount" attribute:

```
docker run -d --name my-container -p 8080:80 --mount type=bind,source="$(pwd)"/modules/docker/html,target=/usr/share/nginx/html nginx:latest
```

A big difference between the "-v" and the "--mount" options is that:

- "-v" it creates a source folder in your local machine if the specified path didn't exist before;
- "--mount" do not create a source folder if it doesn't exist;

## Working with volumes

- https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6704

</br>

- Bind mounts as seen previously, are not the same concept as volumes;
- (terminal) Create a named volume: `docker volume create my-volume`;
- (terminal) List your volumes: `docker volume ls`;
- (terminal) Inspect your volume: `docker volume inspect my-volume`;
- (terminal) Run a container connected with your volume:

```
docker run --name my-container -d --mount type=volume,source=my-volume,target=/app nginx
```

- (terminal) Execute your container in interactive mode and with teletypewriter: `docker exec -it my-container bash`
- (terminal) Create a new file within that container: `touch app/index.html`;
- (terminal) Exit your container: `exit`
- (terminal) Create a new container and mount the same volume to it:

```
docker run --name my-second-container -d --mount type=volume,source=my-volume,target=/app nginx
```

- (terminal) Execute your second container with -it: `docker exec -it my-container bash`
- (terminal) Navigate to /app: `cd /app`;
- (terminal) List the content of that folder: `ls`
- (terminal) Verify that that folder contains the index.html file you have created within the first container;

Another way of creating a new container, binding it to a volume is by using "-v:

- (terminal) Create a new container and mount the same volume to it:

```
docker run --name my-third-container -d -v my-volume:/app nginx
```

Now lets stop and remove all containers and volumes:

- (terminal) Stop and remove all containers: `docker rm $(docker ps -a -q) -f`;
- (terminal) Remove all volumes that have "my-volume" substring in their name: `docker volume rm $(docker volume ls -qf name=my-volume)`
- (terminal) Force remove all volumes that exist but are not connected to any container: `docker volume rm $(docker volume ls -qf dangling=true)`;

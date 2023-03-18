# Full Cycle 3.0

[Full Cycle](https://curso.fullcycle.com.br/page/lancamento/) course notes

</br>

- [DevOps](#devops)

  - [Docker](#docker) (40%)
  - [Patterns and Advanced Techniques with Git and Github] (0%)
  - [Continuous Integration] (0%)
  - [Kubernetes] (0%)
  - [Service Mesh with Istio] (0%)
  - [API Gateway with Kong and Kubernetes] (0%)
  - [Observability] (0%)
  - [Introduction to OpenTelemetry] (0%)
  - [Terraform] (0%)
  - [Ansible] (0%)
  - [GitOps] (0%)
  - [Deploy in Cloud Providers] (0%)

</br>

- [Architecture and Development](#architecture-and-development)
  - [Fundamentals of Software Architecture] (0%)
  - [Communication Between Systems] (0%)
  - [Solid Express] (0%)
  - [Domain Driven Design] (8%)
  - [DDD: Tactical Modeling and Patterns] (0%)
  - [Practical Event Storming] (0%)
  - [Hexagonal Architecture](#hexagonal-architecture) (28%)
  - [Clean Architecture](#clean-architecture) (25%)
  - [Monolithic Systems] (0%)
  - [Microservices Architecture] (0%)
  - [EDA - Event Driven Architecture] (0%)
  - [API Gateway] (0%)
  - [RabbitMQ] (0%)
  - [Apache Kafka] (0%)
  - [Authentication and KeyCloak] (0%)
  - [Practical Project Architecture - Codeflix] (0%)
  - [Practical Project - TypeScript (Back-end)] (0%)
  - [Practical Project - React (Front-end)] (0%)
  - [Video Encoder Microservice - Go Lang] (0%)

</br>

---

## DevOps

### Docker

- [Getting Started](./modules/docker/01-getting-started.md)
  - [Hello World](./modules/docker/01-getting-started.md#hello-world)
  - [Executing Ubuntu](./modules/docker/01-getting-started.md#executing-ubuntu)
  - [Publishing ports with nginx](./modules/docker/01-getting-started.md#publishing-ports-with-nginx)
  - [Removing containers](./modules/docker/01-getting-started.md#removing-containers)
  - [Accessing and changing files in a container](./modules/docker/01-getting-started.md#accessing-and-changing-files-in-a-container)
  - [Getting started with bind mounts](./modules/docker/01-getting-started.md#getting-started-with-bind-mounts)
  - [Working with volumes](./modules/docker/01-getting-started.md#working-with-volumes)
- [Working With Images](./modules/docker/02-working-with-images.md)
  - [Understanding Images and DockerHub](./modules/docker/02-working-with-images.md#understanding-images-and-dockerhub)
  - [Creating first image with Dockerfile](./modules/docker/02-working-with-images.md#creating-first-image-with-dockerfile)
  - [Going Further with Dockerfile](./modules/docker/02-working-with-images.md#going-further-with-dockerfile)
  - [Entrypoint vs Command](./modules/docker/02-working-with-images.md#entrypoint-vs-command)
  - [Docker Entrypoint exec](./modules/docker/02-working-with-images.md#docker-entrypoint-exec)
  - [Publishing image to DockerHub](./modules/docker/02-working-with-images.md#publishing-image-to-dockerhub)

</br>

---

## Architecture and Development

### Hexagonal Architecture

- [Introduction to Hexagonal Architecture](./modules/hexagonal-architecture/01-introduction-to-hexagonal-architecture.md)
  - [Introduction](./modules/hexagonal-architecture/01-introduction-to-hexagonal-architecture.md#introduction)
  - [Project Cycle](./modules/hexagonal-architecture/01-introduction-to-hexagonal-architecture.md#project-cycle)
  - [The Unsustainable Path](./modules/hexagonal-architecture/01-introduction-to-hexagonal-architecture.md#the-unsustainable-path)
  - [Important Notes](./modules/hexagonal-architecture/01-introduction-to-hexagonal-architecture.md#important-notes)
  - [Software Design vs. Architecture](./modules/hexagonal-architecture/01-introduction-to-hexagonal-architecture.md#software-design-vs-architecture)
  - [Presenting Hexagonal Architecture (Ports and Adapters)](./modules/hexagonal-architecture/01-introduction-to-hexagonal-architecture.md#presenting-hexagonal-architecture-ports-and-adapters)
  - [Architecture Dynamics](./modules/hexagonal-architecture/01-introduction-to-hexagonal-architecture.md#architecture-dynamics)
  - [Hexagonal vs Clean vs Onion](./modules/hexagonal-architecture/01-introduction-to-hexagonal-architecture.md#hexagonal-vs-clean-vs-onion)

### Clean Architecture

- [Basic Concepts](./modules/clean-architecture/01-basic-concepts.md)
  - [Introduction](./modules/clean-architecture/01-basic-concepts.md#introduction)
  - [Clean Architecture Origin](./modules/clean-architecture/01-basic-concepts.md#clean-architecture-origin)
  - [Important topics on architecture](./modules/clean-architecture/01-basic-concepts.md#important-topics-on-architecture)
  - [Keep options open](./modules/clean-architecture/01-basic-concepts.md#keep-options-open)
  - [Use cases](./modules/clean-architecture/01-basic-concepts.md#use-cases)
  - [Use cases flow](./modules/clean-architecture/01-basic-concepts.md#use-cases-flow)
  - [Architectural Limits](./modules/clean-architecture/01-basic-concepts.md#architectural-limits)
  - [Input vs Output](./modules/clean-architecture/01-basic-concepts.md#input-vs-output)
  - [Data Transfer Object (DTO)](./modules/clean-architecture/01-basic-concepts.md#data-transfer-object-dto)
  - [Presenters](./modules/clean-architecture/01-basic-concepts.md/#presenters)
  - [Entities vs DDD](./modules/clean-architecture/01-basic-concepts.md#entities-vs-ddd)

# Working with Images

[<- Index](../../README.md)

</br>

# Publishing image to DockerHub

https://plataforma.fullcycle.com.br/courses/242/168/110/conteudos?capitulo=110&conteudo=6710

- (terminal) Create a new dockerfile: `touch _nginx_fullcycle.Dockerfile`;
- (\_nginx_fullcycle.Dockerfile) Add the following content to the file:

```Dockerfile
FROM nginx:latest

COPY html /usr/share/nginx/html

ENTRYPOINT [ "/docker-entrypoint.sh" ]

CMD [ "nginx", "-g", "daemon off;" ]
```

- (terminal) Build the image and naming it using your DockerHub username:

```bash
docker build . -t amaralc/nginx-full-cycle:latest -f "$(pwd)"/modules/docker/_nginx_fullcycle.Dockerfile
```

- (terminal) Run a container with that image:

```bash
docker run --rm -d -p 8080:80 amaralc/nginx-full-cycle:latest
```

### Publish image to DockerHub

If you still do not have an account in DockerHub, go ahead and create one at https://hub.docker.com/

After creating your account, go back to your terminal and:

- (terminal) Login to docker registry: `docker login`
- (terminal) Fill in your username and press enter;
- (terminal) Fill in your password and press enter;
- (terminal) Push your image to docker hub: `docker push amaralc/nginx-full-cycle:latest`

- (browser) Now visit your repository at `https://hub.docker.com/repository/docker/your-username/your-image-name` and verify that it was published. In our case: https://hub.docker.com/repository/docker/amaralc/nginx-full-cycle

_IMPORTANT_: in free mode, docker will automatically delete your images after some time of inactivity (now downloads);

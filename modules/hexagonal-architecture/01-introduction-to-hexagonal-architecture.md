# Hexagonal Architecture

[<- Index](../../README.md)

</br>

### Key notes

- Set clear limits between business complexity and technical complexity;

## Introduction

https://plataforma.fullcycle.com.br/courses/242/168/123/conteudos?capitulo=123&conteudo=6963

- Understand clearly the differences between:
  - business rules complexity;
  - technical complexity;
- Protect the business against technical complexity;
- Set clear limits so that the technical complexity does not invade the business;
- Technical complexity is used only to access the business model;
- Technical complexity can be replaced over time;

## Project Cycle

https://plataforma.fullcycle.com.br/courses/242/168/123/conteudos?capitulo=123&conteudo=6964

- Sustainable software development;
- Avoid developing softwares that depend on frameworks;
- Architecture has to do with the future of the software;
- Avoid thinking of databases before understanding the business rules;

### The unsustainable path

Phase 1:

- Set database;
- Set validations;
- Set web server;
- Create controllers;
- Create views;
- Authentication;

Phase 2:

- Business rules;
- Create APIs;
- Consume APIs;
- Authorization;
- Reports;
- Logs;

Phase 3:

- More accesses;
- Upgrade hardware;
- Add cache;
- Consume external APIs;
- Deal with external API rules;

Phase 4:

- More access;
- More upgrades;
- DB reports;
- Generate commands (CLI) to generate reports;
- V2 of API;

Phase 5:

- Horizontal scale;
- Sessions need to be distributed;
- Cloud uploads with authentication;
- Refactoring;
- Autoscaling;
- CI/CD;

Phase 6:

- GraphQL;
- Constant bugs;
- Distributed logs;
  - Machine with auto-scaling does not exist;
  - Logs are lost;
  - Cloud logs;
- Integrate with CRM;
- Migrate to React;

Phase 7:

- CRM with inconsistencies;
- Containerize application;
- Update CI/CD;
- Memory management in containers;
- Container logs;
- Legacy system;

Phase 8:

- Migrate to microservices;
- Shared database;
- Issues with tracing;
- Double latency;
- Elevated cost;

Phase 9:

- Kubernetes;
- Change CI/CD;
- Queues and messages;
- Lost messages;
  - No [dead letter kill](https://www.kai-waehner.de/blog/2022/05/30/error-handling-via-dead-letter-queue-in-apache-kafka/);
- Consultants;

Phase 10:

- Imagine;

## Important notes

https://plataforma.fullcycle.com.br/courses/242/168/123/conteudos?capitulo=123&conteudo=6965

### Main issues

- Technical debts;
- Missing future vision for the software;
- Missing well-defined limits;
- Exchange components;
- Scale issues;
- Frequent optimizations;
- Prepared to change;
- Anti-corruption layers;
- Abrupt changes;

### Questions

- Is it painful for the developer to develop?
- Could be avoided?
- Is the project returning the investments (ROI)?
- The customer relation is healthy?
- Will the customer being prejudiced by architectural change?
- When dit things went down?
  - Day one;
  - One day after another;
- If you are new to the team, would you judge the previous developers?

## Software design vs Architecture

https://plataforma.fullcycle.com.br/courses/242/168/123/conteudos?capitulo=123&conteudo=6966

"Atividades relacionadas à arquitetura de software são sempre de design. Entretanto, nem todas as atividades de design são sobre arquitetura. O objetivo primário da arquitetura de software é garantir que os atributos de qualidade, restrições de alto nível e os objetivos do negócio, sejam atendidos pelo sistema. Qualquer decisão de design que não tenha relação com este objetivo, não é decisão arquitetural. Todas as decisões de design para um componente que não sejam ‘visíveis’ fora dele, geralmente também não são."

Source: [Elemar Jr.](https://eximia.co/quais-sao-as-diferencas-entre-arquitetura-e-design-de-software/)

## Presenting Hexagonal Architecture (Ports and Adapters)

https://plataforma.fullcycle.com.br/courses/242/168/123/conteudos?capitulo=123&conteudo=6967

"Allow an application to equally be driven by users, programs, automated test or batch scripts, and to be developed and tested in isolation from its eventual run-time devices and databases."

Source: [Cockburn, A.](https://alistair.cockburn.us/hexagonal-architecture/)

## Architecture Dynamics

https://plataforma.fullcycle.com.br/courses/242/168/123/conteudos?capitulo=123&conteudo=6968

- Define limits and protection of business rules;
- Components and decoupling:
  - Logs;
  - Cache;
  - Uppload;
  - Database;
  - Commands;
  - Queues;
  - HTTP, APIs, GraphQL;
  - Easy to transform into microservices;

### Basic logic

| client    | interface | application | interface | server     |
| --------- | --------- | ----------- | --------- | ---------- |
| (left)    |           | (center)    |           | (right)    |
| (outside) |           | (inside)    |           | (outside)  |
| -         |           | -           |           | -          |
| rest      |           | domain      |           | database   |
| cli       |           |             |           | redis      |
| rpc       |           |             |           | filesystem |
| graphql   |           |             |           | lambda     |
| ui        |           |             |           | api        |

### Dependency inversion principle

- High level modules should not depend on low level modules. Both mus depend on abstractions;
- Abstractions should not depend on implementation details. Details must depend on abstractions;
- In practical terms, using abstract classes (or interface) in the constructor, instead of concrete classes is a way of decoupling two modules;

## Hexagonal vs Clean vs Onion

Hexagonal Architecture

- Do not define how code should be organized;
- Concept only of what is in and what is out of the application;

Onion and Clean Architecture

- Respect the same concepts as hexagonal architecture;
- Define application layers and how code should be organized;
- Came after hexagonal architecture;

# Clean Architecture

### Key notes

-

## Introduction

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7890

## Clean Architecture Origin

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7889

- Term created by Robert C. Martin (Uncle Bob) in 2012 in [this article](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html);
- Became a book (Clean Architecture);
  - Reinforce concepts and remove basic gaps;
  - Understand components;
  - Understand architectural limits;
  - Perception of business rules;
- Buzz word;
- An evolution of hexagonal architecture;
- Protect the application domain;
- Low coupling level between layers;
- Layers are architectural limits;
- Use-case oriented;
  - User intent;
- Scream the intention of the system;

## Important topics on architecture

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7891

- Define the software shape;
- Components;
- Components communication;
- Good architecture facilitate the system development, deployment, operation and maintenance;

"The strategy behind that facilitation is to leave as many options open as possible, for as long as possible".

Source: Robert C. Martin., Clean Architecture (Robert C. Martin Series), p.136, Pearson Education.

## Keep options open

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7892

"O propósito primário da arquitetura é suportar o ciclo de vida do sistema. Uma boa arquitetura torna o sistema fácil de entender, fácil de desenvolver, fácil de manter e fácil de implantar. O objetivo final é minimizar o custo da vida útil do sistema, e maximizar a produtividade do programador".

Tradução livre.

Source: Robert C. Martin., Arquitetura Limpa, p.137, Alta Books Editora.

### Rules vs Details

- Business rules bring the real value for the software;
- Implementation details help to support the business rules;
- Implementation details should not impact business rules;

## Use cases

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7893

- Describe software intention;
- Describe clearly each software behavior;
- Implementation details should not impact business rules;

### Use cases vs Single Responsibility Principle (SRP)

- We tend to "reuse" use cases since they are very similar;
- Should I reuse similar use cases? No.
- SRP change because of different reasons;
  - If code change because of different reasons that is an alert that you might be breaking the SRP.
- Pay attention to intentional vs accidental duplication;
- Resist to the temptation to apply DRY principle to every part of the code;

## Use cases flow

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7894

- Use cases tell a story;

  - Example (Clean Architecture, p. 192):
    - Use case: Gather Contact Info for New Loan
    - Input: Name, Address, Birthdate...;
    - Output: Same info for readback + credit score;
    - Primary course:
      - Accept and validate name;
      - Validate address, birthdate, etc;
      - Get credit score;
      - If credit score < 500 activate Denial;
      - Else create Customer and activate Loan Estimation;

- Comparing to Domain Driven Design (DDD), the use cases are in the application level. They depend on the business rules;

## Architectural limits

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7895

- Everything tha does not affect the business rules, must be in a different architectural level.
- Front-end and database choices should not affect the application business rules;
- Business rules can depend on abstractions, but not on concrete implementations;
- See figure 17.2 and 17.3 of Clean Architecture book;

## Input vs Output

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7896

- Everything can be summarized in an input that returns an output;

## Data Transfer Object (DTO)

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7897

- Move data between architectural limits;
- Contains only data (input, or output), with no particular behavior;
- Has no business rule;
- Has no behavior;
- In general, each use case has its input and its output dto;

## Presenters

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7898

- Transformation objects;
- Adequate output DTO to the correct format to deliver the result;
- Remember: a system can have different deliver formats (XML, JSON, Protobuf, GraphQL, CLI, etc);

### Example

```js
const inputDto = new CategoryInputDTO("name");
const outputDto = CreateCategoryUseCase(input);

// Presenters
const jsonResult = CategoryPresenter(output).toJson();
const xmlResult = CategoryPresenter(ouput).toXML();
```

## Entities vs DDD

https://plataforma.fullcycle.com.br/courses/242/168/138/conteudos?capitulo=138&conteudo=7899

- Entities
  - DDD
    - Representation of a unique thing;
  - Clean Architecture
    - Enterprise business rules layer;
    - Global business rules;
    - Change only with business;
    - No explicit rules of what is and how to create an entity;
    - Suggestion: use tactics from DDD to create entities;
      - Aggregates;
      - Entities (DDD);
      - Value objects;
      - Contracts;
    - Entities: aggregates + domain services;

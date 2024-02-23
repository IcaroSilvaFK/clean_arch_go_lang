# Use Cases

- São as intenções -> Ações -> Cada Inteção é um caso de uso
- Clareza de cada comportamento do software
- Detalhes não devem impactar nas regras de negócio
- Frameworks, bancos e apis não devem impactar regras

# Use Cases - SRP

### SRP -> single responsability principle

- Temos a tendência de "reaproveitar" use cases por serem muitos
  parecidos

- Ex: Alterar vs Inserir. Ambos consultam se o registro existe, persistem
  dados. Mas, são Use Cases diferentes. Por que?

  Inserir tem uma inteção
  Alterar tem outra intenção

- SRP (Single Responsibility Principle) => Mudam por razões diferentes

# Use Cases Contam uma história.

Gera um fluxo lógico de acordo com oque a regra precisa.
Dita a intenção doque vai fazer.
A camada em que as regras ficam são chamadas de entidades.

## Limites arquiteturais

Tudo que não impacta diretamente nas regras de negócio deve estar
em um limite arquitetural diferente. Ex: Não será o frontend, banco de
dados que mudarão as regras de negócio da aplicação.

![Limites arquiteturais](https://res.cloudinary.com/dlf01tbs6/image/upload/v1708562875/ahtclkm2mqz3uhvupwpy.png)

Observe a direção da seta. O Banco de Dados conhece as Regras de Negócio. As Regras de Negócio não
conhecem o Bando de Dados.

![The clean Architecture](https://res.cloudinary.com/dlf01tbs6/image/upload/v1708563200/z95zooobnxevs705zjrl.png)

# Input vs Output

- No final do dia, tudo se resume a um input que retorna um output
- Ex: Criar pedido (dados do pedido = input) Pedido criado(dados de retorno do pedido)
- Simplifique seu raciocínio ao criar um software sempre pensando em Input e Output

# DTO (Data Transfer Object)

- Trafegar dados entre os limites arquiteturais.
- Objeto anêmico, sem comportamento.
- Contém dados (Input ou Output).
- Não possui regras de negócio.
- Não possui comportamento
- Não faz nada!
- API -> Controller -> USE CASE -> ENTITY.
- Controller cria um DTO com dados recebidos e envia para o Use Case.
- Use Case executa seu fluxo, pega o resultado, cria um DTO para output e retorna para o Controller.

# Presenters

- Objetos de transformação
- Adequa o DTO de output no formato correto para entregar o resultado
- Lembrando: Um sistema por ter diversos formatos de entrega: ex: XML, JSON, Protobuf,GraphQL, CLI, etc.

```js
const input = new CategoryInputDTO("name");

const output = CreateCategoryUseCase(input);
const jsonResult = CategoryPresenter(output).toJson();
const xmlResult = CategoryPresenter(output).toXML();
```

# Entities vs DDD

- Entities da Clean Architecture <> Entities do DDD
- Clean Architecture define entity como camada de regras de negócio
- Elas se aplicam em qualquer situação.
- Não há definição explicita de como criar as entities.
- Normalmente utilizando táticas do DDD.
- Entities = Agregados + Domain Services.

# Infra é a parte do mundo exterior

- HTPP
- Database
- Graphql
- GRPC
- CLI
- ETC

models:
ID:
model: - github.com/99designs/gqlgen/graphql.ID - github.com/99designs/gqlgen/graphql.Int - github.com/99designs/gqlgen/graphql.Int64 - github.com/99designs/gqlgen/graphql.Int32
Int:
model: - github.com/99designs/gqlgen/graphql.Int - github.com/99designs/gqlgen/graphql.Int64 - github.com/99designs/gqlgen/graphql.Int32

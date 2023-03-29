# task-card

## Features



## Microservices
Aplicação foi desenvolvida em uma arquitetura de microserviços

### task-card-cards
Aplicação de gerencia de cards

### task-card-users
Aplicação de gerencia de usuarios

### task-card-notification
Aplicação para notificações de eventos

### task-card-webpage
Website do sistema de task-card

## Data persistence



## Falta implementar:
* Criar webpage
* Criar ambiente k8s


* Criar testes unitário
* Adicionar bidings nas structs
* Adicionar validações




## Envs

```bash
export DB_USER="user"
export DB_PASSWORD="password"
export DB_HOST="192.168.0.113"
export DB_PORT="3306"
export DB_DATABASE="task_cards"
export server="kafka:29092"
export topicConsume="update"
export groupId="myGroup"
export autoOffSetReset="earliest"
export serviceAccount="notification@taskcard.com"
```
# task-card

Application developed from a software to account for maintenance tasks performed during a working day.

## Uploading application:

```bash
git clone https://github.com/luuisavelino/task-card.git

cd ./task-card/infra

# Uploading MySQL and Apache Kafka database
docker-compose up -d

# Creating namespace
kubectl create ns task-card

# Enabling sidecar
kubectl label namespace task-card istio-injection=enabled

# Uploading the Istio gateway, which will expose the endpoints of the applications
kubectl apply -f k8s-gateway

cd ..

# Uploading manifests
kubectl apply -f task-card-cards/k8s
kubectl apply -f task-card-users/k8s
kubectl apply -f task-card-notification/k8s

# Seeing Manifests
kubectl get deploy -n task-card
kubectl get pods -n task-card
kubectl logs -f ${task-card-notification-pod} -n task-card
```

### Observation:
- It is necessary to have Istio installed.

## Testing application

```bash
# Getting the host
export INGRESS_HOST=$(kubectl get po -l istio=ingressgateway -n istio-system -o jsonpath='{.items[0].status.hostIP}')

# Taking the port that is exposed
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')

# Application address
export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT

# See application address
echo "$GATEWAY_URL"
```

With the application address, just send a request to the endpoints.

To see the documentation for each microseries through Swagger, just enter the following addresses:

```bash
# CARDS application swagger
http://$GATEWAY_URL/api/v1/cards/swagger

# USERS application swagger
http://$GATEWAY_URL/api/v1/users/swagger
```

## Features

Features presented by the application

### Card Management

- API endpoint to "create" new card;
- API endpoint to "list" all cards;
- API endpoint to "list" a specific card;
- API endpoint to "delete" a card;
- API endpoint to "update" a card;
- API endpoint to "move" a card;

     #### Comments:

     - The act of deleting a Card is subject to the user with "Manager" permission;
     - User with role "Technician" can only see his own Task;
     - User with role "Manager" can see all tasks;
     - Cards can only be Edited or Moved by their creator or by users with the "Manager" role;

### User Management

- API endpoint to "create" new user;
- API endpoint to "list" all users;
- API endpoint to "list" a specific user;
- API endpoint to "delete" a user;
- API endpoint to "update" a user;

### Card Movement Notification

- If the card is moved, it will generate an update notification of the Card's status. This notification can be seen in the task-card-notification Application Log.

## Architecture

Application was developed in a microservices architecture, using Kubernetes.

For the database and Apache kafka, we chose to leave them in docker-compose, since they do not belong directly to the microservices app mesh. Another reason for leaving both applications out of cluester was an attempt to assimilate the connection with a database and a kafka in production, which are accessed via IP, as in this case.

Applications are:

### task-card-cards:

Card management application
- Location: k8s
- Namespace: task-card

### task-card-users
    
- User management application
- Location: k8s
- Namespace: task-card

### task-card-notification

- Application that prints notifications of "move" events of cards
- Location: k8s
- Namespace: task-card

### MySQL

- data persistence
- Location: docker-compose

### Apache Kafka

- Message broker to decouple notification logic from the application flow
- Location: docker-compose

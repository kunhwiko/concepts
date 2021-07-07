### Docker Compose 
---
##### Use Cases
```
1. Configure relationships between containers 
2. Simplifies docker commands and makes container specs easy to read 
3. Gets an app running in one command 
```

##### YAML configuration
```
docker-compose.yml consists of:
    1. descriptions of containers / networks / volumes 

docker-compose up 
    --> setup and start all containers

docker-compose down
    --> stop all containers and remove containers/volumes/networks 
```

##### Docker Example
```
docker run -p 80:4000 -v $(pwd):/var/lib/nginx kunko/nginx 
```

##### Docker Compose Example
```
version: '2'

services:
  nginx:
    image: kunko/nginx
    volumes: 
      - .:/var/lib/nginx
    ports:
      - '80:4000'
```
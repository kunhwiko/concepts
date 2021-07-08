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
    --> stop all containers and remove containers
```

##### More Commands 
```
docker-compose up -d 
    --> run in background  

docker-compose top 
    --> check running processes in all containers   

docker-compose down -v 
    --> also get rid of volumes 
```

##### Docker Example
```bash
docker run -p 80:4000 -v $(pwd):/var/lib/nginx kunko/nginx 
```

##### Docker Compose Example
```yaml
version: '2'

services:
  nginx:                         # service name: this will be the name of DNS 
    image: kunko/nginx
    volumes: 
      - .:/var/lib/nginx         # "." prints current directory / this is a bind mount
    ports:
      - "80:4000"
```

##### Multiple Docker Compose Example
```yaml
version: '2'

services:
  drupal:
    image: drupal
    ports:
      - "8080:80"
    volumes:
      # this is a named volume 
      - drupal-modules:/var/www/html/modules 
      - drupal-profiles:/var/www/html/profiles
      - drupal-sites:/var/www/html/sites
      - drupal-themes:/var/www/html/themes
  postgres:
    image: postgres
    environment:
     - POSTGRES_PASSWORD=randompassword

# this is used to define named volumes 
volumes:
  drupal-modules:
  drupal-profiles:
  drupal-sites:
  drupal-themes:
```

##### 
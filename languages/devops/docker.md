### Docker Basics
---
##### Images vs Containers 
```
Image: application we want to run 
Container: instance of an image running as a process 
```

##### Container vs Virtual Machine 
```
Containers 
    --> containers run as a process on the host machine 
    --> containers do not come with operating systems 
    --> container are lightweight 
```

##### Run / Start / Top 
```
docker container run -p 8080:80 -d --name kunko -e MYSQL_RANDOM_ROOT_PASSWORD=yes mysql
    --> starts a new mysql container named "kunko"
    --> "-p" (--publish): opens port 8080 on host and forwards traffic to port 80 on container
    --> "-d" (--detach): starts the image on the background 
    --> "-e" (--env): passes in settings 

docker container start <container id>
    --> starts an existing stopped container
```

##### Top / Log / Inspect
```
docker container top <container id>
    --> check running processes inside a container 

docker container logs <container id>
    --> check logs in a container 

docker container inspect <container id>
    --> retrieve configs of a container 
```

##### List / Stop
```
docker container ls 
    --> view running containers  

docker container ls -a 
    --> view all containers 

docker container stop <container id> 
    --> stops running container but doesn't remove it 
```

##### Remove
```
docker container rm <container id> 
    --> remove non-running container 

docker container rm -f <container id> 
    --> force remove container 
```

##### Run Shell
```
docker container run -it --name kunko ubuntu
    --> run "interactive" "tty" 
    --> open a ubuntu container with a CLI 

docker container run -it --name proxy nginx bash 
    --> open an nginx container with a bash CLI

docker container start -ai kunko 
    --> restart kunko with CLI if it previously had one 

docker container exec -it kunko bash 
    --> open a "new" process with a bash CLI for a running container 
```


### Docker Networks 
---
##### Network Basics 
```
docker container port <container id>
    --> inspect host and container port 

docker container inspect --format '{{ .NetworkSettings.IPAddress }}' <container id> 
    --> get IP address of container 

Docker Networks 
    1. containers are connected to private virtual networks 
    2. each virtual network routes to NAT firewall so they can get out to the Internet or other networks 
    3. best practice is to create a new virtual network for each app 

Network (Driver) Types 
    1. bridge: default virtual networks that receive and send info to the host 
    2. host: 
        --> skips the virtual networking of Docker and containers attach to host interface
        --> attaching containers here improve performance but at the risk of security 
    3. null: network not attached to anything

DNS 
    1. DNS is built-in when opening custom networks so containers in a network can communicate with one another 
```

##### Connect to Networks 
```  
docker network ls 
    --> see all virtual networks 

docker network inspect <network id>
    --> see all containers connected in the network 

docker network create my_app
    --> create a new network called "my_app"

docker container run -d --name kunko --network my_app nginx 
    --> add nginx container to my_app network 

docker network connect <network id> <container id>
    --> connect existing container to network 
```


### Docker Images
---
Image Layers
    1. Images are built as a series of layers, users only need to copy over layers they don't have 
    2. If changes are made, a new layer is built on top of existing layers 
    3. If two different changes are made on top of the same existing layers, 
       two side by side layers are built on top of existing layers, 
       instead of duplicating the existing layers 

Container Layers 
    1. Containers continue to build layers on top of base images 
    2. Containers are just single read/write layers on top of base images 
### Docker
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

##### Networks
```
docker container port <container id>
    --> inspect host and container port 

Docker Networks 
    1. containers are connected to private virtual networks 
    2. each virtual network routes to NAT firewall so they can get out to the Internet or other networks 
    3. best practice is to create a new virtual network for each app 
```
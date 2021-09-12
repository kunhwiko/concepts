### Docker Basics
---
##### Images vs Containers 
```
Image: application we want to run 
Container: instance of an image running as a process 
```

##### Containers vs VM vs Traditional Infrastructures
```
Please reference: https://www.ibm.com/cloud/learn/kubernetes 
```

##### Run / Start / Top 
```
docker container run -p 8080:80 -d --name kunko -e MYSQL_RANDOM_ROOT_PASSWORD=yes mysql
    --> starts a new mysql container named "kunko"
    --> caches the mysql image to image cache if it does not exist 
    --> "-p" (--publish): opens port 8080 on host and forwards traffic to port 80 on container
    --> "-d" (--detach): starts the container on the background 
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
    1. containers are connected via private virtual networks linked to host  
    2. each virtual network routes to NAT firewall/host so they can get out to the Internet or other networks 
    3. best practice is to create a new virtual network for each app 

Network (Driver) Types 
    1. bridge: default virtual networks that receive and send info to the host 
    2. host: 
        1. skips the virtual networking of Docker and containers attach to host interface
        2. attaching containers here improve performance but at the risk of security 
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
##### Layers
```
Image Layers
    1. Images are built as a series of layers, users only need to copy over layers they don't have 
    2. If changes are made, a new layer is built on top of existing layers 
    3. If two different changes are made on top of the same existing layers, 
       two side by side layers are built on top of existing layers, 
       instead of duplicating the existing layers 

Container Layers 
    1. Containers continue to build layers on top of base images 
    2. Containers are just single read/write layers on top of base images 
```

##### Tags
```
Tags
    1. Pointer to a particular image commit / version of an image 
    2. Multiple tags can refer to the same commit, so they have the same image ID 

docker pull <repo>:<tag>
    --> pull request an image to the image cache 

docker image tag <source image repo>:<tag> <target image repo>:<tag>
    --> create a new tag for some image commit 
    --> this does not post to Docker Hub yet 

docker image push <repo>:<tag> 
    --> push layers that haven't been pushed yet to your Docker Hub repo 
```

##### Build 
```
Dockerfile 
    1. used to build images
    2. caches steps to make rebuilding faster   

docker image build -t <tag> .
    1. "-t": add tag name 
    2. ".": build image in current directory using Dockerfile 
```

### Container Lifetime 
---
##### Persistent Data
```
Best Practices 
    1. changes should be made to source app, and then containers should be redeployed
    2. changes should not be made directly to containers as the results will not be reproducible
    3. separation of concerns 

Separation of Concerns 
    1. containers should not contain unique data as they might have to be redeployed 
    2. unique data should be stored in some "persistent data" storage 

Uses of Volumes & Bind Mounts 
    1. separates storage from containers 
    2. allows data to be shared among different containers
    3. allows storages to persist 
```

##### Persistent Data 
```
Volume 
    1. assign volume directory using Dockerfiles 
    2. volumes will not be deleted when a container is removed 

docker container run -d --name kunko -e MYSQL_RANDOM_ROOT_PASSWORD=yes -v mysql-db:/var/lib/mysql mysql
    --> "-v": specifies volume specs 
    --> names volume mysql-db and mounts to container 
    --> var/lib/mysql is the path where the file/directory are mounted in the container 
    --> actual data will reside in /var/lib/docker/volumes/mysql-db...

Bind Mounts 
    1. maps host files/directories to a container file/directories
    2. unlike volumes, is not specified in a Dockerfile   

docker container run -d --name kunko -e MYSQL_RANDOM_ROOT_PASSWORD=yes -v $(pwd):/var/lib/mysql mysql 
    --> mounts the files in the host directory to the /var/lib/mysql container directory 
```
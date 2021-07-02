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
```

##### Run / Start / Top 
```
docker container run --publish 8080:80 --d --name kunko nginx 
    --> starts image nginx
    --> opens port 8080 on host IP and forward traffic to port 80 on container IP
    --> "--d" starts the image on the background 
    --> "--name kunko" names the container "kunko"

docker container start 
    --> starts an existing stopped container
```

##### Top / Log
```
docker container top <container id>
    --> check running processes on container 

docker container logs <container id>
    --> check logs on container 
```

##### List / Stop
```
docker container ls 
    --> view running containers  

docker container ls -a 
    --> view all containers 

docker container stop <container id> 
    --> stop running container but doesn't remove it 
```

##### Remove
```
docker container rm <container id> 
    --> remove non-running container 

docker container rm -f <container id> 
    --> force remove container 
```
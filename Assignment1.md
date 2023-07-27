# infracubator
Exercises and commands to work with Docker and Kubernetes


## Docker Commands 
### Start nginx container
```
docker run nginx
```
### Start nginx container, port forward to local and check nginx running
```
docker run -p 8080:80 nginx
```
open browser and enter `localhost:8080` - should be able to see `Welcome to nginx`

-p is used  for port forwarding, `-p <local port>:<container port>`

### check contaier logs 
1. get container id or container name:
```
docker ps
```
`docker ps` will list all containers running with container status. 

2. run `docker log <container id or name>` command will show the logs of a contaier
eg:
```
docker log my-web-server
``` 
### get into the container console

1. start a container with specifying container name
```
docker run --name my-web-server nginx
```
`my-web-server` is a container name
2. run `exec` to pass new command to running  container
```
docker exec -it my-web-server sh
```
This command will open a console for the container. Try to run basic shell commands here eg: `ls` `cd` `mkdir`

### stop a running contaier 
run `docker stop <container name or id>` to stop a running container
eg:
```
docker stop my-web-server
```

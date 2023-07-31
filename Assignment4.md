# infracubator
Exercises and commands to work with Docker and Kubernetes

## Docker Commands - working with docker volumes

### create a docker volume and attach it to a container
1. run command `docker volume create my-volume` to create a volume
2. attach a volume to the container by running below command:
```
docker run -it --rm -v my-volume:/usr/data --name my-golang golang:latest
```

### add content to docker volume and access it with multiple containers
3. get into shell of above running container with volume attached to it;
```
docker exec -it my-golang sh
```
4. navigate to my-volume mounted on container add a file with content:
```
cd /
cd /usr/data
echo "Simple hello text" > hello.txt
```

5. start another container with same volume attached, navigate to volume mount location and check if `hello.txt` exists
```
docker run -it --rm -v my-volume:/usr/data --name my-golang2 golang:latest
cd /usr/data
ls
```
this will show `hello.txt`. run `cat hello.txt` to check the content

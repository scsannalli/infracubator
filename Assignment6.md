# infracubator
Exercises and commands to work with Docker and Kubernetes

## Docker Commands - working with docker networks

### Working with containers in the same network

1. create a new docker network by running below command: 
```
docker network create my-network
```
2. run redis contaier and attach it to a network created `my-network`

```
docker run --name my-redis -d redis
docker network connect my-network my-redis
```

also by running command `docker container inspect my-redis` validate if `my-network` shows up in the `Networks` sections.

3. Run the go-contaier with the same network created `my-network` and pass redis `container name` instead of `ip address`
```
docker run -it -e REDIS_HOST=my-redis -p 8080:8080 --network my-network go-redis-service
```
`go-redis-service` is the image created in `Assignment5.md` 

4. open another terminal and run below command to check `/vote` endpoint is working fine
```
curl -X GET localhost:8080/vote
```


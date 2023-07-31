# infracubator
Exercises and commands to work with Docker and Kubernetes

## Docker Commands - working with multiple containers

### run a redis container and run go-app with redis enabled
1. clone go-lang repo `git clone https://github.com/docker-ninja/go-app.git`
2. switch to `redis` branch 
```
cd go-app
git checkout redis
```
3. build new image of go-app: 
create a `Dockerfile` to build new go app. 
`Dockerfile` file is placed in `golang-redis` directory for reference. 
Now build an image with docker-app by running command  below:
```
docker build -t go-redis-service .
```
4.  run redis container and redis container ip
```
docker run --name my-redis -d redis
```
Now run below command to get container details and ip address
```
docker container inspect my-redis
```
from the output results get `IPAddress` from `NetworkSettings`

5. Now run the application with `REDIS_HOST` environment variable
```
docker run -it -e REDIS_HOST=172.17.0.2 -p 8080:8080 go-redis-service
```
6. open another terminal and run below command to check `/vote` endpoint is working fine
```
curl -X GET localhost:8080/vote
```




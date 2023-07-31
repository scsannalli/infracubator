# infracubator
Exercises and commands to work with Docker and Kubernetes

## Docker Commands

### Running golang service using golang container
1. Create a go lang app: `application source code is placed under `/golang` folder`
2. Navigate to `/golang` location and build the app using  `golang:latest` image by running below command:
   ```
   docker run -it --rm -v .:/app -w /app golang:latest go build -o out/ ./...
   ```
   this creates a compiled version of go service in `/out` location 
4. run the container to start golang-application:
    ```
    docker run -it --rm -v .:/app -w /app -p 3333:3333 golang:latest ./out/http-server
    ```
5. From another  terminal, Run curl request to golang-application
    ```
    curl -X GET "localhost:3333"
    ```

### tag image with version
 `golang:latest` image can be tagged with specific version `golang:v1` to push under specific registry, run below command:
```
docker tag golang:latest scsannalli/golang:v1
```

### docker history of image
For the image created and tagged in the previous commands, check the image history by running below command: 
```
docker history scsannalli/golang:v1
```

### push the image to docker hub
1. tag an image with repository id/url
    eg:\
    `docker tag golang:latest scsannalli/golang:v1`\
    `scsannalli` is docker id of an user
2. login to docker hub
```
docker login -u <user name> -p <password>
```
3. push image to hub by running below command
```
docker push scsannalli/golang:v1
```
4. Now anyone can pull `scsannalli/golang:v1` image from public repo\
 and run the container


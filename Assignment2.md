# infracubator
Exercises and commands to work with Docker and Kubernetes

## Docker Commands

### Running Docker container of golang app
1. Create a go lang app: `application source code is placed under `/golang` folder`
2. Create docker file to build image with golang app: `Dockerfile is placed under `golang/Dockerfile` `
3. Navigate to `/golang` location and build image `go-service-image` by running below command
   ```
   docker build -t go-lang-service .
   ```
   `go-lang-service` is the user defined image name 
4. run the container to start go-lang-application
    ```
    docker run -it --rm -p 3333:3333 go-lang-service
    ```
5. From another  terminal, Run curl request to go-lang-application
    ```
    curl -X GET "localhost:3333"
    ```

### tag image with version
Above created go-lang-service image will be by default tagged as `go-lang-service:latest`
Now same image can be tagged with specific version `go-lang-service:v1` : run below command
```
docker tag go-lang-service:latest go-lang-service:v1
```

### docker history of image
For the image created and tagged in the previous commands, check the image history by running below command: 
```
docker history go-lang-service:v1
```

### push the image to docker hub
1. tag an image with repository id/url
    eg:\
    `docker tag go-lang-service:v1 scsannalli/sample-go-lang-service:v1`\
    `scsannalli` is docker id of an user
2. login to docker hub
```
docker login -u <user name> -p <password>
```
3. push image to hub by running below command
```
docker push scsannalli/sample-go-lang-service:v1
```
4. Now anyone can pull `scsannalli/sample-go-lang-service:v1` image from public repo\
 and run the container


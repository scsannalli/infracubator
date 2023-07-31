# infracubator
Exercises and commands to work with Docker and Kubernetes

## Docker Commands - multistage docker files

### Running Docker container of golang app using multistage docker file
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
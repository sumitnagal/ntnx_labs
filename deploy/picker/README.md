<img src="../../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>


# Deployment Lab

---
## <img src="../../assets/lab.png" width="auto" height="32"/> Your Mission...

> Build and dockerize a word picker web application.

1. Ensure all the tests are passing!
2. Build an executable by hand to run on your target operating system
   1. Make sure you can set the application version via your build flags
   2. Run the picker service locally
   3. Ensure it prints your custom version upon starting!
   4. Exercise the picker by using the URLS below
3. Modify the provided Dockerfile and .dockerignore files to build a Docker image
   1. Make sure you can set the application version via your build command
4. Run your docker container via locally via Docker
5. Ensure it prints your custom version upon starting
6. Verify your service endpoints are working correctly on the container
7. Record the size of your current docker image!
8. Next change the Docker base image to use a scratch image
9. Rebuild your docker image
10. Repeat the step above to launch your new Docker container and validate the
    service
11. Now the surprise!
    Compare the size of this new image you've just build vs the old image?


## Installation

1. Install [Docker Desktop](https://www.docker.com/products/docker-desktop) on your machine (SKIP IF INSTALLED!)

## Commands

1 Exercise service endpoints

   ```shell
   # Load a dictionary
   # Available dictionaries: words, trump or musicians
   curl -XGET http://localhost:4500/api/v1/load?name=musicians
   # Word picker...
   curl -XGET http://localhost:4500/api/v1/picker
   ```

1. Build a docker images

   ```shell
   docker build --rm -t picker:0.0.1 .
   ```

1. Run a docker image exposing 4500 locally

   ```shell
   docker run -it -p 4500:4500 picker:0.0.1
   ```

1. List your Docker image

   ```shell
   docker images | grep picker
   ```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
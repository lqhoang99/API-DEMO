# TodoList-Echo
This is Todos
# How to build it
docker build --tag apidemo:1.0 .
# How to run it
docker run --publish 3000:3000 --detach --name rb apidemo:1.0
# How to run docker compose up
docker-compose up --build
# How to remove one container 
docker rm -f 492ae81ea5dd
"492ae81ea5dd" is container name(id)
# how to list container 
docker container ls
# How to list local image 
docker image ls
# How to stop application
docker-compose down
OR CTR+c
# how to push it to docker hub
docker tag apidemo:1.0 phucmars/apidemo:1.0
docker push phucmars/apidemo:1.0
# how to pull 
docker pull phucmars/apidemo
# If you want to run your services in the background, you can pass the -d flag (for “detached” mode) to docker-compose up and use docker-compose ps to see what is currently running:
docker-compose up -d
docker-compose ps

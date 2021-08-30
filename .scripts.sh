docker image ls
docker build --tag myapp .
docker image prune
docker system prune
docker run -d -p 8080:8080 myapp
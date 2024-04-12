#Go api 
#Docker container

docker build -t go-server-app .
docker run -e PORT=3001 -p 3001:3001 go-server-app



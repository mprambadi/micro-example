docker rm -f $(docker ps -aq)
docker network rm live-discussion-app

docker network create live-discussion-app

docker run --name app-node -p 4000:4000 -d --network live-discussion-app mprambadi/app-node  
docker run --name app-go -p 4001:4001 -d --network live-discussion-app username/app-go  
docker run --name app-php -p 4002:4002 -d --network live-discussion-app username/app-php  
docker run --name app-nginx -p 80:80 -d --network live-discussion-app -v $(pwd)/app-nginx/default.conf:/etc/nginx/nginx.conf nginx
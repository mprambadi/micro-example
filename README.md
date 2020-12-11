

Build all docker image 
```
docker build -t mprambadi/app-node . 
docker build -t mprambadi/app-go . 
docker build -t mprambadi/app-php . 
```

Create network 
```
docker network create live-discussion-app
```


Run all app container image 
```
docker run --name app-node -p 4000:4000 --network live-discussion-app mprambadi/app-node  
docker run --name app-go -p 4001:4001 --network live-discussion-app username/app-go  
docker run --name app-php -p 4002:4002 --network live-discussion-app username/app-php  
```

Run nginx image 
```
docker run --name app-nginx -p 80:80 --network live-discussion-app -v $(pwd)/app-nginx/default.conf:/etc/nginx/nginx.conf nginx
```
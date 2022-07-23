# Installation
## Create bridge network
``` docker network create testnet ```
## Database
###
### Run with Docker
``` docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 --network testnet --network-alias mysql_host mysql:5.7 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci ```
### Copy and execute sql file
```docker cp ./sql/simpleRest.sql mysql:/home/ ```

```docker exec -it mysql mysql -uroot -p123456 ```

``` mysql < source /home/simpleRest.sql ```

### Mysql user : root pwd : 123456
 <br/>
 <br/>

## back-end Application
### Build

``` docker build -t simple-rest-api . ```

### Run with Docker
``` docker run -d --name simple-rest-api -p 80:80 --network testnet --network-alias simple-rest-api_host simple-rest-api ```



###

full stack web application 
frontend = sveltekit
https://github.com/Sirneij/django_svelte_jwt_auth
backend = go
database = arango

## docker-compose



## docker setup for all
docker run -v /home/thebogie/work/arangodb/collection:/var/lib/arangodb3 -v /home/thebogie/work/arangodb/apps:/var/lib/arangodb3-apps  -p 50001:50001 -p 50002:50002 -p 50003:50003 -d --name stgangdocker stgangdocker 

In Docker:
arangod  --server.authentication=false ; 
cd /stg/back ; ./main

## docker setup for back+db
docker run -v /home/thebogie/work/arangodb/collection:/var/lib/arangodb3 -v /home/thebogie/work/arangodb/apps:/var/lib/arangodb3-apps  -p 50001:50001 -p 50002:50002 --name stggoarangoflutter stggoarangoflutter 

In Docker:
arangod  --server.authentication=false &
cd /stg/back ; ./main
## just database:
docker run -v /home/thebogie/work/arangodb/collection:/var/lib/arangodb3 -v /home/thebogie/work/arangodb/apps:/var/lib/arangodb3-apps  -p 50001:50001  --name stggoarangoflutter stggoarangoflutter 
windows: docker run -v C:/work/arangodbstorage:/var/lib/arangodb3  -p 50001:50001  --name stggoarangoflutter stggoarangoflutter

arangod  --server.authentication=false ; 


## Backend


## Frontend


## Config


## TODO
FrontEnd:
- figure out what needs to be fixed to get it up and running
- 
Backend
- fetch venue by address
  - how to connect to arango and get venue by address
- push venue by address
  - if already exists, update?
  - push element into arango



##Done
- for JSON: arangoexport --server.database smacktalk --collection game --output-directory "games"
- for argango backup/restore: arangodump --server.database smacktalk --output-directory "today"
- docker-compose.yml -> docker exec -it backend_arangodb_db_container_1 sh





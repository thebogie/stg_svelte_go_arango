
# Smacktalk Gaming

full stack web application 
- frontend = sveltekit
- backend = go
- database = arango


### docker setup for all
docker run -v /home/thebogie/work/arangodb/collection:/var/lib/arangodb3 -v /home/thebogie/work/arangodb/apps:/var/lib/arangodb3-apps  -p 50001:50001 -p 50002:50002 -p 50003:50003 -d --name stgangdocker stgangdocker
In Docker:
arangod  --server.authentication=false ; 
cd /stg/back ; ./main

### docker setup for back+db
docker run -v /home/thebogie/work/arangodb/collection:/var/lib/arangodb3 -v /home/thebogie/work/arangodb/apps:/var/lib/arangodb3-apps  -p 50001:50001 -p 50002:50002 --name
In Docker:
arangod  --server.authentication=false &
cd /stg/back ; ./main

### just database:
docker run -v /home/thebogie/work/arangodb/collection:/var/lib/arangodb3 -v /home/thebogie/work/arangodb/apps:/var/lib/arangodb3-apps  -p 50001:50001  --name   
windows: docker run -v C:/work/arangodbstorage:/var/lib/arangodb3  -p 50001:50001  --name stggoarangoflutter stggoarangoflutter
arangod  --server.authentication=false ;
docker-compose -f .\docker-compose.dev.yml up -d arangodb

## Upcoming
- FrontEnd:
  - profile front page. what to hold in it? how to click through to show list of games and venue
  - add contest page

- Backend
  - fetch venue by address
- push venue by address
  - if already exists, update?
  - push element into arango



##Extra
- for JSON: arangoexport --server.database smacktalk --collection game --output-directory "games"
- for argango backup/restore: arangodump --server.database smacktalk --output-directory "today"
- docker-compose.yml -> docker exec -it backend_arangodb_db_container_1 sh





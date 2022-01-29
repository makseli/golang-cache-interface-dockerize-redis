# Dockerize Golang Cache Interface Implementation

For [Golang Notları #6 interface_revision_2 PoC ](https://makseli.medium.com/golang-notlar%C4%B1-6-interface-revision-2-d9af2613df28)

## Requirement
 - git
 - docker-compose
 - internet browser :)

## Run

- git clone https://github.com/makseli/golang-cache-interface-dockerize-redis.git && cd golang-cache-interface-dockerize-redis
- docker-compose up (After making sure other applications are not running on port 5000 and 6379)
- http://localhost:5000 open the address from the browser. In this step, the keys and information determined after the objects are created for the Cache will be saved through the Cache classes. 
- When you open the http://localhost:5000/getMC address from the browser, the text “Data from Local Memory : Kawasaki KLE 500” should be displayed. 
- When you open http://localhost:5000/getRedis from the browser, the text “Data from Redis : Kawasaki KLE 650” should be displayed. 


Continue -> **https://makseli.medium.com/golang-notlar%C4%B1-7-cache-interface-implementation-repo-with-redis-dockerize-77d54832598b**
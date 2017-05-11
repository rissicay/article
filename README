## How to run

1. Increase vm.max_map_count to run elasticsearch 
`sudo sysctl -w vm.max_map_count=262144`
More information can be found on: https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html
2. docker-compose up
3. http://localhost:8080

## Assumptions

1. POST /articles accepts the same output as GET /articles/{id}
2. GET /articles/{id} is exactly the same as given in the example
3. Fields that are empty are emitted in the JSON

## Stack choices

Go was selected as the language of choice because:
1. Its viewed favourably
2. I've used it before

The framework choice is Gin, because I wanted a router but something leaner than
a full framework like beego. Gin seems to fit both these categories.

The data store selected was elasticsearch. I would never use elasticsearch as a
data store in a production system but since this a technical exam I wanted
to have some fun, plus elasticsearch is cool. For these endpoints, elasticsearch
offers no benefits over a SQL data store.

Docker was used for the container environment out of familiarity.

## Code structure

Because of the simplicity of the endpoint, there isnt much structure to the code.
Code that touches the router and code that touches the data store have been 
separated out into different functions.

Error handling, logging and tests have largely been ignored. I wanted to stay around the 
rough working hours listed in the specs. Trying to test this code is pretty difficult.
The main meat of the code is in the fetch and insert functions. These functions are tightly
coupled to elasticsearch. If I had more time, I'd break the calls to elasticsearch out
to other functions so I could mock them.

## Running a development environment

Environmental variables must be set before a typical go dev environment can be 
used. These can be found in 
`docker/.env`

Elasticsearch must be running before web is started

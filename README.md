# DB_RAINEAU_P01

## Description

This repository contains all Golang API code for the databse exercice given in class

## Usage

Launch the containers : 

``` docker-compose up --build```

Create the database : 

``` docker exec -i go-exo-bdd_db_1 sh -c 'exec mysql -ugobdd -p"gobdd" image_gobdd' < ./docker/data/database.sql```

Now you can access the API at localhost:8080.

## Testing the routes

Use the postman collection available in docker/data directory

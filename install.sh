# build docker containers
docker-compose up --build -d

# Wait for the database to be connected
sleep 10

# Create the database
docker exec -i go-exo-bdd_db_1 sh -c 'exec mysql -ugobdd -p"gobdd" image_gobdd' < ./docker/data/database.sql
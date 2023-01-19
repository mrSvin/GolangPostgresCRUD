Download dependencies
````
go get all
````
# Проверяем связь сб, Создаем схему и проверяем ее
````
docker-compose up -d
docker-compose exec pgdb psql -U postgres -c 'CREATE DATABASE db_test'       
docker-compose exec pgdb psql -U postgres -c 'SELECT pid, usename, state, query FROM pg_stat_activity WHERE state IS NOT NULL;'
````
# Создаем таблицу
````
'CREATE TABLE users (id SERIAL NOT NULL, created_at TIMESTAMP, name VARCHAR(100), age INTEGER, verify BOOLEAN)
````
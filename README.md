# avtoDom

make docker container:
> docker run --name=postgres -e POSTGRES_PASSWORD="postgres" -p 5437:5432 -d --rm postgres

make migrations:
> migrate -path ./shemas -database 'postgresql://postgres:postgres@localhost:5437/postgres?sslmode=disable' up
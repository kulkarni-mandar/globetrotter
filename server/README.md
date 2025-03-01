# Server

## Docker command to start database
```
docker run --name postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=globetrotter -p 5432:5432 --rm -d postgres
```
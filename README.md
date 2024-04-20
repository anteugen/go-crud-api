### To retrieve all movies:

curl http://localhost:8080/movies

### To retrieve a specific movie by ID:

curl http://localhost:8080/movies/<movie_id>

### To add a new movie:

curl -X POST -H "Content-Type: application/json" -d '{"name":"Movie Name","description":"Movie Description"}' http://localhost:8080/movies

### To delete a movie by ID:

curl -X DELETE http://localhost:8080/movies/<movie_id>

### To update an existing movie:

curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Name","description":"Updated Description"}' http://localhost:8080/movies/<movie_id>


# admin-handling

## Start Up

* Fork this repository.
* Create a `.env` file with parameters
    * `DB_USERNAME`
    * `DB_PASSWORD`
    * `DB_NAME`
    * `DB_HOST`
    * `DB_PORT`

* Input the MySQL database server information into `.env` according to the format above. This database will contain the data used for this set of API endpoints.
Example:
    * `DB_USERNAME=root`
    * `DB_PASSWORD=12345678`
    * `DB_NAME=test`
    * `DB_HOST=localhost`
    * `DB_PORT=3306`

* Run `go run migration/migration.go` to migrate the database tables into your own database.

## Running the sever
* Using the terminal, change the current working directory to /admin-handling
* Host the server locally by running `go run .` in the terminal

* Add teachers manually into the database by opening a new terminal and running 
    * `curl http://localhost:8080/api/teacher --request "POST" --data '{"teacher": "<email_of_teacher>"'` changing `<email_of_teacher>` to the desired email.
* Add students manually into the database by opening a new terminal and running 
    * `curl http://localhost:8080/api/student --request "POST" --data '{"teacher": "<email_of_student>"'` changing `<email_of_student>` to the desired email.

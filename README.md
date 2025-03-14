# Database of library
My implementation of the course work on the subject of databases and network technologies, which I completed in the 2nd year

## Technologies
- DB: PostgreSQL
- Scripting: Golang

## Information sources
- Books, authors and publishers: `chitai-gorod.ru`
- Customers: API `https://randomdatatools.ru` which generates fake data

## Database schema
![Library database schema](images/schema.png)

## Configuration
Add file `config.yaml` with postgres credentials of the following format to the project directory
```yaml
postgresql:
    host:
    port:
    user:
    password:
    db_name:
    ssl_mode:
```
Example:
```yaml
postgresql:
  host: "localhost"
  port:  5432
  user: "test-user"
  password: "qwert"
  db_name: "test-db"
  ssl_mode: "disable"
```
Also add postgres container data in `cmd/env.sh`.
All fields except port are of the string format, port is integer.
Add docker parameters to `cmd/run.sh` to the appropriate export commands.

## Running
You need installed Go 1.22 and docker.
Run with:
```bash
cmd/run.sh
```
Dump of db will be in `dump.sql` file in project directory

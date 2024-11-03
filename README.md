# Database of library
My implementation of the course work on the subject of databases and network technologies, which I completed in the 2nd year

## Technologies
- DB: PostgreSQL
- Scripting: Golang

## Information sources
- Books, authors and publishers: `chitai-gorod.ru`
- Customers: API `https://randomdatatools.ru` which generates fake data

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
All fields except port are of the string format, port is integer
Add docker parameters to `cmd/run.sh` to the appropriate export commands

## Running
You need installed golang and docker
Run
```bash
cmd/run.sh
```
Dump of db will be in `dump.sql` file in project directory

## Just simple rest API with go and gin

### Requirements
- Go 1.17+
- Reflex - https://github.com/cespare/reflex
- VSCode extension Rest Client or similar to make requests from api.http text file

### Install
- create `.env` file - `cp .env-sample .env` and set values if necessary

#### Create Postgres table
- exec to postgres docker container
- connect to DB - `psql -U postgres gorestdb` (password `postgres`)
- add `uuid-ossp` extention - `CREATE EXTENSION "uuid-ossp";`
- create table - `CREATE TABLE albums (id uuid DEFAULT uuid_generate_v4() PRIMARY KEY, title VARCHAR, artist VARCHAR, year SMALLINT);`
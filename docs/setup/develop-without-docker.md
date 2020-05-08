# Developing without Docker

You'll need:
- Go 1.13
- Postgres 11 or newer (SQLite is not supported)

1. Create a Postgres database and a user.

   For example, these commands create:
   - a database named `komfy` 
   - a user with usernamed `komfy` and password `rocks` 

   on a local Postgres instance:
   ```
   sudo -u postgres bash
   psql -c "CREATE USER komfy WITH PASSWORD 'rocks'"
   psql -c "CREATE DATABASE komfy WITH OWNER komfy"
   ```

   The resulting database URL (`DATABASE_URL` in the `.env` file) will be
   `postgres://komfy:rocks@localhost:5432/komfy`

   For more documentation on these commands please refer to the official Postgres documentation:
   - https://postgresql.org/docs/current/sql-createuser.html for `CREATE USER`
   - https://postgresql.org/docs/current/sql-createdatabase.html for `CREATE DATABASE`
   
   For other clients (e. g. pgWeb, pgAdmin, dBeaver, DataGrip, etc.) please refer to their respective documentation.

2. Fill out the `.env` file as shown in the 
   [`.env.example` file](https://github.com/komfy/api/blob/master/.env.example).

   Tip: `.env` files can be parsed by any Bourne shell-compatible shell (e. g. sh, ash, dash, bash),
   so, if you want to use values from the `.env` file in your shell, run
   `source .env` or `. ./.env`. 
   Note: this doesn't work in Fish as it has a different way of defining variables.

3. Run the `sql/populate_tables.sql` script on your Postgres database.
   On a local Postgres instance one can run `sudo -u postgres psql -f "sql/populate_tables.sql"`

4. Run `task -w dev` in project folder.


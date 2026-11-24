# AWS DevOps Lab

The app uses the existing Docker container named `mysql`, with MySQL published on port `3306`. From the repository root, create the `aws_devops_lab.items` table, seed three items, and grant the app's `mysql` account read access. This command uses the local container's root password `mysql` and can be rerun.

```sh
docker exec -i -e MYSQL_PWD=mysql mysql mysql -u root < infrastructure/mysql-init.sql
```

Start the Go backend and both frontends.

```sh
docker compose up --build
```

| Service | Local URL |
| --- | --- |
| Frontend A | <http://localhost:8091> |
| Frontend B | <http://localhost:8092> |
| Backend health endpoint | <http://localhost:8080/api/health> |
| Items from MySQL | <http://localhost:8080/api/items> |

Nginx forwards both frontends' `/api/` requests to Go. `/api/health` checks that Go is reachable, and `/api/items` reads the MySQL table on each request.

Compose connects to MySQL through `host.docker.internal:3306` on Docker Desktop. The local defaults for the app username and password are both `mysql`. Override the connection settings with `MYSQL_ADDR`, `MYSQL_USER`, and `MYSQL_PASSWORD`.

To run the backend directly with Go 1.26 or later, start it from `apps/backend` with `MYSQL_ADDR=127.0.0.1:3306 MYSQL_USER=mysql MYSQL_PASSWORD=mysql go run .`.

Set up:
1) download library for migration https://github.com/golang-migrate/migrate/releases
2) check config.yaml and write your database date
3) run the command in the console:
  migrate -database postgres://USER:PASSWORD@localhost:5432/DB?sslmode=disable -path migrations up
4) check swagger docs

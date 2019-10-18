# mysqlhealth
Basic health check for mysql exposed in an http endpoint.  

This healthcheck is intended for those who need an http endpoint to check the status of mysql, for examplke it could be used in consul service discovery or as healtcheck or readyness check for kubernetes. The check implemens a connection to mysql and issues a `SELECT 1` statement to check availability of the mysql server. Thise service intention is to improve upon doing a tcp check since that doesn't actually indicates the mysql server is availabel and ready.  


# Dependencies

- GO 1.12 (Tested with)
- Gorila Mux (github.com/gorilla/mux)
- GO sql driver (github.com/go-sql-driver/mysql)
- Automake (if you need to use the Makefile for building the binaries)

# Build

```
# just run
go build
```

# How to use
```bash
export  MYSQL_SOURCE_NAME='user:password@tcp(IP:PORT)/DATABASE' # Leave database empty if you just want to use the mysql local database to check the overall availability of the server
mysqlhealth -help
./mysqlhealth 0.1 Exposes an http health endpoint for mysql health checks.
It uses MYSQL_SOURCE_NAME for mysql connection environment variable with following format: https://github.com/go-sql-driver/mysql#dsn-data-source-name
Default value is "mysql:mysql@tcp(localhost:3306)/".

Usage: ./mysqlhealth [flags]

Flags:
  -version
        Print version information and exit.
  -web.health-path string
        Path under which to expose health endpoint (default "/health")
  -web.listen-address string
        Address to listen on for web interface (default ":42005")
```

# Disclaimer
This software has been built on my free time and it shouldn't be considered as production ready.  
The software doesn't take care of any security best practices as it uses plain text HTTP and doesn't implement any form of authentication or rate limiting.  
USE AR YOUR OWN RISK! 

# TODO

- Add Dockerfile
- use golang docker container for building to prevent the need of having golang installed
- Add tests (it so simple that probably doesn't need one)
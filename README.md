# mysqlhealth
Basic health check for mysql exposed in an http endpoint.  

This healthcheck is intended for those who need an http endpoint to check the status of mysql, for examplke it could be used in consul service discovery or as healtcheck or readyness check for kubernetes. The check implemens a connection to mysql and issues a `SELECT 1` statement to check availability of the mysql server. Thise service intention is to improve upon doing a tcp check since that doesn't actually indicates the mysql server is availabel and ready.  


# Dependencies

- GO 1.12 (Tested with)
- Gorila Mux (github.com/gorilla/mux)
- GO sql driver (github.com/go-sql-driver/mysql)

# Build

```
# just run
go build
```

# Disclaimer
This software has been built on my free time and it shouldn't be considered as production ready.  
The software doesn't take care of any security best practices as it uses plain text HTTP and doesn't implement any form of authentication or rate limiting.  
USE AR YOUR OWN RISK! 

# TODO

- Add Documentation to README
- Add makefile
- Add Dockerfile
- Add tests (it so simple that probably doesn't need one)
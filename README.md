# implight-backend
Backend to store highlights via chrome extension

## Migration Process
Tool used - [golang-migrate](https://github.com/golang-migrate/migrate)

Create migration:

```
migrate create -ext sql -dir db/migrations create_token
```

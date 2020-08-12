# {CNG} Accounts micro service
Golang with Gin router and Postgres as database.

### Running service:

#### `$ make` 

### Avaliable routes:

#### POST
- /api/validate
- /api/new-account
- /api/account-details

getting IP from docker-compose 

```docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' cng-accounts-api_accounts_api_1```

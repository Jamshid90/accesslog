# Access log service

## Database migrations

#### Install golang-migrations
```
brew install golang-migrate
```
from: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

## Create migrations example
```
 migrate create -ext sql -dir migrations -seq create_xxxx_table
```

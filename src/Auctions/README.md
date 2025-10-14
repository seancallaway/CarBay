# Auction Service

This service is written in [Go](https://go.dev/) and utilizes [gin](https://gin-gonic.com/) and [gorm](https://gorm.io/).

It's spec can be found [here](https://github.com/seancallaway/CarBay/wiki/Auction-Service-Spec).

## DB Migrations

Set the [appropriate environment variables](inits/db.go) and run `go run migrations/migrations.go`.

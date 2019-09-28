#!/bin/sh

# mage
go get -v -u -d github.com/magefile/mage
cd $GOPATH/src/github.com/magefile/mage
go run bootstrap.go

# sqlboiler
go get -v -u -t github.com/volatiletech/sqlboiler
go get -v -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

# sql-migrate
go get -v -u github.com/rubenv/sql-migrate/...

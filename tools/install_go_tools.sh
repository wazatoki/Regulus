#!/bin/sh

# mage
GO111MODULE=off go get -u -d github.com/magefile/mage
cd $HOME/go/1.3.1/src/github.com/magefile/mage
go run bootstrap.go

# sqlboiler
# Install sqlboiler v4
GO111MODULE=off go get -u -t github.com/volatiletech/sqlboiler
# Install an sqlboiler driver - these are seperate binaries, here we are
# choosing postgresql
GO111MODULE=off go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

# sql-migrate
GO111MODULE=off go get -v -u github.com/rubenv/sql-migrate/...

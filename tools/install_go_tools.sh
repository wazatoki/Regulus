#!/bin/sh

# mage
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go

# sqlboiler
# Install sqlboiler v4

GO111MODULE=on go get -u -t github.com/volatiletech/sqlboiler/v4
# Install an sqlboiler driver - these are seperate binaries, here we are
# choosing postgresql
GO111MODULE=on go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql

# sql-migrate
GO111MODULE=off go get -v -u github.com/rubenv/sql-migrate/...

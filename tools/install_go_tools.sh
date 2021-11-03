#!/bin/sh

# add path
echo "export PATH=$PATH:/workspace/Regulus/tools" >> ~/.bashrc

# mage
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go

# sqlboiler
# Install sqlboiler v4

go install github.com/volatiletech/sqlboiler/v4@latest
# Install an sqlboiler driver - these are seperate binaries, here we are
# choosing postgresql
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

# sql-migrate
GO111MODULE=off go get -v -u github.com/rubenv/sql-migrate/...

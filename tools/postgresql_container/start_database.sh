#!/bin/sh

docker run --name regulus_db\
	-v /home/vagrant/tools/postgresql_container/postgresql:/mnt\
	-e POSTGRES_USER=regulus\
	-e POSTGRES_PASSWORD=regulus\
	-e POSTGRES_DB=regulusdb\
	-p 5432:5432\
	-d postgres:11.3

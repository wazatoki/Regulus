#!/bin/sh

docker run --name postgresql\
	-v /home/vagrant/tools/postgresql_container/postgresql:/mnt\
	-e POSTGRES_USER=demo\
	-e POSTGRES_PASSWORD=demo\
	-e POSTGRES_DB=demodb\
	-p 5432:5432\
	-d postgres:11.3

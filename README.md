# Regulus

## Overview

The development environment is assumed to be linux.

## Dependency
* Docker
* anyenv
* goenv
* nodenv
* postgresql
* sqlboiler


## DB default Setup

* url: localhost:5432/regulus
* username: root
* password: password

### db setup by docker on linux

execute tools/postgresql_container/start_database.sh

### production setup

prease setup production database by postgresql

## Usage

1. anyenv update plugin install

    git clone https://github.com/znz/anyenv-update.git $(anyenv root)/plugins/anyenv-update

    anyenv update

1. install golang

    goenv install 1.12.9
  
    goenv local 1.12.9

1. install nodejs

    nodenv install 12.9.1

    nodenv local 12.9.1

1. install go tools

    ./tools/install_go_tools.sh

1. download dependency liblalyies

    cd src

    go build

    cd js

    npm install



## Licence


## Author
wazatoki

## References

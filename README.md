# Ijah Shop

The implementation of the well-known interview test Ijah Shop. The goal is to design and implement a system based from the existing Ijah's Excel inventory management.  
For the complete instruction, please see [this link](https://www.dropbox.com/s/7omfzk55284omjd/ijah.pdf?dl=0)

## Architecture Overview

It uses Onion / Hexagonal / Clean architectures in a simpler way. It uses 3 layers: the deepest domain layer, the infrastructure (databases), and the application layer (http server, serialization, etc)

![diagram](https://ucc9296cffb79b303700649379c9.dl.dropboxusercontent.com/cd/0/inline/ASSG_xlLxRUjmgLtyVt-YrZUULaLGrQR5IqoUVvANMi6Xp34TR4lElrjc4Pv9aRDaB1lRffYM-C9yck0i3WxJTX24R7ufqPAXwlY5aK3tqlKYV9sMCVv1cfYT9UC6LSk7qN_Kt8X3voh1rd66jgGNuxtYSCbajVuMxyoyqN32ZkwfKMTx1mKf-J70rRjLkIebLE/file)

## Requirements and Dependencies
- Go 1.11
- [golang/dep](https://github.com/golang/dep)
- [labstack/echo](https://github.com/labstack/echo)
- [satori/go.uuid](https://github.com/satori/go.uuid)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

## Installation
- run `dep ensure` to install dependencies
- run `go run main.go` to run the server

## Notice
- There is no DDL for tables yet. You need to manually create the table.
- Database configuration is still hardcoded
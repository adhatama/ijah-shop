# Ijah Shop

The implementation of the well-known interview test Ijah Shop. The goal is to design and implement a system based from the existing Ijah's Excel inventory management.  
For the complete instruction, please see [this link](https://www.dropbox.com/s/7omfzk55284omjd/ijah.pdf?dl=0)

**Disclaimer: I didn't apply to that company for an interview. This repo is just for my attempt to implement Go architecture for learning purposes.**

## Architecture Overview

It uses Onion / Hexagonal / Clean architectures in a simpler way. It uses 3 layers: the deepest domain layer, the infrastructure (databases), and the application layer (http server, serialization, etc)

![diagram](https://dl.dropboxusercontent.com/s/tisp0b23cc9sjcc/diagram.png)

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

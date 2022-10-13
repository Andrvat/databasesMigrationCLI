# Databases Migration CLI
CLI application that provides databases migration operations {up [ver], down [ver]}


## Table of Contents

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Quick Start](#quick-start)

## Introduction

This is a CLI application that provides databases migration operations {up [ver], down [ver]}. 


## Requirements
The application can be run locally or in a docker container, the requirements for each setup are listed below.

### Local
* [Golang latest](https://golang.org/dl/)
* [Library migrate/migrate for Golang](https://github.com/golang-migrate/migrate)

### Docker
* [Docker](https://www.docker.com/get-docker)


## Quick Start
TODO

### Launch Requirements
Before run this CLI app you need to export the following environment variables:

* POSTGRES_DSN
  * Format: postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
  * Example: postgresql://super:123123@localhost:5432/awesome?key=value&sslmode=disable


* SOURCE_URL
  * Format: /path/to/migrations/folder/from/root. Default: ./
  * Example: /home/golang/project/migrations

### USAGE
migrator [command]

### Available Commands
* completion
  * Generate the autocompletion script for the specified shell
* down
  * Down migration operation that makes the database version the initially or corresponding to a given version
* help
  * Help about any command
* up
  * Up migration operation that makes the database version the latest or corresponding to a given version


# Golang meetup

This repo consists code example with information, how to make REST API service using Golang

The main idea is to make API for meetup registration with methods, such as:

1. Create member - `POST` request with query params name and email to create new member
2. Find - `GET` request with no params, to get all members
3. Find by id - `GET` request with id param, to get member data by id
4. Delete by id - `DELETE` request with id param, to remove member data by id

## Part 1

This directory consists first iteration of our code without any refactoring

## Part 2

This directory consists final version of our code divided into controllers and models 

## How to run

`cd part-*`

### Using docker

`docker-compose build`

`docker-compose up`

### Using Go

Firstly you need to run mongoDB on your local machine, then

`go build .`

`./part-*`




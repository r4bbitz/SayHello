# Go Clean Architechture


[![Build Status](http://img.shields.io/travis/prongbang/goclean.svg)](https://travis-ci.org/prongbang/goclean)
[![Codecov](https://img.shields.io/codecov/c/github/prongbang/goclean.svg)](https://codecov.io/gh/prongbang/goclean) 
[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/goclean)](https://goreportcard.com/report/github.com/prongbang/goclean)


### Install

```
go get -u github.com/r4bbitz/SayHello
```

### build
```
go build
```
### Run SayHello

```
Run SayHello.exe
```

### Install Newman  Automation test Cli
```
npm i newman
```


### Automation Test
newman run   automationtest/SayHelloAutomationTest.postman_collection.json -d automationtest/input.json

###URL
GET http://localhost:1991/api/v1/sayHello

###Expect Result test
```
automationtest/input.json
```

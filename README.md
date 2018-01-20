# CCII is serializer/deserialzier

[![Build Status](https://travis-ci.org/sirkon/ccii.svg?branch=master)](https://travis-ci.org/sirkon/ccii)

I don't know exactly why did I decide to give such a weird name. Just let it go.

##### Installation

```bash
go get -u github.com/sirkon/ccii
```
or
```bash
dep ensure -add github.com/sirkon/ccii
```

##### Purpose

The main purpose is to pass parameters into [protocol buffers](https://github.com/google/protobuf) compiler's plugins

##### Usage:

* Marshal
    ```go
    data, err := ccii.Marshal(obj)
    ```
* Unmarshal
    ```go
    err := ccii.Unmarshal(data, &dest)
    ```
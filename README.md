CCII is serializer/deserialzier

[![Build Status](https://travis-ci.org/sirkon/ccii.svg?branch=master)](https://travis-ci.org/sirkon/ccii)

I don't know exactly why should I decided to give such a weird name. So, let's ccii is a `ccii`.

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
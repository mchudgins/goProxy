# goProxy
[![Travis](https://travis-ci.org/mchudgins/goProxy.svg?branch=master)](https://travis-ci.org/mchudgins/goProxy)
trivial reverse proxy (written in go), does not support TLS

Usage:
```sh
reverseProxy -listen <:port> -target http://localhost:8000
```


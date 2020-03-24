# tikv-proxy

A gRPC proxy for [TiKV](https://tikv.org)


## Getting started

Please follow instructions in [TiKV official document](https://tikv.org/docs/3.0/tasks/try) to deploy a tiny TiKV cluster locally. Then run the proxy server:

```sh
$ docker run -it --rm --name tikv-proxy --network tikv -p 7788:7788 xiaogaozi/tikv-proxy:v0.0.1 --debug --pd-addrs pd.tikv:2379
```

Open another terminal to run the example:

```sh
$ make example
$ ./_build/example localhost:7788
```


## Development

```sh
$ make gen-proto  # Generate Go code of protocol buffers
$ make server     # Build server
$ docker build -t xiaogaozi/tikv-proxy:v0.0.1 -f ./deploy/Dockerfile . && docker rmi $(docker images -q -f dangling=true)  # Build Docker image
```

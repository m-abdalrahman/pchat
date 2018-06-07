# pchat: whisper-chat-example

## Running the example

The example assumes that there is a running Whisper v6 node. 

example for this, use `geth` with the following parameters to exposing an RPC interface at URL http://localhost:8545:

    $ geth <usual p2p flags> --shh --rpc


`--shh` is the option that enables Whisper v6 for the node.

`--rpc` enables the HTTP RPC interface.

## Download

download your version from [here](https://github.com/mabdalrahman/pchat/releases/tag/v1.0.0-alpha1)

## Installation From Source
`pchat` requires Go 1.10.2 or later, [dep](https://github.com/golang/dep) to managing dependency package and [xgo](https://github.com/karalabe/xgo) for cross compile.

```
$ go get -u github.com/mabdalrahman/pchat
```

## Usage
```
pchat -username "user" 

mandatory argument:
     -username                     set your username

optional arguments:
     -url                          set RPC URL for running Whisper v6 node default "http://localhost:8545"
     -privatekey                   set your private key

```

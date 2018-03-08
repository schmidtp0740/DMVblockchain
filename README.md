# DMV Blockchain

## Prereqs
 - have Docker and Docker-compose installed

## Description
Simple blockchain project that show how to transfer assets on the ledger. Transfer used vehicles onto the blockchain.

## Go Backend
- when cloning this repo, place in the following directory
``` 
$ mkdir -p $GOPATH/src/github/schmidtp0740 && cd $GOPATH/src/github.com/schmidtp0740
$ git clone github.com/schmidtp0740/DMVblockchain
```

- create an .env in DMVblockchain/backend with the address of Oracle Blockchain Cloud Service, like so
```
http://192.168.10.0:4001
```
# go-chord
Chord peer-to-peer distributed hash table protocol implemented in Go

[Wikipedia page explaining the chord protocol](https://en.wikipedia.org/wiki/Chord_(peer-to-peer))

This project is primarily for me to learn how to build distributed systems using the Go programming language.

### Things to do:
1. Implement Node type
2. Implement simple file system for writing to different nodes
3. Create Dockerfiles and docker-compose to ease deployment / dev environment
4. Build a cli for writing and reading files from the system

To compile proto files:
protoc -I supernode/ supernode/supernode.proto --go_out=plugins=grpc:supernode

protoc -I node/ node/node.proto --go_out=plugins=grpc:node

# go-chord
Chord peer-to-peer distributed hash table protocol implemented in Go

[Wikipedia page explaining the chord protocol](https://en.wikipedia.org/wiki/Chord_(peer-to-peer))

## Chord Entities:

### Supernode
There is only one supernode in the system. The supernode is in charge of
maintaining the state of the system. It keeps track of all the nodes in the
system and synchronizes node joins. When a client wants to access the system
for reading and writing, it contacts the supernode to get a random node in the
system.

### Node
The system contains many nodes. Each node has an Id, a finger table,
which is used for routing read and write requests to the correct nodes, and a
memory mapped file system for reading and writing files. The filesystem uses the
afero file system abstraction, which allows for different future implementations
such as a cache on read file system, or a copy on write file system.

## Key Concepts:

### Building the System
The supernode is used to orchestrate nodes joining the system. When a node wants
to join, it sends a join request to the supernode. The supernode will respond
with the list of nodes that are currently in the system. The node will then
update it's local finger table and will find it's predecessor in the list. Once
the node has taken care of it's local state, it will send a request to update
the distributed hash table of each of the nodes in the list. After the system
has been stabilized, the node will send a post join request to the supernode to
let it know the node has joined the system properly and that the system is in
a stable state.

### Routing
Routing in the system is handled by each of the nodes. Each node has a finger
table which is used to handle routing a request in O(log(n)) time. Routing works
by assigning each node a distinct ID. The finger table uses successor(ID + 2^i)
for keeping routing information of nodes across the system. This allows a node
to route requests quickly without having to keep track of every node in the
system. This means a node only has to keep track of 20 nodes when there are
1 million nodes in the system.

### Reading and Writing
Reading and Writing to the system is done by contacting the random node returned
from the supernode. When a request is made to the node, it hashes the filename,
checks if the file belongs to the node, if it does then it's read/written. If it
does not belong to the node, it's forwarded to the closest preceding node in the
node's finger table recursively.



### Things to do:
1. Use Cobra to build modern cli for go-chord
2. Create write and read for cli
2. Create Dockerfiles and docker-compose to ease deployment and integration testing

To compile proto files:
./protos.sh
Need to go into .pb.go files and change package names for correct importing

# Delve development

## State centric development

The state of the game should always be in the state that can be persisted and restored. This means that no state is outside the central state model.

## Fast iterations

To ensure fast iterations on the code base, we use air for live reloading on save.
This means that in order to keep the current place when live reloading the client needs to be able to load a snapshot of the state.
The best solution is probably the ability to save/load snapshots directly in the client. This will also be a good solution for replicating game states in bug reports.

## Architecture

Node based architecture

## Folder layout

src/scenes/ -> all individual scenes: players, bullets, monsters, items ...
src/levels/ -> ie. town, dungeon_level_1, ... 

## Missing nodes

* NetworkingNode
This node is responsible for networking. Both the client and server will need this node.
The node will need a root node reference for which it will traverse the scene tree. and replicate any changes to the authority.

* AudioNode


## Networking

Any scene tree that incoorporates networking needs to have a networking node. This node will synchronize the scene tree with the server (authority).

TODO:

* Meta data associated with each field
    * Authority
    * Replicated ( for networking )
    * Modified

* some indicator of who is the authority for a given field
    * is there some sort of trait/annotation functionality in golang ?
* determine serialization format: protobuf, ...
* on a side note, the serialization can also be used for snapshots  
* upon connecting to the server, the client will need to request the current state of the scene tree ( which is a full replication of the current state of the scene tree)
* game loop updates should only be changed fields, and fields that the client is allowed to see.
* we can use the node-path as part of the key in the replication map.
    * ie. /level1/players/player2 {
        name: "Vihaela",
        transform: {
            position: {
                x: 10,
                y: 12,
            }
            rotation: {
                x: 43,
                y: 0,
            }
        }
    }
* changeset generators to be able determine if values has changed
* constructors should only take primitive types that can be serialized in the creation function. 
    - All derived nodes references should be derived during a init fase. 
    - This is to ensure that the tree has been created before any querying of the tree is performed.
    - To enforce this constraint we can ensure that the game context is not created until after the level has been created.


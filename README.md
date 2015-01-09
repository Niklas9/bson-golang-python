BSON test with Go & Python services
===================================

Just doing some testing to communicate between a Go and a Python service using
BSON (http://bsonspec.org/, binary JSON) over Unix/TCP sockets.

To optimize network packet sizes, one should probably consider msgpack
(http://msgpack.org/) instead of BSON. BSON was created with data alteration
considered as well, and can worst case be even greater in byte size to send
over the network than JSON (due to the length prefixes and explicit array
indices).
BSON is the default format for MongoDB though, both for client-server
communication but also for storage - which means that if the received BSON
can be trusted we can quite safely just forward it to a MongoDB connection
without any deserialization etc.

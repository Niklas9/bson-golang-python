BSON test with Go & Python service
==================================

Just doing some testing to communicate between a Go and a Python service using
BSON (http://bsonspec.org/, binary JSON) over TCP.

To optimize network packet sizes, one should probably consider msgpack
(http://msgpack.org/) instead of BSON. BSON was created with data alteration
consideration as well, and can worst case be even greater in byte size to send
over the network than JSON.
BSON is the default format for MongoDB, both for client-server communication but
also for storage - in other words if the received BSON structure can be trusted
we can quite safely just forward it to a MongoDB connection without any
deserialization etc.


import socket

import bson


class BSONSocketConnection(object):

    socket = None
    host = None
    port = None

    def __init__(self, host, port):
        self.host = host
        self.port = port

    def connect(self):
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.connect((self.host, self.port))

    def close(self):
        self.socket.close()

    def send(self, data):
        d = bson.BSON.encode(data)
        self.socket.send(d)
        print 'sent type: %s, data: %r, length: %d' % (type(d), d, len(d))


if __name__ == '__main__':
    conn = BSONSocketConnection('localhost', 1337)
    conn.connect()
    conn.send({'message': 'hello yo'})
    #conn.close()

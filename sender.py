
import socket

import bson


class BSONSocketConnection(object):

    socket = None
    host = None
    port = None

    def __init__(self, host=None, port=None, socket_address=None):
        self.host = host
        self.port = port
        self.socket_address = socket_address
        t = socket.AF_INET
        if host is None:
            t = socket.AF_UNIX
        self.socket = socket.socket(t, socket.SOCK_STREAM)
        bson.patch_socket()  # brings native bson encoding/decoding to socket

    def connect(self):
        if self.host is None:
            self.socket.connect(self.socket_address)
        else:
            self.socket.connect((self.host, self.port))

    def close(self):
        self.socket.close()

    def send(self, data):
        self.socket.sendobj(data)
        d = bson.dumps(data)
        print 'sent type: %s, data: %r, length: %d' % (type(d), d, len(d))


if __name__ == '__main__':
    #conn = BSONSocketConnection(host='localhost', port=1337)
    conn = BSONSocketConnection(socket_address='/tmp/whatever.sock')
    conn.connect()
    conn.send({'text': 'hello', 'issent': True})
    conn.close()

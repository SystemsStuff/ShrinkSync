import os
from pyftpdlib.handlers import FTPHandler
from pyftpdlib.servers import FTPServer
from pyftpdlib.authorizers import DummyAuthorizer

authorizer = DummyAuthorizer()
authorizer.add_user('foo', 'bar', '.', perm='elradfmwMT')
authorizer.add_anonymous(os.getcwd())

handler = FTPHandler
handler.authorizer = authorizer

server = FTPServer(('0.0.0.0', 21), handler)
server.serve_forever()
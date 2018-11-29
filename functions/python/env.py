import os

def Main(req):
    return { "res": dict(os.environ) }, None

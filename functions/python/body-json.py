import json

def Main(req):
    o = json.loads(req.body.encode())
    return { "status": o["status"] }, None

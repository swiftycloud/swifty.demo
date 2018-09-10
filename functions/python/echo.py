def Main(req):
    print(req)
    print(req.args)
    try:
        print(req.claims)
    except:
        print("no claims")
    try:
        print(req.body)
    except:
        print("no body")
    return {"name2": req.args["name"] }, {"status": 201}

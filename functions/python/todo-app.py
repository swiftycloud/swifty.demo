import bson
import json
import swifty

#
# GET /tasks            -- list tasks
# POST /tasks $BODY     -- add new task
# GET /tasks/ID         -- get info about task
# PUT /tasks/ID         -- update task (except status)
# DELETE /tasks/ID      -- remove task
# POST /tasks/ID/done   -- mark task as done
#

def toTask(obj):
    return { 'id': str(obj['_id']), 'task': obj['task'], 'status': obj['status'] }

def fromTask(body, q):
    b = json.loads(body)
    if 'task' in b:
        q['task'] = b['task']

def Main(req):
    db = swifty.MongoDatabase('tasks')
    col = db['tasks']

    p = req.path.split('/')

    if p[0] != 'tasks':
        return {}, { 'status': 404 }

    q = { 'owner': req.claims['cookie'] }
    if len(p) == 1:
        if req.method == 'GET':
            if 'status' in req.args:
                q['status'] = req.args['status']
            return [ toTask(x) for x in col.find(q) ], None

        if req.method == 'POST':
            q['status'] = 'new'
            fromTask(req.body, q)
            col.insert_one(q)
            return {}, None

    q['_id'] = bson.ObjectId(p[1])
    if len(p) == 2:
        if req.method == 'GET':
            return toTask(col.find_one(q)), None

        if req.method == 'PUT':
            e = { }
            fromTask(req.body, e)
            col.update_one(q, { '$set': e })
            return {}, None

        if req.method == 'DELETE':
            col.delete_one(q)
            return {}, None

    if len(p) == 3:
        if p[2] == 'done' and req.method == 'POST':
            col.update_one(q, { '$set': { 'status': 'done' } })
            return {}, None

    return {}, { 'status': 404 }

import json
import os
import sys
import requests

TOKEN = '653041643:AAFYtZOS4-OLP2uzMSnUBFHHTKXUzmyvRhU'
BASE_URL = "https://api.telegram.org/bot{}".format(TOKEN)

def Main(req):
    try:
        data = json.loads(req.body.encode())
        message = str(data["message"]["text"])
        chat_id = data["message"]["chat"]["id"]
        first_name = data["message"]["chat"]["first_name"]

        response = "Please /start, {}".format(first_name)

        if "start" in message:
            response = "Hello {}! Type /help to get list of actions.".format(first_name)

        if "help" in message:
            response = "/about - get information about Swifty"

        if "about" in message:
            response = ("Swifty is the serverless platform that allows startups, developers and enterprises to develop and run application backend with minimal time-to-market, costs and without infrastructure management.\n"
            			"Start creating your backend at\n"
                       	"https://swifty.cloud")

        data = {"text": response.encode("utf8"), "chat_id": chat_id}
        url = BASE_URL + "/sendMessage"
        requests.post(url, data)

    except Exception as e:
        print(e)

    return {"statusCode": 200}, None

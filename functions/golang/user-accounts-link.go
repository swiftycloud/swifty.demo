/*
 * Funciton to add facebook ID to existing pwd-based account for
 * swifty Auth-as-a-Service
 * This function needs some mware and env setup:
 * - env:
 *      SWIFTY_AUTH_NAME=<name>
 * - mw:
 *      type:   mongo
 *      name:   <name>_mgo
 * - account:
 *      type:   facebook
 *      name:   <name>
 *      client: <client_id>
 *      secret: <secret_id>
 *
 * Also this FN needs <name>_jwt as its auth context
 *
 * API:
 *	/link
 *	-X POST
 *	-d {"code": <facebook-code>, "redirect_uri": <app-redirect-uri>}
 *	-H Authorization: bearer JWT
 */

package main

import (
	"os"
	"fmt"
	"swifty"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type linkResp struct {
	Error	string	`json:"error,omitempty"`
}

func doFacebookReq(url string) (map[string]interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error talking to facebook: %s", err.Error())
		return nil, err
	}

	ts, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Printf("Error reading resp body: %s", err.Error())
		return nil, err
	}

	var tok map[string]interface{}

	err = json.Unmarshal(ts, &tok)
	if err != nil {
		fmt.Printf("Error unmarshalling resp: %s", err.Error())
		return nil, err
	}

	return tok, nil
}

type facebookUserInfo struct {
	Id	string
}

type linkReq struct {
	Code	string	`json:"code"`
	URI	string	`json:"redirect_uri"`
}

func getCredsFromFacebook(n string, rq *linkReq) (*facebookUserInfo, error) {
	var url string

	url = "https://graph.facebook.com/v3.0/oauth/access_token"
	url += "?client_id=" + os.Getenv("ACC_FACEBOOK" + n + "_CLIENT")
	url += "&client_secret=" + os.Getenv("ACC_FACEBOOK" + n + "_SECRET")
	url += "&code=" + rq.Code
	url += "&redirect_uri=" + rq.URI

	toki, err := doFacebookReq(url)
	if err != nil {
		return nil, err
	}

	tok := toki["access_token"].(string)

	url = "https://graph.facebook.com/v3.0/me"
	url += "?access_token=" + tok
	url += "&fields=name,email"
	usr, err := doFacebookReq(url)
	if err != nil {
		return nil, err
	}

	return &facebookUserInfo{ Id: usr["id"].(string) }, nil
}


func Main(req *Request) (interface{}, *Responce) {
	if req.Method != "POST" || req.Path != "link" {
		return "Not found", &Responce{Status: 404}
	}

	aun := strings.ToUpper(os.Getenv("SWIFTY_AUTH_NAME"))

	db, err := swifty.MongoDatabase(aun + "_mgo")
	if err != nil {
		panic("Mongo DB not attached")
	}

	var rq linkReq

	err = json.Unmarshal([]byte(req.Body), &rq)
	if err != nil {
		fmt.Println(err)
		return &linkResp{Error: "Error unmarshalling body"}, nil
	}

	fui, err := getCredsFromFacebook(aun, &rq)
	if err != nil {
		fmt.Printf("Cannot talk to FB: %s\n", err.Error())
		return &linkResp{Error: "Error talking to facebook"}, nil
	}

	err = db.C("Users").Update(bson.M{"_id": bson.ObjectIdHex(req.Claims["cookie"].(string))},
			bson.M{"$set": bson.M{"facebookid": fui.Id}})
	if err != nil {
		fmt.Printf("Cannot link ids %s->%s: %s\n", req.Claims["cookie"], fui.Id, err.Error())
		return &linkResp{Error: "Error linking ids"}, nil
	}

	return &linkResp{}, nil
}

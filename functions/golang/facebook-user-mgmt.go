/*
 * Facebook account management for swifty Auth-as-a-Service
 * This function needs some mware and env setup:
 * - env:
 *      SWIFTY_AUTH_NAME=<name>
 * - mw:
 *      type:   mongo
 *      name:   <name>_mgo
 * - mw:
 *      type:   authjwt
 *      name:   <name>_jwt
 * - account:
 *      type:   facebook
 *      name:   <name>
 *      client: <client_id>
 *      secret: <secret_id>
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
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type authResp struct {
	Error	string	`json:"error,omitempty"`
	Token	string	`json:"token,omitempty"`
	Name	string	`json:"name,omitempty"`
	Email	string	`json:"email,omitempty"`
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
	Name	string
	Email	string
}

type signinReq struct {
	Code	string	`json:"code"`
	URI	string	`json:"redirect_uri"`
	Token	string	`json:"token"`
}

func getCredsFromFacebook(rq *signinReq) (*facebookUserInfo, error) {
	var url string

	if rq.Token == "" {
		n := strings.ToUpper(os.Getenv("SWIFTY_AUTH_NAME"))
		url = "https://graph.facebook.com/v3.0/oauth/access_token"
		url += "?client_id=" + os.Getenv("ACC_FACEBOOK" + n + "_CLIENT")
		url += "&client_secret=" + os.Getenv("ACC_FACEBOOK" + n + "_SECRET")
		url += "&code=" + rq.Code
		url += "&redirect_uri=" + rq.URI

		toki, err := doFacebookReq(url)
		if err != nil {
			return nil, err
		}

		rq.Token = toki["access_token"].(string)
	}

	url = "https://graph.facebook.com/v3.0/me"
	url += "?access_token=" + rq.Token
	url += "&fields=name,email"
	usr, err := doFacebookReq(url)
	if err != nil {
		return nil, err
	}

	ret := facebookUserInfo{ Id: usr["id"].(string) }
	if i, ok := usr["name"]; ok {
		ret.Name = i.(string)
	}
	if i, ok := usr["email"]; ok {
		ret.Email = i.(string)
	}

	return &ret, nil
}

func doSignin(auth *swifty.AuthCtx, rq *signinReq) interface{} {
	fbui, err := getCredsFromFacebook(rq)
	if err != nil {
		return &authResp{Error: "Error talking to facebook"}
	}

	urec := make(map[string]interface{})
	urec["facebookid"] = fbui.Id
	/* XXX Set additional facebbok data here */

	chg := mgo.Change{
		Upsert: true,
		Remove: false,
		Update: bson.M { "$setOnInsert": urec },
		ReturnNew: true,
	}

	_, err = auth.UsersCol.Find(bson.M{"facebookid": fbui.Id}).Apply(chg, &urec)
	if err != nil {
		fmt.Printf("Error signing up: %s", err.Error())
		return &authResp{Error: "Error signing in"}
	}

	jwt, err := auth.MakeJWT(map[string]interface{}{ "facebookid": fbui.Id, "cookie": urec["_id"] })
	if err != nil {
		return &authResp{Error: "Error signing in"}
	}

	return &authResp{Token: jwt, Name: fbui.Name, Email: fbui.Email}
}

func Main(req *Request) (interface{}, *Responce) {
	if req.Method != "POST" || req.Path != "signin" {
		return "Not found", &Responce{Status: 404}
	}

	auth, err := swifty.AuthContext()
	if err != nil {
		fmt.Println(err)
		panic("Can't get auth context")
	}

	var rq signinReq

	err = json.Unmarshal([]byte(req.Body), &rq)
	if err != nil {
		fmt.Println(err)
		panic("Cannot unmarshal body")
	}

	return doSignin(auth, &rq), nil
}

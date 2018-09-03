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
}

func doFacebookReq(url string) (map[string]string, error) {
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

	var tok map[string]string

	err = json.Unmarshal(ts, &tok)
	if err != nil {
		fmt.Printf("Error unmarshalling resp: %s", err.Error())
		return nil, err
	}

	return tok, nil
}

func getCredsFromFacebook(args map[string]string) (string, error) {
	tok, ok := args["token"]
	if ok {
		goto got_token
	}

	n := strings.ToUpper(os.Getenv("SWIFTY_AUTH_NAME"))
	url := "https://graph.facebook.com/v3.0/oauth/access_token"
	url += "?client_id=" + os.Getenv("ACC_FACEBOOK" + n + "_CLIENT")
	url += "&client_secret=" + os.Getenv("ACC_FACEBOOK" + n + "_SECRET")
	url += "&code=" + args["code"]
	url += "&redirect_uri=" + args["redirect_uri"]

	toki, err := doFacebookReq(url)
	if err != nil {
		return "", err
	}

	tok = toki["access-token"]
got_token:
	url = "https://graph.facebook.com/v3.0/me"
	url += "?access_token=" + tok
	usr, err := doFacebookReq(url)
	if err != nil {
		return "", err
	}

	return usr["id"], nil
}

func doSignin(auth *swifty.AuthCtx, args map[string]string) interface{} {
	fbid, err := getCredsFromFacebook(args)
	if err != nil {
		return &authResp{Error: "Error talking to facebook"}
	}

	urec := make(map[string]interface{})
	urec["facebookid"] = fbid
	/* XXX Set additional facebbok data here */

	chg := mgo.Change{
		Upsert: true,
		Remove: false,
		Update: bson.M { "$setOnInsert": urec },
		ReturnNew: true,
	}

	_, err = auth.UsersCol.Find(bson.M{"facebookid": fbid}).Apply(chg, &urec)
	if err != nil {
		fmt.Printf("Error signing up: %s", err.Error())
		return &authResp{Error: "Error signing in"}
	}

	jwt, err := auth.MakeJWT(map[string]interface{}{ "facebookid": fbid, "cookie": urec["_id"] })
	if err != nil {
		return &authResp{Error: "Error signing in"}
	}

	return &authResp{Token: jwt}
}

func Main(req *Request) (interface{}, *Responce) {
	auth, err := swifty.AuthContext()
	if err != nil {
		fmt.Println(err)
		panic("Can't get auth context")
	}

	switch req.Path {
	case "signin":
		return doSignin(auth, req.Args), nil
	}

	return "Invalid action", nil
}

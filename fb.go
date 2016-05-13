package go_social

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"log"
	"github.com/rizkisunaryo/go_social/model"
	"strings"
	"github.com/rizkisunaryo/go_http_request"
	"github.com/rizkisunaryo/go_json"
)

func HttpHandleFbLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Header.Set("Connection", "close")
	defer w.Header().Set("Connection", "close")

	header := r.RemoteAddr + ": go_social.HttpHandleFbLogin: "

	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err!=nil {
		log.Println(header, "1:", err.Error())
		fmt.Fprintf(w, "{\"Message\":\""+strings.Replace(err.Error(),"\"","",-1)+"\"}")
		return
	}

	log.Println(header,string(body))



	var req model.FbLoginReq
	json.Unmarshal(body, &req)

	if req.Id=="" || req.Token=="" {
		log.Println(header, "2")
		fmt.Fprintf(w, "{\"Message\":\"incomplete request\"}")
		return
	}



	var out model.FbLoginOut
	body,err = go_http_request.GetInterface("https://graph.facebook.com/"+req.Id+"?access_token="+req.Token, &out, 30)
	if err!=nil {
		log.Println(header, "3", err.Error())
		fmt.Fprintf(w, "{\"Message\":\""+strings.Replace(err.Error(),"\"","",-1)+"\"}")
		return
	}

	var resp model.FbLoginResp
	if out.Message!="" || out.Type!="" {
		resp.Status="11"
		resp.Message=out.Message
		resp.Type=out.Type
		resp.Code=out.Code
	} else {
		resp.Status="0"
	}



	go_json.PrintJson(header,w,resp)
}

func HttpHandleFbLogin2() (
string,
func(string, http.ResponseWriter, *http.Request),
func(string, http.ResponseWriter, *http.Request),
func(string, http.ResponseWriter, *http.Request),
func(string, http.ResponseWriter, *http.Request)) {

	return "go_social.HttpHandleFbLogin", handleFbLogin, handleFbLogin, handleFbLogin, handleFbLogin
}

func handleFbLogin(header string, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err!=nil {
		log.Println(header, "1:", err.Error())
		fmt.Fprintf(w, "{\"Message\":\""+strings.Replace(err.Error(),"\"","",-1)+"\"}")
		return
	}

	log.Println(header,string(body))



	var req model.FbLoginReq
	json.Unmarshal(body, &req)

	if req.Id=="" || req.Token=="" {
		log.Println(header, "2")
		fmt.Fprintf(w, "{\"Message\":\"incomplete request\"}")
		return
	}



	var out model.FbLoginOut
	body,err = go_http_request.GetInterface("https://graph.facebook.com/"+req.Id+"?access_token="+req.Token, &out, 30)
	if err!=nil {
		log.Println(header, "3", err.Error())
		fmt.Fprintf(w, "{\"Message\":\""+strings.Replace(err.Error(),"\"","",-1)+"\"}")
		return
	}

	var resp model.FbLoginResp
	if out.Message!="" || out.Type!="" {
		resp.Status="11"
		resp.Message=out.Message
		resp.Type=out.Type
		resp.Code=out.Code
	} else {
		resp.Status="0"
	}



	go_json.PrintJson(header,w,resp)
}
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	serviceUrl string
	// TODO(gj): configurable logging?
}

func NewClient() Client {
	return Client{serviceUrl: "http://localhost:8080"}
}

func NewClientWithServiceUrl(serviceUrl string) Client {
	return Client{serviceUrl}
}

type Instance interface {
	Id() string
	Type() string
}

type Type interface {
	Type() string
}

type AuthorizeReq struct {
	ActorType    string `json:"actor_type"`
	ActorId      string `json:"actor_id"`
	Action       string `json:"action"`
	ResourceType string `json:"resource_type"`
	ResourceId   string `json:"resource_id"`
}

type AuthorizeRes struct {
	allowed bool
}

func (c Client) Authorize(actor Instance, action string, resource Instance) (bool, error) {
	payload := AuthorizeReq{
		ActorType:    actor.Type(),
		ActorId:      actor.Id(),
		Action:       action,
		ResourceType: resource.Type(),
		ResourceId:   resource.Id(),
	}

	reqBodyJson, e := json.Marshal(payload)
	if e != nil {
		return false, e
	}

	reqBody := bytes.NewBuffer(reqBodyJson)
	res, e := http.Post(c.serviceUrl+"/authorize", "application/json", reqBody)
	if e != nil {
		return false, e
	}
	defer res.Body.Close()

	resBodyJson, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return false, e
	}
	if res.StatusCode != 200 {
		return false, errors.New(string(resBodyJson))
	}

	var resBody AuthorizeRes
	e = json.Unmarshal(resBodyJson, &resBody)
	if e != nil {
		return false, e
	}
	return resBody.allowed, nil
}

type ListReq struct {
	ActorType    string `json:"actor_type"`
	ActorId      string `json:"actor_id"`
	Action       string `json:"action"`
	ResourceType string `json:"resource_type"`
}

type ListRes struct {
	results []int
}

func (c Client) List(actor Instance, action string, resource Type) ([]int, error) {
	payload := ListReq{
		ActorType:    actor.Type(),
		ActorId:      actor.Id(),
		Action:       action,
		ResourceType: resource.Type(),
	}

	reqBodyJson, e := json.Marshal(payload)
	if e != nil {
		return nil, e
	}

	reqBody := bytes.NewBuffer(reqBodyJson)
	res, e := http.Post(c.serviceUrl+"/list", "application/json", reqBody)
	if e != nil {
		return nil, e
	}
	defer res.Body.Close()

	resBodyJson, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return nil, e
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(resBodyJson))
	}

	var resBody ListRes
	e = json.Unmarshal(resBodyJson, &resBody)
	if e != nil {
		return nil, e
	}
	return resBody.results, nil
}

type RelationReq struct {
	FromId   string `json:"from_id"`
	FromType string `json:"from_type"`
	Relation string `json:"relation"`
	ToId     string `json:"to_id"`
	ToType   string `json:"to_type"`
}

func (c Client) AddRelation(from Instance, name string, to Instance) error {
	reqBody := RelationReq{
		FromId:   from.Id(),
		FromType: from.Type(),
		Relation: name,
		ToId:     to.Id(),
		ToType:   to.Type(),
	}
	reqBodyJson, e := json.Marshal(reqBody)
	if e != nil {
		return e
	}
	reqBodyBytes := bytes.NewBuffer(reqBodyJson)
	res, e := http.Post(c.serviceUrl+"/relations", "application/json", reqBodyBytes)
	if e != nil {
		return e
	}
	defer res.Body.Close()
	resBody, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return e
	}
	if res.StatusCode != 200 {
		return errors.New(string(resBody))
	}
	return nil
}

func (c Client) DeleteRelation(from Instance, name string, to Instance) error {
	reqBody := RelationReq{
		FromId:   from.Id(),
		FromType: from.Type(),
		Relation: name,
		ToId:     to.Id(),
		ToType:   to.Type(),
	}
	reqBodyJson, e := json.Marshal(reqBody)
	if e != nil {
		return e
	}
	reqBodyBytes := bytes.NewBuffer(reqBodyJson)
	req, e := http.NewRequest("DELETE", c.serviceUrl+"/relations", reqBodyBytes)
	if e != nil {
		return e
	}
	req.Header.Set("Content-Type", "application/json")
	res, e := http.DefaultClient.Do(req)
	if e != nil {
		return e
	}
	defer res.Body.Close()
	resBody, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return e
	}
	if res.StatusCode != 200 {
		return errors.New(string(resBody))
	}
	return nil
}

type Role struct {
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	Role         string `json:"role"`
	ActorId      string `json:"actor_id"`
	ActorType    string `json:"actor_type"`
}

func (c Client) AddRole(resource Instance, name string, actor Instance) error {
	reqBody := Role{
		ResourceId:   resource.Id(),
		ResourceType: resource.Type(),
		Role:         name,
		ActorId:      actor.Id(),
		ActorType:    actor.Type(),
	}
	reqBodyJson, e := json.Marshal(reqBody)
	if e != nil {
		return e
	}
	reqBodyBytes := bytes.NewBuffer(reqBodyJson)
	res, e := http.Post(c.serviceUrl+"/roles", "application/json", reqBodyBytes)
	if e != nil {
		return e
	}
	defer res.Body.Close()
	resBody, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return e
	}
	if res.StatusCode != 200 {
		return errors.New(string(resBody))
	}
	return nil
}

func (c Client) DeleteRole(resource Instance, name string, actor Instance) error {
	reqBody := Role{
		ResourceId:   resource.Id(),
		ResourceType: resource.Type(),
		Role:         name,
		ActorId:      actor.Id(),
		ActorType:    actor.Type(),
	}
	reqBodyJson, e := json.Marshal(reqBody)
	if e != nil {
		return e
	}
	reqBodyBytes := bytes.NewBuffer(reqBodyJson)
	req, e := http.NewRequest("DELETE", c.serviceUrl+"/roles", reqBodyBytes)
	if e != nil {
		return e
	}
	req.Header.Set("Content-Type", "application/json")
	res, e := http.DefaultClient.Do(req)
	if e != nil {
		return e
	}
	defer res.Body.Close()
	resBody, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return e
	}
	if res.StatusCode != 200 {
		return errors.New(string(resBody))
	}
	return nil
}

// TODO(gj): Do we need equivalent of Oso::Client::get_roles in Ruby client?
func (c Client) GetResourceRoleForActor(resource Instance, role string, actor Instance) ([]Role, error) {
	req, e := http.NewRequest("GET", c.serviceUrl+"/roles", nil)
	if e != nil {
		return nil, e
	}
	q := req.URL.Query()
	q.Set("actor_type", actor.Type())
	q.Set("actor_id", actor.Id())
	q.Set("role", role)
	q.Set("resource_type", resource.Type())
	q.Set("resource_id", resource.Id())
	req.URL.RawQuery = q.Encode()
	res, e := http.DefaultClient.Do(req)
	if e != nil {
		return nil, e
	}
	defer res.Body.Close()
	resBodyJson, e := ioutil.ReadAll(res.Body)
	if e != nil {
		return nil, e
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(resBodyJson))
	}
	var resBody []Role
	e = json.Unmarshal(resBodyJson, &resBody)
	if e != nil {
		return nil, e
	}
	return resBody, nil
}

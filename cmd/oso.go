package main

import (
	"fmt"
	"log"

	oso "github.com/osohq/go-oso"
)

type User struct {
	id int
}

func (u User) Id() string {
	return fmt.Sprint(u.id)
}

func (u User) Type() string {
	return "User"
}

type Repo struct {
	id int
}

func (r Repo) Id() string {
	return fmt.Sprint(r.id)
}

func (r Repo) Type() string {
	return "Repo"
}

func main() {
	oso := oso.NewClient("http://localhost:8080")
	allowed, e := oso.Authorize(User{id: 1}, "read", Repo{id: 2})
	if e != nil {
		log.Fatalln(e)
	}
	log.Printf("Authorize: %v\n", allowed)

	results, e := oso.List(User{id: 1}, "read", Repo{})
	if e != nil {
		log.Fatalln(e)
	}
	log.Printf("List: %v\n", results)

	e = oso.AddRelation(Repo{id: 2}, "parent", Repo{id: 3})
	if e != nil {
		log.Fatalln(e)
	}
	log.Printf("AddRelation: success")

	e = oso.DeleteRelation(Repo{id: 2}, "parent", Repo{id: 3})
	if e != nil {
		log.Fatalln(e)
	}
	log.Printf("DeleteRelation: success")

	e = oso.AddRole(Repo{id: 2}, "member", User{id: 1})
	if e != nil {
		log.Fatalln(e)
	}
	log.Printf("AddRole: success")

	roles, e := oso.GetResourceRoleForActor(Repo{id: 2}, "member", User{id: 1})
	if e != nil {
		log.Fatalln(e)
	}
	log.Printf("GetResourceRoleForActor: %v", roles)

	e = oso.DeleteRole(Repo{id: 2}, "member", User{id: 1})
	if e != nil {
		log.Fatalln(e)
	}
	log.Printf("DeleteRole: success")
}

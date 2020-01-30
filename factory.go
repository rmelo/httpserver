package main

import "github.com/rmelo/httpserver/entity"

//NewMember creates an instance of entity.Member
func NewMember() interface{} {
	return &entity.Member{}
}

//NewTime creates an instance of entity.Member
func NewTime() interface{} {
	return &entity.Time{}
}

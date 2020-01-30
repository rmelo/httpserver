package main

import "github.com/rmelo/httpserver/entity"

func NewMember() interface{} { return &entity.Member{} }
func NewTime() interface{}   { return &entity.Time{} }

package main

import (
	"context"
	"fmt"
)

type Values struct {
	m map[string]string
}

func (v Values) Get(key string) string {
	return v.m[key]
}

func main() {
	v := Values{map[string]string{
		"1": "one",
		"2": "two",
	}}

	// store context with ctx["myvalues"]=v
	c2 := context.WithValue(context.Background(), "myvalues", v)

	// get value: ctx.Value["myvalues"]
	// fmt.Println(c2.Value("myvalues").(Values).Get("1"))
	vv := c2.Value("myvalues").(Values)
	for _, v1 := range vv.m {
		fmt.Println("v", v1)
	}
}

# go-stringify

[![Build Status](https://img.shields.io/travis/vicanso/go-stringify.svg?label=linux+build)](https://travis-ci.org/vicanso/go-stringify)


Stringify function for golang, it just like javascript's `JSON.stringify`. It supports customized replacement function, just like converting password to "***", cutting the string of which length is longer than 30.

## API

`String(interface{}, Replacer)`

- `interface{}` data to json stringify 
- `Replacer` function for customized replacement

```go
Replacer func(key string, value interface{}) (replace bool, newValue string)
```


```go
package main

import (
	"fmt"

	stringify "github.com/vicanso/go-stringify"
)

type (
	Other struct {
		Name        string                 `json:"name,omitempty"`
		Age         int                    `json:"age,omitempty"`
		VIP         bool                   `json:"vip,omitempty"`
		Information map[string]interface{} `json:"information,omitempty"`
	}
	Me struct {
		ID          string                 `json:"id,omitempty"`
		Account     string                 `json:"account,omitempty"`
		Password    string                 `json:"password,omitempty"`
		Secret      string                 `json:"-"`
		Name        string                 `json:"name,omitempty"`
		Age         int                    `json:"age,omitempty"`
		VIP         bool                   `json:"vip,omitempty"`
		Detail      string                 `json:"detail,omitempty"`
		Information map[string]interface{} `json:"information,omitempty"`
		Boss        Other                  `json:"boss,omitempty"`
		Leader      *Other                 `json:"leader,omitempty"`
		Workmates   []Other                `json:"workmates,omitempty"`
		Friends     []*Other               `json:"friends,omitempty"`
		Nil         []*Other               `json:"nil,omitempty"`
	}
)

func main() {
	s := getData()
	// {"account":"t***t","password":"***","name":"foo","age":18,"detail":"GitHub is where people build s...","information":{"male":true,"bestFriend":{"name":"peter"},"fav":"reading","weight":70},"boss":{"name":"boss"},"leader":{"age":30},"workmates":[{"name":"jack"},{"name":"tas"}],"friends":[{"name":"tom"}]}
	fmt.Println(stringify.String(s, replacer))
}

func replacer(key string, value interface{}) (bool, string) {
	// password is hidden by ***
	if key == "password" {
		return true, `"***"`
	}
	// mask the account
	if key == "account" {
		v := value.(string)
		return true, `"` + v[0:1] + "***" + v[len(v)-1:] + `"`
	}
	str, ok := value.(string)
	// cut the string
	if ok && len(str) > 30 {
		return true, `"` + str[0:30] + `..."`
	}
	// others not change
	return false, ""
}

func getData() *Me {
	friends := make([]*Other, 2)
	friends = append(friends, &Other{
		Name: "tom",
	})
	peter := &Other{
		Name: "peter",
	}
	return &Me{
		ID:       "",
		Account:  "test",
		Password: "pwd",
		Secret:   "none",
		Name:     "foo",
		Age:      18,
		VIP:      false,
		Detail:   "GitHub is where people build software. More than 31 million people use GitHub to discover, fork, and contribute to over 100 million projects.",
		Information: map[string]interface{}{
			"fav":        "reading",
			"weight":     70,
			"male":       true,
			"bestFriend": peter,
		},
		Boss: Other{
			Name: "boss",
		},
		Leader: &Other{
			Age: 30,
		},
		Workmates: []Other{
			Other{
				Name: "jack",
			},
			Other{
				Name: "tas",
			},
		},
		Friends: friends,
	}
}
```

## benchmark

```bash
goos: darwin
goarch: amd64
pkg: github.com/vicanso/go-stringify
BenchmarkStringify-4   	   50000	     23350 ns/op	    3320 B/op	     144 allocs/op
BenchmarkMarshal-4     	  300000	      5548 ns/op	    1504 B/op	      16 allocs/op
PASS
ok  	github.com/vicanso/go-stringify	3.146s
```
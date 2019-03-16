package stringify

import (
	"encoding/json"
	"strings"
	"testing"
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

func getTestData() *Me {
	friends := make([]*Other, 2)
	friends = append(friends, &Other{
		Name: "tom",
	})
	peter := &Other{
		Name: "peter\" don\"",
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

func replacer(key string, value interface{}) (bool, string) {
	if key == "password" {
		return true, `"***"`
	}
	if key == "account" {
		v := value.(string)
		return true, `"` + v[0:1] + "***" + v[len(v)-1:] + `"`
	}
	str, ok := value.(string)
	if ok && len(str) > 30 {
		return true, `"` + str[0:30] + `..."`
	}
	return false, ""
}

func TestString(t *testing.T) {
	s := getTestData()
	str := String(s, replacer)
	checkList := []string{
		`"account":"t***t"`,
		`"password":"***"`,
		`"name":"foo"`,
		`"age":18`,
		`"detail":"GitHub is where people build s..."`,
		`"fav":"reading"`,
		`"weight":70`,
		`"male":true`,
		`"bestFriend":{"name":"peter\" don\""}`,
		`"boss":{"name":"boss"}`,
		`"leader":{"age":30}`,
		`{"name":"jack"}`,
		`{"name":"tas"}`,
		`{"name":"tom"}`,
	}
	if len(str) != 312 {
		t.Fatalf("json stringify fail")
	}
	for _, v := range checkList {
		if !strings.Contains(str, v) {
			t.Fatalf("json stringify fail, it should contains %s", v)
		}
	}
}

func BenchmarkString(b *testing.B) {
	b.ReportAllocs()
	s := getTestData()

	for i := 0; i < b.N; i++ {
		String(s, nil)
	}
}

func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	s := getTestData()
	for i := 0; i < b.N; i++ {
		buf, _ := json.Marshal(s)
		_ = string(buf)
	}
}

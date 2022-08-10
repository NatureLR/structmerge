package structmerge

import (
	"testing"
)

type Main struct {
	Fa string
	Fb string
	Fc []string
	Fx string `merge:"fs"`
}

type A struct {
	Fa string
}

type B struct {
	Fb string
	FF string `merge:"fs"`
	Fc C
}

type C struct {
	Fc []string
}

func TestMain(t *testing.T) {
	var main Main
	a := A{Fa: "fa_value"}
	b := B{Fb: "fb_value", FF: "fffffff", Fc: C{[]string{"fc is a struct"}}}

	err := Merge(&main, a, b)
	if err != nil {
		t.Log(err)
		return
	}
	if main.Fa != "fa_value" || main.Fb != "fb_value" || main.Fx != "fffffff" {
		t.Log("test faild")
	}
	t.Log(main)
}

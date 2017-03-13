package main

import "github.com/kintone/go-kintone"

func newApp() *kintone.App {
	return &kintone.App{
		Domain:   "r0pri.cybozu.com",
		ApiToken: "orHrvJ3Y8Urgpdz1MhMdCV1bJRQfT1U34ZMFTdbz",
		AppId:    61,
	}
}

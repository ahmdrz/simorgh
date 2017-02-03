package main

import (
	"go-srp/src/srp"
	"sync"
)

type Authentication struct {
	mutex        sync.Mutex
	keys         map[string]*srp.Server
	accepts      map[string]int64
	salt         []byte
	verifier     []byte
	identityHash []byte
}

func InitAuthentication() *Authentication {
	return &Authentication{
		keys:    make(map[string]*srp.Server),
		accepts: make(map[string]int64),
	}
}

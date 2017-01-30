package main

import (
	"go-srp/src/srp"
)

type Authentication struct {
	keys         map[string]*srp.Server
	accepts      map[string]bool
	salt         []byte
	verifier     []byte
	identityHash []byte
}

func InitAuthentication() *Authentication {
	return &Authentication{
		keys:    make(map[string]*srp.Server),
		accepts: make(map[string]bool),
	}
}

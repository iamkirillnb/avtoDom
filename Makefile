SHELL := /bin/bash

gen-go:
	sh ./gen-go.sh


build: gen-go

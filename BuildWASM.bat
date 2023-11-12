@echo off
set GOOS=js
set GOARCH=wasm
go build -o game.wasm .
set GOOS=
set GOARCH=
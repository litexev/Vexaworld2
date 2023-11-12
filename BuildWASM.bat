@echo off
set GOOS=js
set GOARCH=wasm
go build -o game.wasm .
set GOOS=
set GOARCH=
copy game.wasm .\web\game.wasm
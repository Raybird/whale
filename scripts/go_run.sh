#! /bin/bash

# container 裡面用 air 來啟動開發時期的 gin 服務,來達成 live-reloading
# https://github.com/cosmtrek/air

# clone and install 
go mod download

# https://github.com/cosmtrek/air
# ./air
# go run cmd/web/main.go

/app/.air
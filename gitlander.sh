#!/bin/bash

case $1 in

"register" )
    publickey=$(cat $3)
    curl "http://gitlander.com/api/user?name=$2" -H "Content-Type: text/plain;charset=UTF-8" --data-binary $"$publickey" --compressed
;;

"init" )
    curl "http://gitlander.com/api/repo?name=$2&repo=$3"
;;

*)
echo "--help commands:"
echo "register <username> <ssh public key file>"
echo "init <username> <repo name>"
echo "example clone:"
echo "git clone <username>@gitlander:<repo>.git"
esac
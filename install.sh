#!/bin/bash
[ "$UID" -eq 0 ] || { echo "This script must be run as root."; exit 1;}

unamestr="$(uname -s)"
email=""
email_confirm=""
case "${unamestr}" in
    Linux*)
        go build -o ./bin/btop
        sudo mv ./bin/btop /usr/local/bin
        while [ -z $email ] || [ $email != $email_confirm ]
        do
            read -p 'Email adress of your commits : ' email
            read -p 'Confirm your email : ' email_confirm
        done
        echo $email > .btop_config
        mv .btop_config ~/
        echo "Btop is available on your computer. Try \"btop\" for start";;
    Darwin*)
        go build -o ./bin/btop
        sudo mv ./bin/btop /usr/local/bin
        while [ -z $email ] || [ $email != $email_confirm ]
        do
            read -p 'Email adress of your commits : ' email
            read -p 'Confirm your email : ' email_confirm
        done
        echo $email > .btop_config
        mv .btop_config ~/
        echo "Btop is available on your computer. Try \"btop\" for start 🚀";;
    *)          echo "Your system is not compatible"
esac

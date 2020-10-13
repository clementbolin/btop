#!/bin/bash
if [[ $EUID -ne 0 ]]; then
    echo "You must run this with superuser priviliges.  Try \"sudo ./install.sh\"" 2>&1
    exit 1
else
    echo "Installing btop..."
fi

function intall_linux() {
    if go build -o ./bin/btop &> logs ; then
        email = ""
        email_confirm = ""
        sudo mv ./bin/btop /usr/local/bin
        while [ -z $email ] || [ $email != $email_confirm ]
        do
            read -p 'Email adress of your commits : ' email
            read -p 'Confirm your email : ' email_confirm
        done
        echo $email > .btop_config
        mv .btop_config ~/
        rm ./logs
        echo "Btop is available on your computer. Try \"btop\" for start";
    fi
    echo "Your compil has failed";
    tail -n 10 logs
}

function intall_unix() {
    if go build -o ./bin/btop &> logs ; then
        email=""
        email_confirm=""
        sudo mv ./bin/btop /usr/local/bin
        while [ -z $email ] || [ $email != $email_confirm ]
        do
            read -p 'Email adress of your commits : ' email
            read -p 'Confirm your email : ' email_confirm
        done
        echo $email > .btop_config
        mv .btop_config ~/
        rm ./logs
        echo "Btop is available on your computer. Try \"btop\" for start 🚀";
    fi
    echo "Your compil has failed";
    tail -n 10 logs
}

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     intall_linux;;
    Darwin*)    intall_unix;;
    *)    echo "Project not avaible on your system";;
esac

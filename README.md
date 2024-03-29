# Btop
![Go](https://github.com/ClementBolin/btop/workflows/Go/badge.svg)
[![Build Status](https://travis-ci.org/ClementBolin/btop.svg?branch=master)](https://travis-ci.org/ClementBolin/btop)

Btop is your personal terminal dashboard.
The goal of this project, is to display many information of your system
* Process run
* Docker information (container, system)
* Uptime
* Power charge
* Git statistics
* System information (os, version)
* Command history 🚀


#### Example

![Btop example](./assets/btop.gif)


#### Feature

- [x] Process run
- [x] Docker information (container, system)
- [x] Uptime
- [x] Power charge
- [x] Git statistics
- [x] System information (os, version)
- [x] Command history
- [ ] Choice your module
- [ ] Choice position of each module
- [ ] Github Module
- [ ] Trello module
- [ ] Updater

#### How To Install
if you don't have the Golang compiler, you can download it [here](https://golang.org)

    git clone https://github.com/ClementBolin/btop
    cd btop
    chmod +x ./install.sh
    sudo ./install.sh

#### How to run with docker

    git clone https://github.com/ClementBolin/btop
    cd btop
    docker build -t btop . && docker run -it btop /bin/bash
    ./bin/btop

This project is based on [WTF](https://github.com/wtfutil/wtf) and use [tview](https://github.com/rivo/tview) and [tcell](https://github.com/gdamore/tcell) library

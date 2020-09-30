FROM golang:1.15

LABEL version="1.0" maintainer="Clement Bolin <clement.bolin@epitech.eu>"

ARG workdir="btop"

WORKDIR /${workdir}

COPY . .

RUN apt-get update -y && apt-get install git 
RUN git clone https://github.com/ClementBolin/btop && cd btop
RUN chmod +x ./install.sh

CMD ["./install.sh"]

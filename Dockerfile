FROM golang:1.15

LABEL version="1.0" maintainer="Clement Bolin <clement.bolin@epitech.eu>"

ARG workdir="top-go"

WORKDIR /${workdir}

RUN apt-get update -y

# Active when project is end

# COPY . .
# RUN apt-get update -y && apt-get install git 
# RUN git clone https://github.com/ClementBolin/top-go.git && cd top-go
# RUN make build 

# CMD ["./bin/${workdir}"]

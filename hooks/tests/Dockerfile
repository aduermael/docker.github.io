#
# Docs validator
#

FROM golang:1.7.3-alpine
MAINTAINER Gaetan de Villele <gaetan@docker.com>

# -----------------------
# INSTALL GIT
# -----------------------
RUN apk update
RUN apk add git

# -----------------------
# INSTALL GO PROGRAM
# -----------------------

COPY src /go/src

WORKDIR /go/src/validator
RUN go install

# when running the container, MOUNT docs repo in /docs

# -----------------------
# START SHELL
# -----------------------

CMD ash

FROM ubuntu:latest
LABEL authors="teten"

ENTRYPOINT ["top", "-b"]
FROM scratch
WORKDIR /
MAINTAINER Niclas Geiger

ADD main /

EXPOSE 443

CMD ["/main"]
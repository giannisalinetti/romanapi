FROM centos:7

RUN yum install -y golang && yum clean all -y
RUN useradd roman
USER roman

RUN mkdir -p /home/roman/go/{src,pkg,bin}
ENV GOPATH=/home/roman/go
RUN mkdir -p $GOPATH/src/github.com/gbs/romanapi

ADD romanserver $GOPATH/src/github.com/gbs/romanapi/romanserver
ADD romanNumerals $GOPATH/src/github.com/gbs/romanapi/romanNumerals
WORKDIR $GOPATH/src/github.com/gbs/romanapi/romanserver

RUN go install

EXPOSE 8000

CMD ["/home/roman/go/bin/romanserver"]

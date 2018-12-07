FROM centos:7
ADD goweb /opt/goweb
ADD goconf /opt/conf/goconf
ENTRYPOINT ["/opt/goweb"]


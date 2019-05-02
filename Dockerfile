FROM busybox
VOLUME /golang/gopath/src/github.com/aloxc/gappweb /root/gappweb
ENTRYPOINT ["/root/gappweb/gappweb"]
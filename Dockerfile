FROM alpine

COPY gopath/bin/kroger-app-test /go/bin/kroger-app-test

ENTRYPOINT /go/bin/kroger-app-test

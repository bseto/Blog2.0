FROM golang:1.17.5
WORKDIR /persistentchat
RUN go get github.com/bseto/blog2/backend
RUN go install github.com/bseto/blog2/backend/cmd/persistentchat@cd82a7a
RUN ls -alh ${GOPATH}/bin
CMD ${GOPATH}/bin/persistentchat
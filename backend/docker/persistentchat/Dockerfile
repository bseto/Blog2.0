FROM backend:latest
WORKDIR /backend/cmd/persistentchat
RUN go build
CMD ./persistentchat

#RUN go get github.com/bseto/blog2/backend
#RUN ls -alh ${GOPATH}/bin
#RUN go install github.com/bseto/blog2/backend/cmd/persistentchat@latest
#RUN ls -alh ${GOPATH}/bin
#CMD ${GOPATH}/bin/persistentchat
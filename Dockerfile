FROM golang:latest
ENV LANG=C.UTF-8
ARG app_name=GooooooooG
ARG http_port=9000

# Install beego & bee
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get github.com/json-iterator/go

# Expose port to public
EXPOSE $http_port

# Copy the source code from current directory to /go/src in container
# Place the Dockerfile into source code directory and build Docker image
RUN mkdir -p /go/src/$app_name
COPY . /go/src/$app_name
WORKDIR /go/src/$app_name

# Launch Beego when the container is created
CMD bee run
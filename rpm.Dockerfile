# ============================================================
# Compile binary
# ============================================================
FROM golang:1.12 AS buildenv
RUN apt-get update && apt-get install -y unzip

ARG PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
RUN unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
RUN unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
RUN rm -f $PROTOC_ZIP

RUN go get -u github.com/asdine/genji/...
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN echo $GOPATH
RUN ls -al $GOPATH/bin

WORKDIR /go/src/github.com/Megalithic-LLC/on-prem-emaild
COPY . .
RUN make generate && make


# ============================================================
# Build an RPM package
# ============================================================
FROM centos:7
RUN yum install -y automake gcc libffi-devel make rpm-build ruby-devel rubygems
RUN gem install --no-ri --no-rdoc fpm -v '1.11.0'

COPY emaild.service /usr/lib/systemd/system/on-prem-emaild.service
RUN mkdir -p /var/on-prem/emaild
COPY --from=buildenv /go/bin/on-prem-emaild /usr/bin/


# Create RPM package

WORKDIR /

ARG VERSION=0.0.0
ARG ITERATION=0

RUN fpm \
  --input-type dir \
  --output-type rpm \
  --name on-prem-emaild \
  --description "On-Prem Email" \
  --version ${VERSION} \
  --iteration ${ITERATION} \
  /usr/bin/on-prem-emaild \
  /usr/lib/systemd/system/on-prem-emaild.service \
  /var/on-prem/emaild

# Provide an easy means of copying the archive out of the image

VOLUME /output
CMD cp -v *.rpm /output

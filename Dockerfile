FROM grpc/go:latest
# RUN apt-get update && apt-get -y install apt-utils && apt-get install -y unzip && apt-get install wget && apt-get install -y curl 
# RUN apt-get clean

RUN export GOBIN=$GOPATH/bin && mkdir -p $GOPATH/src/tange
ADD . $GOPATH/src/tange
WORKDIR $GOPATH/src/tange
RUN cp ./.tange.yaml.sample ~/.tange.yaml
RUN make build 
EXPOSE 61613


RUN mkdir -p /bin/tange
RUN export PATH="$PATH:/bin/tange"
 
ADD tange /bin/tange

WORKDIR /bin/tange
# 
# EXPOSE 61613
# 
ENTRYPOINT ./tange serve


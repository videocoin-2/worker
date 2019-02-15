FROM debian:jessie-slim AS release

LABEL maintainer="Videocoin" description="transcoding client streams"

RUN apt update && apt upgrade -y
RUN apt install ca-certificates wget build-essential -y

WORKDIR /opt/

RUN wget https://johnvansickle.com/ffmpeg/builds/ffmpeg-git-amd64-static.tar.xz

RUN tar -xvf ffmpeg-git-amd64-static.tar.xz
RUN rm ffmpeg-git-amd64-static.tar.xz
RUN mv ffmpeg-*/* /usr/local/bin
RUN rm -rf ffmpeg*

ADD keys keys
ADD release/transcoder-linux-amd64 ./


EXPOSE 50051 50052 50053 50054 50055

ENTRYPOINT [ "./transcoder-linux-amd64" ]

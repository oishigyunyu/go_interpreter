# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
#RUN mkdir /go/src/work
## コンテナログイン時のディレクトリ指定
#WORKDIR /go/src/work
# ホストのファイルをコンテナの作業ディレクトリに移行



RUN apt-get update -qq \
        && apt-get install -y --no-install-recommends \
        sudo curl wget apt-transport-https gnupg \
        ca-certificates \
        gcc build-essential git vim 

ARG username=go_penguin
ARG wkdir=/home/${username}/work

# echo "username:password" | chpasswd
# root password is "root"

RUN echo "root:root" | chpasswd && \
    adduser --disabled-password --gecos "" "${username}" && \
    echo "${username}:${username}" | chpasswd && \
    echo "%${username}    ALL=(ALL)   NOPASSWD:    ALL" >> /etc/sudoers.d/${username} && \
    chmod 0440 /etc/sudoers.d/${username} 
    
WORKDIR ${wkdir}
RUN chown ${username}:${username} ${wkdir}
USER ${username}
RUN sudo chmod a+w /
ENV GOPATH $HOME/.go
ENV PATH $GOPATH/bin:$PATH

CMD ["bash"]

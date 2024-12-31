FROM alpine:latest

RUN apk add --no-cache perl wget xz tar

ADD texlive.profile /data/texlive.profile

RUN wget http://mirror.ctan.org/systems/texlive/tlnet/install-tl-unx.tar.gz && \
    tar -xzf install-tl-unx.tar.gz && \
    cd install-tl-* && \
    ./install-tl --profile /data/texlive.profile

ENV PATH="/usr/local/texlive/2024/bin/x86_64-linuxmusl:${PATH}"

RUN tlmgr install \
        pgf \
        xcolor \
        enumitem \
        titlesec \
        fontawesome \
        etoolbox \
        fira \
        xkeyval \
        fontaxes \
        fontspec

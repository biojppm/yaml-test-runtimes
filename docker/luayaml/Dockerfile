ARG FROM
FROM $FROM
ARG REPO
ARG TAG

RUN apk add \
        lua \
        lua-dev \
        yaml \
        yaml-dev \
 && true

RUN git clone --branch=v$TAG $REPO /build

WORKDIR /build

RUN build-aux/luke \
 && build-aux/luke install \
 && true

RUN cp -r \
        /usr/local/share/lua \
        /export/usr/local/share/ \
 && cp -r \
        /usr/local/lib/lua \
        /export/usr/local/lib/ \
 && mkdir -p /export/usr/bin \
 && cp /usr/bin/lua* /export/usr/bin/ \
 && mkdir -p /export/usr/lib \
 && cp -P /usr/lib/liblua* /export/usr/lib/ \
 && true

COPY bin/* /export/usr/local/bin/

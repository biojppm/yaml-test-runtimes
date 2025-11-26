ARG FROM
FROM $FROM
ARG REPO
ARG TAG

RUN apk update \
 && apk add \
        maven \
        musl \
        openjdk11 \
 && true

ENV JAVA_HOME=/usr/lib/jvm/java-11-openjdk
ENV PATH="$JAVA_HOME/bin:$PATH"

WORKDIR /build

RUN git init --quiet \
 && git remote add origin $REPO \
 && git fetch origin $TAG \
 && git reset --hard FETCH_HEAD \
 && true

RUN mvn clean install

COPY java /java

WORKDIR /java

RUN version=$TAG \
 && version=${version#snakeyaml-} \
 && mvn clean package -DskipTests -Dsnakeversion="$version" \
 && true

RUN mkdir -p /export/java \
 && mv target/snake2json-*-jar-with-dependencies.jar /export/java

COPY bin/* /export/usr/local/bin/

ARG FROM
FROM $FROM
ARG TAG

RUN apk add python3-dev

RUN pip install --break-system-packages pyyaml==$TAG

RUN mkdir -p /export/usr/lib/python3.12/site-packages \
 && cp -Pr \
        /usr/lib/python3.12/site-packages/* \
        /export/usr/lib/python3.12/site-packages/ \
 && true

COPY bin/* /export/usr/local/bin/

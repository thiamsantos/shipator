FROM openresty/openresty:1.17.8.1-alpine

ENV SHIPATOR_VERSION=0.1.0-rc2
ENV SHIPATOR_PLACEHOLDER=__ENV__
ENV SHIPATOR_TARGET=/app/shipator/html/index.html
ENV SHIPATOR_PREFIX=REACT_APP

RUN rm /etc/nginx/conf.d/default.conf
COPY default.conf /etc/nginx/conf.d/

WORKDIR /shipator
COPY start.sh /shipator/start.sh

RUN wget "https://github.com/brainnco/shipator/releases/download/v$SHIPATOR_VERSION/shipator-v$SHIPATOR_VERSION-linux-amd64.tar.gz" && \
    tar -xf "shipator-v$SHIPATOR_VERSION-linux-amd64.tar.gz" && \
    rm "shipator-v$SHIPATOR_VERSION-linux-amd64.tar.gz" && \
    mv shipator /usr/bin/shipator && \
    shipator -version

CMD ["/shipator/start.sh"]

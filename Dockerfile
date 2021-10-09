FROM golang:1.17.2

WORKDIR /go/src/code_sim

COPY . .

ENV CONFIG_PATH=/go/src/code_sim/config.yml
ENV ENV=prd
ENV PY_LEXICAL_ANALYZER_PATH=/go/src/code_sim/transformer/python-lexical-analyzer/analyze.py
ENV NETWORK_INTERFACE=eth2

CMD ["/bin/bash", "-c", "./out/code_sim"]
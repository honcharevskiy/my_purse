FROM python:3.8

SHELL ["/bin/bash", "-c"]
USER root

COPY . my_purse/
WORKDIR my_purse/
RUN pip install .

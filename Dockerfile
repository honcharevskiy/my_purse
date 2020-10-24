FROM python:3.8

SHELL ["/bin/bash", "-c"]
USER root

RUN apt-get update && apt-get install -y pipenv
COPY . my_purse/
WORKDIR my_purse/
RUN pip install .

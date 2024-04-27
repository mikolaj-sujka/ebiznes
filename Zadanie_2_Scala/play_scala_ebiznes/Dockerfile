FROM ubuntu:22.04

ENV SCALA_VERSION 2.13.13
ENV SBT_VERSION 1.9.9

WORKDIR /app
COPY . /app

RUN apt-get update && \
    apt-get install -y wget curl gnupg2 openjdk-11-jdk

RUN apt-get install -y openjdk-8-jdk

RUN wget https://downloads.lightbend.com/scala/${SCALA_VERSION}/scala-${SCALA_VERSION}.tgz && \
    tar -xzf scala-${SCALA_VERSION}.tgz && \
    mv scala-${SCALA_VERSION} /usr/share/scala && \
    rm scala-${SCALA_VERSION}.tgz
    
ENV PATH="/usr/share/scala/bin:${PATH}"

RUN echo "deb https://repo.scala-sbt.org/scalasbt/debian all main" | tee /etc/apt/sources.list.d/sbt.list && \
    echo "deb https://repo.scala-sbt.org/scalasbt/debian /" | tee -a /etc/apt/sources.list.d/sbt.list && \
    curl -sL "https://keyserver.ubuntu.com/pks/lookup?op=get&search=0x2EE0EA64E40A89B84B2DF73499E82A75642AC823" | apt-key add && \
    apt-get update && \
    apt-get install -y sbt=${SBT_VERSION}

WORKDIR /app
COPY . /app

CMD scala -version && sbt sbtVersion && sbt run
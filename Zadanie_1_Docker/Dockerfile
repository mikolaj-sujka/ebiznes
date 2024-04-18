FROM ubuntu:22.04

ENV PYTHON_VERSION=python3.8
ENV KOTLIN_VERSION=1.8.20
ENV JAVA_VERSION=8.0.345-zulu

# 3.0 - update package manager, install python and pip
# 3.5 - instalacja JAVA 8 i Kotlin
# 4.0 - instalacja Gradle i JDBC SQLite w ramach projektu na Gradle 

RUN ln -snf /usr/share/zoneinfo/Europe/Warsaw /etc/localtime && echo Europe/Warsaw > /etc/timezone
RUN apt-get update && apt-get install -y zip unzip curl
RUN apt-get install -y software-properties-common 
RUN add-apt-repository -y ppa:deadsnakes/ppa
RUN apt install -y ${PYTHON_VERSION} python3-pip openjdk-8-jdk 
RUN apt-get install -y wget unzip
RUN curl -s "https://get.sdkman.io" | bash && \
    /bin/bash -c "source $HOME/.sdkman/bin/sdkman-init.sh && \
    yes | sdk install java ${JAVA_VERSION} && \
    yes | sdk install kotlin"

ENV KOTLIN_HOME=$HOME/.sdkman/candidates/kotlin/current
ENV PATH=$PATH:$KOTLIN_HOME/bin

RUN wget https://services.gradle.org/distributions/gradle-8.7-bin.zip -P /tmp && \
    unzip -d /opt/gradle /tmp/gradle-8.7-bin.zip && \
    rm /tmp/gradle-8.7-bin.zip
ENV GRADLE_HOME=/opt/gradle/gradle-8.7
ENV PATH=$PATH:$GRADLE_HOME/bin

# Kopiowanie projektu do obrazu
COPY . /app
WORKDIR /app

# Budowanie projektu za pomocą Gradle
RUN /bin/bash -c "gradle build --no-daemon"

# Sprawdź wersję Pythona, Java, Kotlin, gradle
CMD python3.8 --version && java -version && kotlinc -version && gradle -version
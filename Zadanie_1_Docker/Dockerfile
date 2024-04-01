FROM ubuntu:22.04

ENV PYTHON_VERSION=python3.8

# 3.0 - update package manager, install python and pip
# 3.5 - instalacja JAVA 8 i Kotlin
# 4.0 - instalacja Gradle
RUN ln -snf /usr/share/zoneinfo/Europe/Warsaw /etc/localtime && echo Europe/Warsaw > /etc/timezone
RUN apt-get update 
RUN apt-get install -y software-properties-common 
RUN add-apt-repository -y ppa:deadsnakes/ppa
RUN apt install -y ${PYTHON_VERSION} python3-pip openjdk-8-jdk 
RUN apt-get install -y wget unzip
RUN wget https://github.com/JetBrains/kotlin/releases/download/v1.6.10/kotlin-compiler-1.6.10.zip && \
    unzip kotlin-compiler-1.6.10.zip -d /opt && \
    ln -s /opt/kotlinc/bin/kotlinc /usr/bin/kotlinc
RUN wget https://services.gradle.org/distributions/gradle-7.4.2-bin.zip && \
    unzip gradle-7.4.2-bin.zip -d /opt/gradle && \
    ln -s /opt/gradle/gradle-7.4.2/bin/gradle /usr/bin/gradle
# Sprawdź wersję Pythona, Java, Kotlin
CMD python3.8 --version && java -version && kotlinc -version
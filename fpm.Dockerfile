FROM debian:8
RUN apt-get update && apt-get install -y ruby ruby-dev build-essential
RUN gem install fpm
WORKDIR /packaging

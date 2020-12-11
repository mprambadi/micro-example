FROM php:7.4-cli

WORKDIR /app

COPY . /app

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN composer install -vvv

EXPOSE 4002

CMD [ "php","-S","0.0.0.0:4002","-t","public" ]

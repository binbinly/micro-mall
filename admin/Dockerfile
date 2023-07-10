FROM php:8.1-cli

WORKDIR /usr/src/mall

COPY . .

# see: https://github.com/chialab/docker-php
# Download script to install PHP extensions and dependencies
ADD https://raw.githubusercontent.com/mlocati/docker-php-extension-installer/master/install-php-extensions /usr/local/bin/

# 安装扩展
RUN chmod uga+x /usr/local/bin/install-php-extensions && sync

RUN DEBIAN_FRONTEND=noninteractive apt-get update -q \
    && DEBIAN_FRONTEND=noninteractive apt-get install -qq -y \
      curl git zip unzip \
    && install-php-extensions \
      bcmath \
      bz2 \
      calendar \
      exif \
      gd \
      intl \
      ldap \
      mysqli \
      opcache \
      pdo_mysql \
      pdo_pgsql \
      pgsql \
      redis \
      soap \
      xsl \
      zip \
      sockets

# Install Composer
RUN php -r "copy('https://install.phpcomposer.com/installer', 'composer-setup.php');" \
  && php composer-setup.php && php -r "unlink('composer-setup.php');" && mv composer.phar /usr/local/bin/composer

RUN cp .env.docker .env && composer install --no-dev --optimize-autoloader --prefer-dist

EXPOSE 8000

ENTRYPOINT [ "php", "artisan", "serve"]
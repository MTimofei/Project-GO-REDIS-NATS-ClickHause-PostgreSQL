FROM ubuntu:latest
# Базовый образ - Ubuntu latest

# Установка пакетов для компиляции плагина
RUN apt-get update && \
    apt-get install -y git gcc make postgresql-server-dev-* cmake clickhouse-client
# Обновление репозиториев и установка необходимых пакетов: git, gcc, make, postgresql-server-dev-* (необходимо для сборки плагина postgres_fdw), cmake, clickhouse-client

# Клонирование репозитория clickhouse-dbi
RUN git clone https://github.com/ClickHouse/clickhouse-dbi.git && \
    cd clickhouse-dbi && \
    git checkout v0.1.0
# Клонирование репозитория clickhouse-dbi с github.com и переключение на тег v0.1.0

# Сборка плагина postgres_fdw для Clickhouse
RUN cd clickhouse-dbi && \
    mkdir build && \
    cd build && \
    cmake .. && \
    make -j4 && \
    cp src/Server/dbms/programs/clickhouse-odbc-bridge/postgres_fdw.so /usr/lib/clickhouse/udfs/
# Сборка плагина postgres_fdw для Clickhouse. Создание директории build в репозитории clickhouse-dbi, переход в нее, выполнение cmake, компиляция плагина с помощью make, и копирование скомпилированного плагина в директорию /usr/lib/clickhouse/udfs/

# Установка Postgres
RUN apt-get install -y postgresql
ENV POSTGRES_USER myuser
ENV POSTGRES_PASSWORD mypassword
# Установка СУБД PostgreSQL и установка переменных среды для имени пользователя и пароля

# Установка Clickhouse
RUN apt-get install -y clickhouse-server
ENV CLICKHOUSE_PASSWORD mypassword
# Установка СУБД ClickHouse и установка переменной среды для пароля

# Установка Nats
RUN apt-get install -y gnatsd
# Установка сервера сообщений NATS

# Установка Redis
RUN apt-get install -y redis-server
ENV REDIS_PASSWORD mypassword
# Установка СУБД Redis и установка переменной среды для пароля

EXPOSE 5432
EXPOSE 8123
EXPOSE 9000
EXPOSE 4222
EXPOSE 6222
EXPOSE 8222
EXPOSE 6379
# Открытие портов для доступа к СУБД PostgreSQL (5432), СУБД ClickHouse (8123 и 9000), серверу сообщений NATS (4222, 6222, 8222) и СУБД Redis (6379)

CMD ["/bin/bash"] 

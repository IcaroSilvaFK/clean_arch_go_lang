FROM mysql

COPY initdb.sql /docker-entrypoint-initdb.d/initdb.sql
FROM postgres:9.6
ENV POSTGRES_DB=opensirenedb
ENV POSTGRES_USER=sir
ENV POSTGRES_PASSWORD=sirensiren
ADD sql/schema.sql /docker-entrypoint-initdb.d/schema.sql

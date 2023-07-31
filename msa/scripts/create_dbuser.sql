create user profi identified by profi


CREATE ROLE profi LOGIN password 'profi1' SUPERUSER;

CREATE ROLE user_name PASSWORD 'tYPe_YoUr_PaSSwOrD' NOSUPERUSER CREATEDB CREATEROLE INHERIT LOGIN;

create table all_users

sudo apt-get update && sudo apt-get install pgadmin4 pgadmin4-apache2 -y


CREATE DATABASE logins
    WITH
    OWNER = profi
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;


--select * from users

--users_id_seq

CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL PRIMARY KEY,
    user_name varchar(255),
    first_name varchar(255),
    last_name varchar(255),
    email varchar(255),
    phone varchar(255)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to profi;
    


CREATE TABLE public.all_users
(
    id serial,
    user_name character varying(255)[],
    first_name character varying(255)[],
    last_name character varying(255)[],
    email character varying(255),
    phone character varying(255)[],
    CONSTRAINT pk_all_users PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.all_users
    OWNER to profi;
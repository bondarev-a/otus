--create user profi
CREATE ROLE profi LOGIN password 'profi1' SUPERUSER;

CREATE DATABASE logins
    WITH
    OWNER = profi
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

GRANT ALL PRIVILEGES ON DATABASE logins TO profi;


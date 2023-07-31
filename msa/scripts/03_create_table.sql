
CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL PRIMARY KEY,
    user_name varchar(255),
    first_name varchar(255),
    last_name varchar(255),
    email varchar(255),
    phone varchar(255)
) TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to logins;


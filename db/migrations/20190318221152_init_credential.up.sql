-- Table: public.credentials
CREATE TABLE public.credentials
(
    uuid uuid NOT NULL,
    username character(180) COLLATE pg_catalog."default" NOT NULL,
    email character(180) COLLATE pg_catalog."default" NOT NULL,
    password character(180) COLLATE pg_catalog."default" NOT NULL,
    active boolean NOT NULL DEFAULT false,
    created timestamp without time zone NOT NULL DEFAULT now(),
    deleted timestamp without time zone,
    CONSTRAINT credentials_pkey PRIMARY KEY (uuid),
    CONSTRAINT credential_username_unique UNIQUE (username)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

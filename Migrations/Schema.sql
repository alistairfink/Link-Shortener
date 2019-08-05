-- Table: public.link

-- DROP TABLE public.link;

CREATE TABLE public.link
(
    id text COLLATE pg_catalog."default" NOT NULL,
    link text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT link_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.link
    OWNER to postgres;
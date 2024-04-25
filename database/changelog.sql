--liquibase formatted sql

--changeset pidanou.eang:1 labels:example-label context:example-context
--comment: init
CREATE TABLE client (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    address TEXT NOT NULL,
    postal_code TEXT NOT NULL,
    city TEXT NOT NULL,
    country TEXT NOT NULL,
    phone TEXT NOT NULL,
    email TEXT NOT NULL,
    date_created DATE NOT NULL default now(),
    PRIMARY KEY (id)
);



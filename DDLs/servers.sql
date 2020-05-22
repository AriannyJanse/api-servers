-- public.servers definition

-- Drop table

-- DROP TABLE public.servers;

CREATE TABLE servers (
	id_server INT8 NOT NULL DEFAULT unique_rowid(),
	hostname VARCHAR NOT NULL,
	logo VARCHAR NOT NULL,
	title VARCHAR NOT NULL,
	is_down BOOL NOT NULL,
	createdat TIMESTAMP NULL DEFAULT now():::TIMESTAMP,
	updatedat TIMESTAMP NULL,
	CONSTRAINT servers_pkey PRIMARY KEY (id_server ASC),
	UNIQUE INDEX servers_hostname_key (hostname ASC),
	FAMILY "primary" (id_server, hostname, logo, title, is_down, createdat, updatedat)
);
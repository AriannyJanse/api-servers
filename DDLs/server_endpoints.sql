-- public.server_endpoints definition

-- Drop table

-- DROP TABLE public.server_endpoints;

CREATE TABLE server_endpoints (
	id_server_endpoint INT8 NOT NULL DEFAULT unique_rowid(),
	endpoint_id_server INT8 NOT NULL,
	endpoint_address VARCHAR NOT NULL,
	endpoint_ssl_grade VARCHAR NULL DEFAULT '':::STRING,
	endpoint_country VARCHAR NOT NULL,
	endpoint_owner VARCHAR NOT NULL,
	createdat TIMESTAMP NULL DEFAULT now():::TIMESTAMP,
	updatedat TIMESTAMP NULL,
	CONSTRAINT server_endpoints_pkey PRIMARY KEY (id_server_endpoint ASC),
	CONSTRAINT server_endpoints_id_server_fkey FOREIGN KEY (endpoint_id_server) REFERENCES servers(id_server) ON DELETE CASCADE,
	INDEX server_endpoints_auto_index_server_endpoints_id_server_fkey (endpoint_id_server ASC),
	FAMILY "primary" (id_server_endpoint, endpoint_id_server, endpoint_address, endpoint_ssl_grade, endpoint_country, endpoint_owner, createdat, updatedat)
);


-- public.server_endpoints foreign keys

ALTER TABLE public.server_endpoints ADD CONSTRAINT server_endpoints_id_server_fkey FOREIGN KEY (endpoint_id_server) REFERENCES servers(id_server) ON DELETE CASCADE;
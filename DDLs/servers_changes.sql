-- public.servers_changes definition

-- Drop table

-- DROP TABLE public.servers_changes;

CREATE TABLE servers_changes (
	id_server_changes INT8 NOT NULL DEFAULT unique_rowid(),
	changes_id_server INT8 NOT NULL,
	ssl_grade VARCHAR NULL DEFAULT '':::STRING,
	previous_ssl_grade VARCHAR NULL DEFAULT '':::STRING,
	createdat TIMESTAMP NULL DEFAULT now():::TIMESTAMP,
	updatedat TIMESTAMP NULL,
	CONSTRAINT servers_changes_pkey PRIMARY KEY (id_server_changes ASC),
	CONSTRAINT servers_changes_id_server_fkey FOREIGN KEY (changes_id_server) REFERENCES servers(id_server) ON DELETE CASCADE,
	INDEX servers_changes_auto_index_servers_changes_id_server_fkey (changes_id_server ASC),
	FAMILY "primary" (id_server_changes, changes_id_server, ssl_grade, previous_ssl_grade, createdat, updatedat)
);


-- public.servers_changes foreign keys

ALTER TABLE public.servers_changes ADD CONSTRAINT servers_changes_id_server_fkey FOREIGN KEY (changes_id_server) REFERENCES servers(id_server) ON DELETE CASCADE;
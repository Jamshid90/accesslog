CREATE TABLE IF NOT EXISTS access_log (
    "id" character varying(50) NOT NULL,
    "user_id" character varying(50) DEFAULT '',
    "method" character varying(500) DEFAULT '',
    "url" character varying(500) DEFAULT '',
    "data" jsonb,
    "created_at" timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT access_log_id_pkey PRIMARY KEY (id));
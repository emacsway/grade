CREATE TABLE event_log (
    tenant_id uuid NOT NULL,
    stream_type varchar(128) NOT NULL,
    stream_id varchar(255) NOT NULL,
    stream_position int NOT NULL,
    event_type varchar(60) NOT NULL,
    event_version int8 NOT NULL,
    payload jsonb NOT NULL,
    metadata jsonb NULL,
    CONSTRAINT event_log_event_id_uniq UNIQUE (metadata->>'event_id'),
    CONSTRAINT event_log_pk PRIMARY KEY (tenant_id, stream_type, stream_id, stream_position)
);

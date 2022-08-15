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

CREATE TABLE recognizer (
    tenant_id uuid NOT NULL,
    member_id uuid NOT NULL,
    grade int8 NOT NULL DEFAULT 0,
    available_endorsement_count NOT NULL,
    pending_endorsement_count NOT NULL DEFAULT 0,
    created_at timestamp with time zone NOT NULL,
    version int NOT NULL,
    CONSTRAINT recognizer_pk PRIMARY KEY (tenant_id, member_id)
);

CREATE TABLE endorsement (
    tenant_id uuid NOT NULL,
    specialist_id uuid NOT NULL,
    specialist_grade int8 NOT NULL DEFAULT 0,
    specialist_version int NOT NULL,
    recognizer_id uuid NOT NULL,
    recognizer_grade int8 NOT NULL DEFAULT 0,
    recognizer_version int NOT NULL,
    artifact_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    CONSTRAINT endorsement_uniq UNIQUE (tenant_id, specialist_id, artifact_id, recognizer_id),
    CONSTRAINT endorsement_uniq UNIQUE (tenant_id, recognizer_id, recognizer_version),
    CONSTRAINT endorsement_pk PRIMARY KEY (tenant_id, specialist_id, specialist_version)
);
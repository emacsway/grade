CREATE TABLE tenant (
        id bigserial CONSTRAINT tenant_pk PRIMARY KEY,
        name varchar(150) NOT NULL,
        created_at timestamp with time zone NOT NULL,
        version integer NOT NULL
);


CREATE TABLE member (
    tenant_id integer NOT NULL REFERENCES tenant(id) ON DELETE CASCADE,
    member_id bigint NOT NULL,
    status smallint NOT NULL,
    first_name varchar(150) NOT NULL,
    last_name varchar(150) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    version integer NOT NULL,
    CONSTRAINT member_pk PRIMARY KEY (tenant_id, member_id)
);


CREATE FUNCTION make_member_seq() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
    execute format('CREATE SEQUENCE IF NOT EXISTS member_seq_%s', NEW.id);
    return NEW;
end
$$;
CREATE TRIGGER make_member_seq AFTER INSERT ON tenant FOR EACH ROW EXECUTE PROCEDURE make_member_seq();


CREATE FUNCTION drop_member_seq() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
execute format('DROP SEQUENCE IF EXISTS member_seq_%s', OLD.id);
return NEW;
end
$$;
CREATE TRIGGER drop_member_seq AFTER DELETE ON tenant FOR EACH ROW EXECUTE PROCEDURE drop_member_seq();


CREATE FUNCTION fill_in_member_seq() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
    NEW.member_id := nextval('member_seq_' || NEW.tenant_id);
    RETURN NEW;
end
$$;
CREATE TRIGGER fill_in_member_seq BEFORE INSERT ON member FOR EACH ROW EXECUTE PROCEDURE fill_in_member_seq();

-- tested up to here

CREATE TABLE event_log (
    tenant_id integer NOT NULL REFERENCES tenant(id) ON DELETE CASCADE,
    stream_type varchar(128) NOT NULL,
    stream_id varchar(255) NOT NULL,
    stream_position integer NOT NULL,
    event_type varchar(60) NOT NULL,
    event_version smallint NOT NULL,
    payload jsonb NOT NULL,
    metadata jsonb NULL,
    CONSTRAINT event_log_event_id_uniq UNIQUE (metadata->>'event_id'),
    CONSTRAINT event_log_pk PRIMARY KEY (tenant_id, stream_type, stream_id, stream_position)
);


CREATE TABLE endorser (
    tenant_id integer NOT NULL REFERENCES tenant(id) ON DELETE CASCADE,
    member_id bigint NOT NULL REFERENCES member(id) ON DELETE CASCADE,
    grade smallint NOT NULL DEFAULT 0,
    available_endorsement_count NOT NULL,
    pending_endorsement_count NOT NULL DEFAULT 0,
    created_at timestamp with time zone NOT NULL,
    version integer NOT NULL,
    CONSTRAINT endorser_pk PRIMARY KEY (tenant_id, member_id)
);


CREATE TABLE endorsement (
    tenant_id integer NOT NULL REFERENCES tenant(id) ON DELETE CASCADE,
    specialist_id bigint NOT NULL REFERENCES member(id) ON DELETE CASCADE,
    specialist_grade smallint NOT NULL DEFAULT 0,
    specialist_version integer NOT NULL,
    artifact_id bigint NOT NULL,
    endorser_id bigint NOT NULL REFERENCES member(id) ON DELETE CASCADE,
    endorser_grade smallint NOT NULL DEFAULT 0,
    endorser_version integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    CONSTRAINT endorsement_uniq UNIQUE (tenant_id, specialist_id, artifact_id, endorser_id),
    CONSTRAINT endorsement_uniq UNIQUE (tenant_id, endorser_id, endorser_version),
    CONSTRAINT endorsement_pk PRIMARY KEY (tenant_id, specialist_id, specialist_version)
);

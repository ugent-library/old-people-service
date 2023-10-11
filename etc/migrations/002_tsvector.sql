CREATE EXTENSION IF NOT EXISTS unaccent;
DO
$$BEGIN
	CREATE TEXT SEARCH CONFIGURATION usimple ( COPY = simple );
EXCEPTION
	WHEN unique_violation THEN
	   NULL;
END;$$;

ALTER TEXT SEARCH CONFIGURATION usimple ALTER MAPPING FOR hword, hword_part, word WITH unaccent, simple;

ALTER TABLE organization
ADD COLUMN IF NOT EXISTS ts tsvector GENERATED ALWAYS AS
(
	to_tsvector('simple', jsonb_path_query_array(identifier,'$.**{2}')) ||
	to_tsvector('simple', public_id) ||
	to_tsvector('usimple',name_dut) ||
	to_tsvector('usimple', name_eng)
) STORED;
CREATE INDEX IF NOT EXISTS organization_ts_idx ON organization USING GIN(ts);
ALTER TABLE person
	ADD COLUMN IF NOT EXISTS ts tsvector GENERATED ALWAYS AS (
		to_tsvector('usimple',name)
	) STORED;
CREATE INDEX IF NOT EXISTS person_ts_idx ON person USING GIN(ts);

---- create above / drop below ----

DROP INDEX IF EXISTS person_ts_idx;
ALTER TABLE person DROP COLUMN IF EXISTS ts;
DROP INDEX IF EXISTS organization_ts_idx;
ALTER TABLE organization DROP COLUMN IF EXISTS ts;

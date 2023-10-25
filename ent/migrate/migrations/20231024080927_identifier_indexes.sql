CREATE INDEX IF NOT EXISTS organization_identifier_gismo_id_gin ON organization USING GIN((identifier->'gismo_id') jsonb_ops);
CREATE INDEX IF NOT EXISTS organization_identifier_ugent_id_gin ON organization USING GIN((identifier->'ugent_id') jsonb_ops);
CREATE INDEX IF NOT EXISTS organization_identifier_biblio_id_gin ON organization USING GIN((identifier->'biblio_id') jsonb_ops);
CREATE INDEX IF NOT EXISTS organization_identifier_ugent_memorialis_id_gin ON organization USING GIN((identifier->'ugent_memorialis_id') jsonb_ops);
CREATE INDEX IF NOT EXISTS person_identifier_ugent_username_gin ON person USING GIN((identifier->'ugent_username') jsonb_ops);
CREATE INDEX IF NOT EXISTS person_identifier_ugent_id_gin ON person USING GIN((identifier->'ugent_id') jsonb_ops);
CREATE INDEX IF NOT EXISTS person_identifier_historic_ugent_id_gin ON person USING GIN((identifier->'historic_ugent_id') jsonb_ops);
CREATE INDEX IF NOT EXISTS person_identifier_old_biblio_id_gin ON person USING GIN((identifier->'old_biblio_id') jsonb_ops);
ALTER TABLE "organization_parent"
    ADD CONSTRAINT "organization_parent_ref_organization_id"
    FOREIGN KEY ("organization_id") REFERENCES "organization" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE "organization_parent"
    ADD CONSTRAINT "organization_parent_ref_parent_organization_id"
    FOREIGN KEY ("parent_organization_id") REFERENCES "organization" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;

ALTER TABLE "organization_person"
    ADD CONSTRAINT "organization_person_ref_organization_id"
    FOREIGN KEY ("organization_id") REFERENCES "organization" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE "organization_person"
    ADD CONSTRAINT "organization_person_ref_person_id"
    FOREIGN KEY ("person_id") REFERENCES "person" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;

---- create above / drop below ----

ALTER TABLE "organization_parent"
    DROP CONSTRAINT IF EXISTS "organization_parent_ref_organization_id";
ALTER TABLE "organization_parent"
    DROP CONSTRAINT IF EXISTS "organization_parent_ref_parent_organization_id";

ALTER TABLE "organization_person"
    DROP CONSTRAINT IF EXISTS "organization_person_ref_organization_id";
ALTER TABLE "organization_person"
    DROP CONSTRAINT IF EXISTS "organization_person_ref_person_id";

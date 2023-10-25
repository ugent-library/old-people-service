ALTER TABLE "organization_parent" 
    ADD CONSTRAINT "organization_parent_ref_organization_id" 
    FOREIGN KEY ("organization_id") REFERENCES "organization" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE "organization_parent"
    ADD CONSTRAINT "organization_parent_ref_parent_organization_id"
    FOREIGN KEY ("parent_organization_id") REFERENCES "organization" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
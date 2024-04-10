-- UserProjectRelations
ALTER TABLE "user_project_relations"
    ADD CONSTRAINT "FK_UserProjectRelations_ProjectID" FOREIGN KEY ("project_id") REFERENCES "projects" ("project_id");
ALTER TABLE "user_project_relations"
    ADD CONSTRAINT "FK_UserProjectRelations_UserID" FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
-- ProjectLinkRelations
ALTER TABLE "project_link_relations"
    ADD CONSTRAINT "FK_ProjectLinkRelations_ProjectID" FOREIGN KEY ("project_id") REFERENCES "projects" ("project_id");
ALTER TABLE "project_link_relations"
    ADD CONSTRAINT "FK_ProjectLinkRelations_LinkID" FOREIGN KEY ("link_id") REFERENCES "links" ("link_id");
-- UserLinkRelations
ALTER TABLE "user_link_relations"
    ADD CONSTRAINT "FK_UserLinkRelations_UserID" FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
ALTER TABLE "user_link_relations"
    ADD CONSTRAINT "FK_UserLinkRelations_LinkID" FOREIGN KEY ("link_id") REFERENCES "links" ("link_id");
-- LinkRoutes
ALTER TABLE "link_routes"
    ADD CONSTRAINT "FK_LinkRoutes_LinkID" FOREIGN KEY ("link_id") REFERENCES "links" ("link_id");
ALTER TABLE "link_routes"
    ADD CONSTRAINT "FK_LinkRoutes_PlatformID" FOREIGN KEY ("platform_id") REFERENCES "platforms" ("platform_id");
ALTER TABLE "link_routes"
    ADD CONSTRAINT "FK_LinkRoutes_AvailabilityZoneID" FOREIGN KEY ("a_zone_id") REFERENCES "availability_zones" ("a_zone_id");
-- ClicksStats
ALTER TABLE "clicks_stats"
    ADD CONSTRAINT "FK_ClicksStats_RouteID" FOREIGN KEY ("route_id") REFERENCES "link_routes" ("route_id");

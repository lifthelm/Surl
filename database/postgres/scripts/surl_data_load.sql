-- Users
COPY "users" ("user_id", "username", "email", "password", "registration_date", "user_role")
FROM '/scripts/data/users.csv' DELIMITER ',' CSV HEADER;

-- Projects
COPY "projects" ("project_id", "name", "description")
FROM '/scripts/data/projects.csv' DELIMITER ',' CSV HEADER;

-- Links
COPY "links" ("link_id", "def_long_url", "short_url", "create_date")
FROM '/scripts/data/links.csv' DELIMITER ',' CSV HEADER;

-- Platforms
COPY "platforms" ("platform_id", "description")
FROM '/scripts/data/platforms.csv' DELIMITER ',' CSV HEADER;

-- AvailabilityZones
COPY "availability_zones" ("a_zone_id", "description")
FROM '/scripts/data/availability_zones.csv' DELIMITER ',' CSV HEADER;

-- LinkRoutes
COPY "link_routes" ("route_id", "link_id", "platform_id", "a_zone_id", "long_url")
FROM '/scripts/data/link_routes.csv' DELIMITER ',' CSV HEADER;

-- ClicksStats
COPY "clicks_stats" ("record_id", "route_id", "click_date_time", "user_ip_address", "user_agent", "user_country")
FROM '/scripts/data/click_stats.csv' DELIMITER ',' CSV HEADER;

-- UserProjectRelation
COPY "user_project_relations" ("relation_id", "user_id", "project_id", "user_role")
FROM '/scripts/data/user_project_relation.csv' DELIMITER ',' CSV HEADER;

-- ProjectLinkRelation
COPY "project_link_relations" ("relation_id", "project_id", "link_id")
FROM '/scripts/data/project_link_ownership.csv' DELIMITER ',' CSV HEADER;

-- UserLinkRelation
COPY "user_link_relations" ("relation_id", "user_id", "link_id", "user_role")
FROM '/scripts/data/user_link_relation.csv' DELIMITER ',' CSV HEADER;

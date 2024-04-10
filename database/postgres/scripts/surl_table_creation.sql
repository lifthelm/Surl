-- surl
CREATE TABLE "users" (
    "user_id" serial PRIMARY KEY,
    "username" VARCHAR(64) NOT NULL UNIQUE,
    "email" VARCHAR(128) NOT NULL,
    "password" VARCHAR(64) NOT NULL,
    "registration_date" TIMESTAMP NOT NULL,
    "user_role" VARCHAR(16) NOT NULL
);

CREATE TABLE "projects" (
    "project_id" serial PRIMARY KEY,
    "name" VARCHAR(128) NOT NULL,
    "description" TEXT
);

CREATE TABLE "links" (
    "link_id" serial PRIMARY KEY,
    "def_long_url" VARCHAR(256) NOT NULL,
    "short_url" VARCHAR(32) NOT NULL UNIQUE,
    "create_date" TIMESTAMP NOT NULL
);

CREATE TABLE "user_project_relations" (
    "relation_id" serial PRIMARY KEY,
    "user_id" INT REFERENCES "users" ("user_id") NOT NULL,
    "project_id" INT REFERENCES "projects" ("project_id") NOT NULL,
    "user_role" VARCHAR(16) NOT NULL
);

CREATE TABLE "project_link_relations" (
    "relation_id" serial PRIMARY KEY,
    "project_id" INT REFERENCES "projects" ("project_id") NOT NULL,
    "link_id" INT REFERENCES "links" ("link_id") NOT NULL
);

CREATE TABLE "user_link_relations" (
    "relation_id" serial PRIMARY KEY,
    "user_id" INT REFERENCES "users" ("user_id") NOT NULL,
    "link_id" INT REFERENCES "links" ("link_id") NOT NULL,
    "user_role" VARCHAR(16) NOT NULL
);

CREATE TABLE "platforms" (
    "platform_id" serial PRIMARY KEY,
    "description" TEXT
);

CREATE TABLE "availability_zones" (
    "a_zone_id" serial PRIMARY KEY,
    "description" TEXT
);

CREATE TABLE "link_routes" (
    "route_id" serial PRIMARY KEY,
    "link_id" INT REFERENCES "links" ("link_id") NOT NULL,
    "platform_id" INT REFERENCES "platforms" ("platform_id") NOT NULL,
    "a_zone_id" INT REFERENCES "availability_zones" ("a_zone_id") NOT NULL,
    "long_url" VARCHAR(256) NOT NULL
);

CREATE TABLE "clicks_stats" (
    "record_id" serial PRIMARY KEY,
    "route_id" INT REFERENCES "link_routes" ("route_id") NOT NULL,
    "click_date_time" TIMESTAMP NOT NULL,
    "user_ip_address" VARCHAR(16) NOT NULL,
    "user_agent" VARCHAR(512) NOT NULL,
    "user_country" VARCHAR(64) NOT NULL
);
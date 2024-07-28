-- Create "deneme" table
CREATE TABLE "deneme" ("id" uuid NOT NULL, "count" integer NOT NULL, "test_id" uuid NULL, "is_active" boolean NOT NULL, "deneme_type" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "test" table
CREATE TABLE "test" ("id" uuid NOT NULL, "name" character varying NOT NULL, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create "account" table
CREATE TABLE "account" ("id" uuid NOT NULL, "name" character varying NOT NULL, "surname" character varying NOT NULL, "deneme_id" uuid NULL, PRIMARY KEY ("id"));
-- Create "group" table
CREATE TABLE "group" ("id" uuid NOT NULL, "name" character varying NOT NULL, "surname" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "account_group" table
CREATE TABLE "account_group" ("id" smallserial NOT NULL, "account_id" uuid NOT NULL, "group_id" uuid NOT NULL, PRIMARY KEY ("id"));

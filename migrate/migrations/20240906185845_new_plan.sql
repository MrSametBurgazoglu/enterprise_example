-- Create "deneme" table
CREATE TABLE "deneme" ("id" uuid NOT NULL, "id" uuid NOT NULL, "test_id" uuid NULL, "count" integer NOT NULL, "is_active" boolean NOT NULL, "deneme_type" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "test" table
CREATE TABLE "test" ("id" uuid NOT NULL, "id" uuid NOT NULL, "name" character varying NOT NULL, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create "account" table
CREATE TABLE "account" ("id" uuid NOT NULL, "id" uuid NOT NULL, "name" character varying NOT NULL, "surname" character varying NOT NULL, "deneme_id" uuid NULL, PRIMARY KEY ("id"));
-- Create "group" table
CREATE TABLE "group" ("id" uuid NOT NULL, "id" uuid NOT NULL, "name" character varying NOT NULL, "surname" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "account_group" table
CREATE TABLE "account_group" ("id" smallserial NOT NULL, "account_id" uuid NOT NULL, "group_id" uuid NOT NULL, PRIMARY KEY ("id"));
-- Modify "deneme" table
ALTER TABLE "deneme" ADD CONSTRAINT "deneme_test_id" FOREIGN KEY ("test_id") REFERENCES "test" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "account" table
ALTER TABLE "account" ADD CONSTRAINT "account_deneme_id" FOREIGN KEY ("deneme_id") REFERENCES "deneme" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "account_group_id" FOREIGN KEY ("group_id") REFERENCES "group" ("account_id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "group" table
ALTER TABLE "group" ADD CONSTRAINT "group_account_id" FOREIGN KEY ("account_id") REFERENCES "account" ("group_id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "account_group" table
ALTER TABLE "account_group" ADD CONSTRAINT "account_group_account" FOREIGN KEY ("account_id") REFERENCES "account" ("account_id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "account_group_group" FOREIGN KEY ("group_id") REFERENCES "group" ("group_id") ON UPDATE NO ACTION ON DELETE NO ACTION;

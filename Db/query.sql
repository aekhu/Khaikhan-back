create database khairkhan;
create user admin with encrypted password 'password';
grant all privileges on database khairkhan to admin;

--change database and change user as admin: \c guush admin
CREATE TABLE "Users" (
  "id" SERIAL PRIMARY KEY,
  "role" varchar(20),
  "email" varchar(30),
  "passsword" varchar(30),
  "firstname" varchar(30),
  "lastname" varchar(30),
  "birth_date" date,
  "registered_date" date,
  "phone" char(8),
  "event_id" varchar(10)
);

CREATE TABLE "Event" (
  "id" varchar(20) PRIMARY KEY,
  "title" text,
  "description" text,
  "trail_id" varchar(120),
  "photo_id" varchar(20),
  "category" varchar(20),
  "begin_date" datetime,
  "owner_id" varchar(10),
  "user_id" varchar(10)
);

CREATE TABLE "Trail" (
  "id" varchar(20) PRIMARY KEY,
  "title" text,
  "description" text,
  "photo_id" varchar(10),
  "category" varchar(20),
  "length" varchar(20)
);

CREATE TABLE "Tips" (
  "id" varchar(20) PRIMARY KEY,
  "title" text,
  "description" text,
  "type" text,
  "user_id" varchar(20),
  "publish_date" datetime
);

CREATE TABLE "Photo" (
  "id" varchar(20) PRIMARY KEY,
  "owner_id" varchar(10),
  "url" text
);

CREATE TABLE "Category" (
  "id" varchar(20) PRIMARY KEY,
  "name" text
);

CREATE TABLE "Event_user" (
  "id" SERIAL PRIMARY KEY,
  "user_id" varchar(20),
  "event_id" varchar(20)
);

CREATE TABLE "Tips_user" (
  "id" SERIAL PRIMARY KEY,
  "user_id" varchar(20),
  "tip_id" varchar(20)
);

CREATE TABLE "Trail_user" (
  "id" SERIAL PRIMARY KEY,
  "user_id" varchar(20),
  "trail_id" varchar(20)
);

ALTER TABLE "Photo" ADD FOREIGN KEY ("owner_id") REFERENCES "Users" ("id");

ALTER TABLE "Users" ADD FOREIGN KEY ("id") REFERENCES "Event_user" ("user_id");

ALTER TABLE "Event" ADD FOREIGN KEY ("id") REFERENCES "Event_user" ("event_id");

ALTER TABLE "Users" ADD FOREIGN KEY ("id") REFERENCES "Tips_user" ("user_id");

ALTER TABLE "Tips" ADD FOREIGN KEY ("id") REFERENCES "Tips_user" ("tip_id");

ALTER TABLE "Users" ADD FOREIGN KEY ("id") REFERENCES "Trail_user" ("user_id");

ALTER TABLE "Trail" ADD FOREIGN KEY ("id") REFERENCES "Trail_user" ("trail_id");

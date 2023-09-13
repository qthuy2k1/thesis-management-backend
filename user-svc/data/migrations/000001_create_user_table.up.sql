CREATE TABLE "users" (
  "id" serial PRIMARY KEY not null,
  "class" varchar NOT NULL,
  "major" varchar,
  "phone" varchar,
  "photo_src" varchar,
  "role" varchar not null,
  "name" varchar not null,
  "email" varchar not null,
  "classroom_id" integer not null
);
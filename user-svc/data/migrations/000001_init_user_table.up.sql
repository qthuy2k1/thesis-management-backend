CREATE TABLE "users" (
  "id" VARCHAR PRIMARY KEY not null,
  "class" varchar NOT NULL,
  "major" varchar,
  "phone" varchar,
  "photo_src" varchar,
  "role" varchar not null,
  "name" varchar not null,
  "email" varchar not null,
  "hashed_password" varchar
);

CREATE TABLE "members" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "classroom_id" INTEGER NOT NULL,
    "member_id" VARCHAR NOT NULL,
    "status" VARCHAR,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE,
    "is_defense" BOOLEAN
);


CREATE TABLE IF NOT EXISTS "student_defs" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "user_id" VARCHAR NOT NULL,
    "instructor_id" VARCHAR NOT NULL,
    "time_slots_id" INTEGER DEFAULT 0
);

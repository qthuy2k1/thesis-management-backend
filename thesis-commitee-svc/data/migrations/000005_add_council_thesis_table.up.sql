CREATE TABLE IF NOT EXISTS "councils" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "lecturer_id" VARCHAR NOT NULL,
    "thesis_id" INTEGER
);

CREATE TABLE IF NOT EXISTS "schedule" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "room_id" INTEGER,
    "thesis_id" INTEGER
);


CREATE TABLE IF NOT EXISTS "time_slots" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "schedule_id" INTEGER
);

CREATE TABLE IF NOT EXISTS "thesis" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE "thesis_commitees"
ADD COLUMN "time_slots_id" INTEGER;

ALTER TABLE "thesis_commitees"
ADD COLUMN "time" VARCHAR;
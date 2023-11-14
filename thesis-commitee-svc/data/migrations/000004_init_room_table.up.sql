CREATE TABLE IF NOT EXISTS "rooms" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "name" VARCHAR NOT NULL,
    "type" VARCHAR NOT NULL,
    "school" VARCHAR NOT NULL,
    "description" VARCHAR
);

ALTER TABLE "thesis_commitees"
DROP COLUMN "room";

ALTER TABLE "thesis_commitees"
ADD COLUMN "room_id" INTEGER NOT NULL;

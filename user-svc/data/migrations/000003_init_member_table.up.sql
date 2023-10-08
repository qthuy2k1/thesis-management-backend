CREATE TABLE "members" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "classroom_id" INTEGER NOT NULL,
    "member_id" VARCHAR NOT NULL,
    "status" VARCHAR
);

ALTER TABLE "users"
DROP COLUMN "classroom_id";
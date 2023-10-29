DROP TABLE IF EXISTS "rooms";


ALTER TABLE "thesis_commitees"
ADD COLUMN "room" VARCHAR NOT NULL;

ALTER TABLE "thesis_commitees"
DROP COLUMN "room_id" INTEGER NOT NULL;

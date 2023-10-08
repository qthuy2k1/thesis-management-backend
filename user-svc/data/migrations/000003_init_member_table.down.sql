DROP TABLE IF EXISTS "members";

ALTER TABLE "users"
ADD COLUMN "classroom_id" INTEGER;
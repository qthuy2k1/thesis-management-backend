ALTER TABLE "submissions"
ALTER COLUMN "attachment_id" TYPE INTEGER USING (attachment_id::integer);
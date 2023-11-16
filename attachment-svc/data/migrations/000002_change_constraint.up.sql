ALTER TABLE "attachments" ALTER COLUMN submission_id DROP NOT NULL;
ALTER TABLE "attachments" ALTER COLUMN exercise_id DROP NOT NULL;

ALTER TABLE "attachments" ADD COLUMN post_id INTEGER;
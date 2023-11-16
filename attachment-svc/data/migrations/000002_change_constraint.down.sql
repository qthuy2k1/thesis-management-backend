ALTER TABLE "attachments" ALTER COLUMN submission_id SET NOT NULL;
ALTER TABLE "attachments" ALTER COLUMN exercise_id SET NOT NULL;

ALTER TABLE "attachments" DROP COLUMN post_id;
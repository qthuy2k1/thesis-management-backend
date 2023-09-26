CREATE TABLE "attachments" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "file_url" TEXT NOT NULL,
    "status" VARCHAR NOT NULL,
    "submission_id" INTEGER NOT NULL,
    "exercise_id" INTEGER NOT NULL,
    "author_id" VARCHAR NOT NULL,
    "created_at" timestamp not null default current_timestamp(0)::timestamp without time zone
);
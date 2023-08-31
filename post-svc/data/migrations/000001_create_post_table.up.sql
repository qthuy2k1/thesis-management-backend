CREATE TABLE "posts" (
  "id" serial PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "content" text NOT NULL,
  "classroom_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

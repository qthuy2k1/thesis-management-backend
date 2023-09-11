CREATE TABLE "submissions" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" integer NOT NULL,
  "exercise_id" integer NOT NULL,
  "submission_date" timestamp NOT NULL,
  "status" varchar NOT NULL
);

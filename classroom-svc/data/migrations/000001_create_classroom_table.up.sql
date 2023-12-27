CREATE TABLE "classrooms" (
  "id" serial primary key not null,
  "status" varchar not null,
  "created_at" timestamp not null default current_timestamp(0)::timestamp without time zone,
  "updated_at" timestamp not null default current_timestamp(0)::timestamp without time zone,
  "lecturer_id" VARCHAR NOT NULL,
  "class_course" VARCHAR,
  "quantity_student" INTEGER,
  "topic_tags" VARCHAR
);

CREATE TABLE "waiting_lists" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "classroom_id" INTEGER NOT NULL,
    "user_id" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::timestamp without time zone,
    "is_defense" BOOLEAN,
    "status" VARCHAR
);

CREATE INDEX waiting_lists_classroom_id ON waiting_lists(classroom_id);
CREATE INDEX waiting_lists_user_id ON waiting_lists(user_id);



CREATE TABLE "exercises" (
  "id" serial PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "content" text NOT NULL,
  "classroom_id" integer NOT NULL,
  "deadline" timestamp NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE,
  "reporting_stage_id" INTEGER,
  "author_id" VARCHAR
);

CREATE TABLE "posts" (
  "id" serial PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "content" text NOT NULL,
  "classroom_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE,
  "reporting_stage_id" INTEGER,
  "author_id" VARCHAR
);

CREATE TABLE "reporting_stages" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "label" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "value" VARCHAR
);

CREATE TABLE "submissions" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" VARCHAR NOT NULL,
  "exercise_id" INTEGER NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE
);

CREATE INDEX submissions_execise_id ON submissions(exercise_id);


CREATE TABLE "attachments" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "file_url" TEXT NOT NULL,
    "status" VARCHAR NOT NULL,
    "submission_id" INTEGER,
    "exercise_id" INTEGER,
    "post_id" INTEGER,
    "author_id" VARCHAR NOT NULL,
    "created_at" timestamp not null default current_timestamp(0)::timestamp without time zone,
    "name" VARCHAR,
    "size" INTEGER,
    "type" VARCHAR,
    "thumbnail" VARCHAR
);
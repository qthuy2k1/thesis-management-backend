CREATE TABLE "users" (
  "id" serial PRIMARY KEY not null,
  "fullname" varchar not null,
  "email" varchar not null,
  "role" varchar not null,
  "classroom_id" integer not null,
  "score" integer,
  "reset_token" varchar,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE
);
CREATE TABLE "users" (
  "id" serial PRIMARY KEY not null,
  "fullname" varchar not null,
  "email" varchar not null,
  "role" varchar not null,
  "classroom_id" integer not null,
  "score" integer,
  "reset_token" varchar,
  "created_at" timestamp not null default current_timestamp,
  "updated_at" timestamp not null default current_timestamp
);
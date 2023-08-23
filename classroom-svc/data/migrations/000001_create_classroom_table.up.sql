CREATE TABLE "classrooms" (
  "id" serial primary key not null,
  "title" varchar unique not null,
  "description" varchar,
  "status" varchar not null,
  "created_at" timestamp not null default current_timestamp,
  "updated_at" timestamp not null default current_timestamp
);

CREATE TABLE "waiting_lists" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "classroom_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::timestamp without time zone
);
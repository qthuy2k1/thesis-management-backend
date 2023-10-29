CREATE TABLE IF NOT EXISTS "student_defs" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "user_id" VARCHAR NOT NULL,
    "instructor_id" VARCHAR NOT NULL
);
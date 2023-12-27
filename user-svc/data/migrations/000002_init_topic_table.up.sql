CREATE TABLE "topics" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "title" VARCHAR NOT NULL,
    "type_topic" VARCHAR NOT NULL,
    "member_quantity" INTEGER NOT NULL,
    "student_id" VARCHAR NOT NULL,
    "member_email" VARCHAR NOT NULL,
    "description" VARCHAR
);
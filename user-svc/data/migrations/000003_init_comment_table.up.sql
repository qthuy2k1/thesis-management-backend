CREATE TABLE "comments" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "user_id" VARCHAR NOT NULL,
    "post_id" INTEGER,
    "exercise_id" INTEGER,
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(0)::TIMESTAMP WITHOUT TIME ZONE
);

CREATE INDEX comments_post_id ON comments(post_id);
CREATE INDEX comments_execise_id ON comments(exercise_id);
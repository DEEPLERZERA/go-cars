CREATE TABLE "cars" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar NOT NULL,
    "price" integer NOT NULL,
    "brand" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

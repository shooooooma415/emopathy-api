CREATE TABLE "groups" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" varchar NOT NULL,
  "password" varchar NOT NULL
);

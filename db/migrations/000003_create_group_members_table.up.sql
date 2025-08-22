CREATE TABLE "group_members" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "group_id" uuid NOT NULL REFERENCES "groups"("id") ON DELETE CASCADE,
  "user_id" uuid NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
  "admin" boolean NOT NULL DEFAULT false,
  UNIQUE("group_id", "user_id")
);

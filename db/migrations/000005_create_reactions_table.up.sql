CREATE TABLE "reactions" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "event_id" uuid NOT NULL REFERENCES "user_events"("id") ON DELETE CASCADE,
  "user_id" uuid NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT NOW(),
  "type" integer NOT NULL,
  UNIQUE("event_id", "user_id")
);

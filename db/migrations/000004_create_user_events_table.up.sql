CREATE TABLE "user_events" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT NOW(),
  "event_name" varchar NOT NULL,
  "emotion" varchar NOT NULL
);

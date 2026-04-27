-- Create index for verification token lookups by identifier.
CREATE INDEX IF NOT EXISTS "verification_identifier_idx" ON "verification" ("identifier");

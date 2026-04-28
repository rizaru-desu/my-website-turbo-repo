DO $$
BEGIN
  IF to_regclass('"user"') IS NULL THEN
    RETURN;
  END IF;

  ALTER TABLE "user"
    ADD COLUMN IF NOT EXISTS "twoFactorEnabled" boolean;

  CREATE TABLE IF NOT EXISTS "twoFactor" (
    "id" text NOT NULL,
    "secret" text NOT NULL,
    "backupCodes" text NOT NULL,
    "userId" text NOT NULL,
    CONSTRAINT "twoFactor_pkey" PRIMARY KEY ("id")
  );

  IF NOT EXISTS (
    SELECT 1
    FROM pg_constraint
    WHERE conname = 'twoFactor_userId_fkey'
  ) THEN
    ALTER TABLE "twoFactor"
      ADD CONSTRAINT "twoFactor_userId_fkey"
      FOREIGN KEY ("userId") REFERENCES "user" ("id")
      ON UPDATE CASCADE ON DELETE CASCADE;
  END IF;

  CREATE UNIQUE INDEX IF NOT EXISTS "twoFactor_userId_key" ON "twoFactor" ("userId");
END $$;

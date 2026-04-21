-- Enum types
CREATE TYPE "BlogCommentStatus" AS ENUM ('PENDING', 'APPROVED', 'REJECTED', 'SPAM');
CREATE TYPE "BlogPostStatus" AS ENUM ('DRAFT', 'PUBLISHED', 'ARCHIVED');
CREATE TYPE "MessageStatus" AS ENUM ('UNREAD', 'READ', 'ARCHIVED');
CREATE TYPE "ProjectAccent" AS ENUM ('RED', 'BLUE', 'CREAM');
CREATE TYPE "ProjectStatus" AS ENUM ('DRAFT', 'PUBLISHED', 'ARCHIVED');
CREATE TYPE "TestimonialRelation" AS ENUM ('CLIENT', 'COLLEAGUE', 'MENTOR', 'OTHER');
CREATE TYPE "TestimonialStatus" AS ENUM ('PENDING', 'APPROVED', 'REJECTED');

-- Table: public."account"
CREATE TABLE "account" (
  "id" text NOT NULL,
  "accountId" text NOT NULL,
  "providerId" text NOT NULL,
  "userId" text NOT NULL,
  "accessToken" text,
  "refreshToken" text,
  "idToken" text,
  "accessTokenExpiresAt" timestamp(3) without time zone,
  "refreshTokenExpiresAt" timestamp(3) without time zone,
  "scope" text,
  "password" text,
  "createdAt" timestamp(3) without time zone NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "account_pkey" PRIMARY KEY (id)
);

-- Table: public."blogComment"
CREATE TABLE "blogComment" (
  "id" text NOT NULL,
  "blogPostId" text NOT NULL,
  "parentId" text,
  "displayName" text NOT NULL,
  "email" text NOT NULL,
  "body" text NOT NULL,
  "status" "BlogCommentStatus" DEFAULT 'PENDING'::"BlogCommentStatus" NOT NULL,
  "reviewedAt" timestamp(3) without time zone,
  "reviewedByUserId" text,
  "ipHash" text NOT NULL,
  "fingerprint" text NOT NULL,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "blogComment_pkey" PRIMARY KEY (id)
);

-- Table: public."blogPost"
CREATE TABLE "blogPost" (
  "id" text NOT NULL,
  "title" text NOT NULL,
  "slug" text NOT NULL,
  "excerpt" text NOT NULL,
  "content" text NOT NULL,
  "tags" text[],
  "category" text NOT NULL,
  "coverImagePlaceholder" text,
  "status" "BlogPostStatus" DEFAULT 'DRAFT'::"BlogPostStatus" NOT NULL,
  "featured" boolean DEFAULT false NOT NULL,
  "publishDate" timestamp(3) without time zone,
  "readingTime" text NOT NULL,
  "seoTitle" text,
  "seoDescription" text,
  "authorName" text NOT NULL,
  "authorUserId" text NOT NULL,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "blogPost_pkey" PRIMARY KEY (id)
);

-- Table: public."certificate"
CREATE TABLE "certificate" (
  "id" text NOT NULL,
  "name" text NOT NULL,
  "issuer" text NOT NULL,
  "year" text NOT NULL,
  "verificationLink" text NOT NULL,
  "credentialId" text,
  "featured" boolean DEFAULT false NOT NULL,
  "sortOrder" integer DEFAULT 0 NOT NULL,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "certificate_pkey" PRIMARY KEY (id)
);

-- Table: public."cvDownloadLog"
CREATE TABLE "cvDownloadLog" (
  "id" text NOT NULL,
  "ipHash" text NOT NULL,
  "userAgent" text NOT NULL,
  "referrer" text DEFAULT 'direct'::text NOT NULL,
  "downloadedAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  CONSTRAINT "cvDownloadLog_pkey" PRIMARY KEY (id)
);

-- Table: public."education"
CREATE TABLE "education" (
  "id" text NOT NULL,
  "degree" text NOT NULL,
  "school" text NOT NULL,
  "period" text NOT NULL,
  "description" text NOT NULL,
  "highlights" text[],
  "sortOrder" integer DEFAULT 0 NOT NULL,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "education_pkey" PRIMARY KEY (id)
);

-- Table: public."experience"
CREATE TABLE "experience" (
  "id" text NOT NULL,
  "role" text NOT NULL,
  "company" text NOT NULL,
  "period" text NOT NULL,
  "location" text NOT NULL,
  "summary" text NOT NULL,
  "achievements" text[],
  "sortOrder" integer DEFAULT 0 NOT NULL,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "experience_pkey" PRIMARY KEY (id)
);

-- Table: public."message"
CREATE TABLE "message" (
  "id" text NOT NULL,
  "senderName" text NOT NULL,
  "senderEmail" text NOT NULL,
  "subject" text NOT NULL,
  "body" text NOT NULL,
  "status" "MessageStatus" DEFAULT 'UNREAD'::"MessageStatus" NOT NULL,
  "readAt" timestamp(3) without time zone,
  "archivedAt" timestamp(3) without time zone,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "message_pkey" PRIMARY KEY (id)
);

-- Table: public."profileContent"
CREATE TABLE "profileContent" (
  "id" text NOT NULL,
  "storageKey" text DEFAULT 'primary'::text NOT NULL,
  "fullName" text NOT NULL,
  "headline" text NOT NULL,
  "shortIntro" text NOT NULL,
  "about" text NOT NULL,
  "location" text NOT NULL,
  "email" text NOT NULL,
  "phone" text NOT NULL,
  "availability" text NOT NULL,
  "primaryCta" text NOT NULL,
  "socialLinks" jsonb NOT NULL,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  "profilePhotoUrl" text,
  "focus" text[] DEFAULT ARRAY[]::text[],
  "stats" jsonb DEFAULT '[]'::jsonb NOT NULL,
  CONSTRAINT "profileContent_pkey" PRIMARY KEY (id)
);

-- Table: public."project"
CREATE TABLE "project" (
  "id" text NOT NULL,
  "title" text NOT NULL,
  "slug" text NOT NULL,
  "summary" text NOT NULL,
  "category" text NOT NULL,
  "year" text NOT NULL,
  "clientOrCompany" text NOT NULL,
  "role" text NOT NULL,
  "duration" text NOT NULL,
  "thumbnailPlaceholder" text NOT NULL,
  "projectUrl" text,
  "githubUrl" text,
  "impactSummary" text NOT NULL,
  "tags" text[],
  "impactBullets" text[],
  "techStack" text[],
  "process" text[],
  "featured" boolean DEFAULT false NOT NULL,
  "status" "ProjectStatus" DEFAULT 'DRAFT'::"ProjectStatus" NOT NULL,
  "sortOrder" integer,
  "accent" "ProjectAccent" DEFAULT 'RED'::"ProjectAccent" NOT NULL,
  "challenge" text NOT NULL,
  "outcome" text NOT NULL,
  "metrics" jsonb NOT NULL,
  "gallery" jsonb NOT NULL,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "project_pkey" PRIMARY KEY (id)
);

-- Table: public."resumeAsset"
CREATE TABLE "resumeAsset" (
  "id" text NOT NULL,
  "storageKey" text DEFAULT 'primary'::text NOT NULL,
  "downloadUrl" text NOT NULL,
  "fileName" text,
  "fileSizeBytes" integer,
  "mimeType" text DEFAULT 'application/pdf'::text NOT NULL,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "resumeAsset_pkey" PRIMARY KEY (id)
);

-- Table: public."session"
CREATE TABLE "session" (
  "id" text NOT NULL,
  "expiresAt" timestamp(3) without time zone NOT NULL,
  "token" text NOT NULL,
  "createdAt" timestamp(3) without time zone NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  "ipAddress" text,
  "userAgent" text,
  "userId" text NOT NULL,
  "impersonatedBy" text,
  CONSTRAINT "session_pkey" PRIMARY KEY (id)
);

-- Table: public."skill"
CREATE TABLE "skill" (
  "id" text NOT NULL,
  "name" text NOT NULL,
  "category" text NOT NULL,
  "level" text NOT NULL,
  "featured" boolean DEFAULT false NOT NULL,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "skill_pkey" PRIMARY KEY (id)
);

-- Table: public."testimonial"
CREATE TABLE "testimonial" (
  "id" text NOT NULL,
  "name" text NOT NULL,
  "role" text NOT NULL,
  "company" text,
  "message" text NOT NULL,
  "rating" integer NOT NULL,
  "relation" "TestimonialRelation" NOT NULL,
  "status" "TestimonialStatus" DEFAULT 'PENDING'::"TestimonialStatus" NOT NULL,
  "featured" boolean DEFAULT false NOT NULL,
  "reviewedAt" timestamp(3) without time zone,
  "createdAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "testimonial_pkey" PRIMARY KEY (id)
);

-- Table: public."twoFactor"
CREATE TABLE "twoFactor" (
  "id" text NOT NULL,
  "secret" text NOT NULL,
  "backupCodes" text NOT NULL,
  "userId" text NOT NULL,
  CONSTRAINT "twoFactor_pkey" PRIMARY KEY (id)
);

-- Table: public."user"
CREATE TABLE "user" (
  "id" text NOT NULL,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "emailVerified" boolean NOT NULL,
  "image" text,
  "createdAt" timestamp(3) without time zone NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  "twoFactorEnabled" boolean,
  "banExpires" timestamp(3) without time zone,
  "banReason" text,
  "banned" boolean,
  "displayUsername" text,
  "isAnonymous" boolean,
  "role" text,
  "username" text,
  CONSTRAINT "user_pkey" PRIMARY KEY (id)
);

-- Table: public."verification"
CREATE TABLE "verification" (
  "id" text NOT NULL,
  "identifier" text NOT NULL,
  "value" text NOT NULL,
  "expiresAt" timestamp(3) without time zone NOT NULL,
  "createdAt" timestamp(3) without time zone NOT NULL,
  "updatedAt" timestamp(3) without time zone NOT NULL,
  CONSTRAINT "verification_pkey" PRIMARY KEY (id)
);

-- Table: public."visitorLog"
CREATE TABLE "visitorLog" (
  "id" text NOT NULL,
  "visitorId" text NOT NULL,
  "ipHash" text NOT NULL,
  "path" text NOT NULL,
  "referrer" text DEFAULT 'direct'::text NOT NULL,
  "referrerSource" text DEFAULT 'direct'::text NOT NULL,
  "userAgent" text NOT NULL,
  "isUniqueDailyVisitor" boolean DEFAULT false NOT NULL,
  "visitedAt" timestamp(3) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
  CONSTRAINT "visitorLog_pkey" PRIMARY KEY (id)
);

-- Foreign keys
ALTER TABLE "account" ADD CONSTRAINT "account_userId_fkey" FOREIGN KEY ("userId") REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "blogComment" ADD CONSTRAINT "blogComment_blogPostId_fkey" FOREIGN KEY ("blogPostId") REFERENCES "blogPost"(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "blogComment" ADD CONSTRAINT "blogComment_parentId_fkey" FOREIGN KEY ("parentId") REFERENCES "blogComment"(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "blogComment" ADD CONSTRAINT "blogComment_reviewedByUserId_fkey" FOREIGN KEY ("reviewedByUserId") REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE "blogPost" ADD CONSTRAINT "blogPost_authorUserId_fkey" FOREIGN KEY ("authorUserId") REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "session" ADD CONSTRAINT "session_userId_fkey" FOREIGN KEY ("userId") REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE "twoFactor" ADD CONSTRAINT "twoFactor_userId_fkey" FOREIGN KEY ("userId") REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE;

-- Indexes
CREATE INDEX "blogComment_blogPostId_idx" ON public."blogComment" USING btree ("blogPostId");
CREATE INDEX "blogComment_blogPostId_parentId_status_createdAt_idx" ON public."blogComment" USING btree ("blogPostId", "parentId", status, "createdAt");
CREATE INDEX "blogComment_blogPostId_status_createdAt_idx" ON public."blogComment" USING btree ("blogPostId", status, "createdAt");
CREATE INDEX "blogComment_fingerprint_idx" ON public."blogComment" USING btree (fingerprint);
CREATE INDEX "blogComment_ipHash_createdAt_idx" ON public."blogComment" USING btree ("ipHash", "createdAt");
CREATE INDEX "blogComment_parentId_idx" ON public."blogComment" USING btree ("parentId");
CREATE INDEX "blogComment_reviewedByUserId_idx" ON public."blogComment" USING btree ("reviewedByUserId");
CREATE INDEX "blogComment_status_idx" ON public."blogComment" USING btree (status);
CREATE INDEX "blogPost_authorUserId_idx" ON public."blogPost" USING btree ("authorUserId");
CREATE INDEX "blogPost_featured_idx" ON public."blogPost" USING btree (featured);
CREATE INDEX "blogPost_publishDate_idx" ON public."blogPost" USING btree ("publishDate");
CREATE UNIQUE INDEX "blogPost_slug_key" ON public."blogPost" USING btree (slug);
CREATE INDEX "blogPost_status_featured_publishDate_idx" ON public."blogPost" USING btree (status, featured, "publishDate");
CREATE INDEX "blogPost_status_idx" ON public."blogPost" USING btree (status);
CREATE INDEX certificate_featured_idx ON public.certificate USING btree (featured);
CREATE INDEX "certificate_sortOrder_idx" ON public.certificate USING btree ("sortOrder");
CREATE INDEX "cvDownloadLog_downloadedAt_idx" ON public."cvDownloadLog" USING btree ("downloadedAt");
CREATE INDEX "cvDownloadLog_ipHash_downloadedAt_idx" ON public."cvDownloadLog" USING btree ("ipHash", "downloadedAt");
CREATE INDEX "education_sortOrder_idx" ON public.education USING btree ("sortOrder");
CREATE INDEX "experience_sortOrder_idx" ON public.experience USING btree ("sortOrder");
CREATE INDEX "message_createdAt_idx" ON public.message USING btree ("createdAt");
CREATE INDEX "message_status_createdAt_idx" ON public.message USING btree (status, "createdAt");
CREATE INDEX message_status_idx ON public.message USING btree (status);
CREATE UNIQUE INDEX "profileContent_storageKey_key" ON public."profileContent" USING btree ("storageKey");
CREATE INDEX project_featured_idx ON public.project USING btree (featured);
CREATE UNIQUE INDEX project_slug_key ON public.project USING btree (slug);
CREATE INDEX "project_sortOrder_idx" ON public.project USING btree ("sortOrder");
CREATE INDEX "project_status_featured_sortOrder_idx" ON public.project USING btree (status, featured, "sortOrder");
CREATE INDEX project_status_idx ON public.project USING btree (status);
CREATE INDEX "project_updatedAt_idx" ON public.project USING btree ("updatedAt");
CREATE UNIQUE INDEX "resumeAsset_storageKey_key" ON public."resumeAsset" USING btree ("storageKey");
CREATE UNIQUE INDEX session_token_key ON public.session USING btree (token);
CREATE INDEX skill_category_featured_idx ON public.skill USING btree (category, featured);
CREATE INDEX skill_category_idx ON public.skill USING btree (category);
CREATE INDEX skill_featured_idx ON public.skill USING btree (featured);
CREATE INDEX "testimonial_createdAt_idx" ON public.testimonial USING btree ("createdAt");
CREATE INDEX testimonial_featured_idx ON public.testimonial USING btree (featured);
CREATE INDEX testimonial_status_featured_idx ON public.testimonial USING btree (status, featured);
CREATE INDEX testimonial_status_idx ON public.testimonial USING btree (status);
CREATE UNIQUE INDEX user_email_key ON public."user" USING btree (email);
CREATE UNIQUE INDEX user_username_key ON public."user" USING btree (username);
CREATE INDEX "visitorLog_path_visitedAt_idx" ON public."visitorLog" USING btree (path, "visitedAt");
CREATE INDEX "visitorLog_referrerSource_visitedAt_idx" ON public."visitorLog" USING btree ("referrerSource", "visitedAt");
CREATE INDEX "visitorLog_visitedAt_idx" ON public."visitorLog" USING btree ("visitedAt");
CREATE INDEX "visitorLog_visitorId_visitedAt_idx" ON public."visitorLog" USING btree ("visitorId", "visitedAt");

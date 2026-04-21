package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

var (
	postgresTextArray = map[string]string{dialect.Postgres: "text[]"}
	postgresJSONB     = map[string]string{dialect.Postgres: "jsonb"}
	postgresTimestamp = map[string]string{dialect.Postgres: "timestamp(3) without time zone"}
)

func table(name string) []entschema.Annotation {
	return []entschema.Annotation{entsql.Annotation{Table: name}}
}

func textID() ent.Field {
	return field.String("id").NotEmpty()
}

func textColumn(name string) ent.Field {
	return field.String(name).NotEmpty()
}

func optionalText(name string) ent.Field {
	return field.String(name).Optional().Nillable()
}

func camelText(name, storageKey string) ent.Field {
	return field.String(name).StorageKey(storageKey).NotEmpty()
}

func optionalCamelText(name, storageKey string) ent.Field {
	return field.String(name).StorageKey(storageKey).Optional().Nillable()
}

func textArray(name string) ent.Field {
	return field.Strings(name).SchemaType(postgresTextArray).Optional()
}

func camelTextArray(name, storageKey string) ent.Field {
	return field.Strings(name).StorageKey(storageKey).SchemaType(postgresTextArray).Optional()
}

func requiredTime(name, storageKey string) ent.Field {
	return field.Time(name).StorageKey(storageKey).SchemaType(postgresTimestamp)
}

func defaultTime(name, storageKey string) ent.Field {
	return field.Time(name).
		StorageKey(storageKey).
		SchemaType(postgresTimestamp).
		Default(time.Now).
		Annotations(entsql.Default("CURRENT_TIMESTAMP"))
}

func updateTime(name, storageKey string) ent.Field {
	return field.Time(name).StorageKey(storageKey).SchemaType(postgresTimestamp).UpdateDefault(time.Now)
}

func optionalTime(name, storageKey string) ent.Field {
	return field.Time(name).StorageKey(storageKey).SchemaType(postgresTimestamp).Optional().Nillable()
}

func pgEnum(name, pgType, def string, values ...string) ent.Field {
	enum := field.Enum(name).
		Values(values...).
		SchemaType(map[string]string{dialect.Postgres: `\"` + pgType + `\"`})
	if def != "" {
		enum = enum.Default(def)
	}
	return enum
}

// Account holds the schema definition for the account table.
type Account struct{ ent.Schema }

func (Account) Annotations() []entschema.Annotation { return table("account") }

func (Account) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		camelText("account_id", "accountId"),
		camelText("provider_id", "providerId"),
		camelText("user_id", "userId"),
		optionalCamelText("access_token", "accessToken"),
		optionalCamelText("refresh_token", "refreshToken"),
		optionalCamelText("id_token", "idToken"),
		optionalTime("access_token_expires_at", "accessTokenExpiresAt"),
		optionalTime("refresh_token_expires_at", "refreshTokenExpiresAt"),
		optionalText("scope"),
		optionalText("password"),
		requiredTime("created_at", "createdAt"),
		requiredTime("updated_at", "updatedAt"),
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("accounts").
			Field("user_id").
			Required().
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

// BlogComment holds the schema definition for the blogComment table.
type BlogComment struct{ ent.Schema }

func (BlogComment) Annotations() []entschema.Annotation { return table("blogComment") }

func (BlogComment) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		camelText("blog_post_id", "blogPostId"),
		optionalCamelText("parent_id", "parentId"),
		camelText("display_name", "displayName"),
		textColumn("email"),
		textColumn("body"),
		pgEnum("status", "BlogCommentStatus", "PENDING", "PENDING", "APPROVED", "REJECTED", "SPAM"),
		optionalCamelText("reviewed_by_user_id", "reviewedByUserId"),
		optionalTime("reviewed_at", "reviewedAt"),
		camelText("ip_hash", "ipHash"),
		textColumn("fingerprint"),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
	}
}

func (BlogComment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", BlogComment.Type).
			StorageKey(edge.Symbol("blogComment_parentId_fkey")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("blog_post", BlogPost.Type).
			Ref("comments").
			Field("blog_post_id").
			Required().
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("parent", BlogComment.Type).
			Ref("children").
			Field("parent_id").
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("reviewed_by", User.Type).
			Ref("reviewed_comments").
			Field("reviewed_by_user_id").
			Unique().
			Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}

func (BlogComment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("blog_post_id").StorageKey("blogComment_blogPostId_idx"),
		index.Fields("blog_post_id", "parent_id", "status", "created_at").StorageKey("blogComment_blogPostId_parentId_status_createdAt_idx"),
		index.Fields("blog_post_id", "status", "created_at").StorageKey("blogComment_blogPostId_status_createdAt_idx"),
		index.Fields("fingerprint").StorageKey("blogComment_fingerprint_idx"),
		index.Fields("ip_hash", "created_at").StorageKey("blogComment_ipHash_createdAt_idx"),
		index.Fields("parent_id").StorageKey("blogComment_parentId_idx"),
		index.Fields("reviewed_by_user_id").StorageKey("blogComment_reviewedByUserId_idx"),
		index.Fields("status").StorageKey("blogComment_status_idx"),
	}
}

// BlogPost holds the schema definition for the blogPost table.
type BlogPost struct{ ent.Schema }

func (BlogPost) Annotations() []entschema.Annotation { return table("blogPost") }

func (BlogPost) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		textColumn("title"),
		textColumn("slug"),
		textColumn("excerpt"),
		textColumn("content"),
		textArray("tags"),
		textColumn("category"),
		optionalCamelText("cover_image_placeholder", "coverImagePlaceholder"),
		pgEnum("status", "BlogPostStatus", "DRAFT", "DRAFT", "PUBLISHED", "ARCHIVED"),
		field.Bool("featured").Default(false),
		optionalTime("publish_date", "publishDate"),
		camelText("reading_time", "readingTime"),
		optionalCamelText("seo_title", "seoTitle"),
		optionalCamelText("seo_description", "seoDescription"),
		camelText("author_name", "authorName"),
		camelText("author_user_id", "authorUserId"),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
	}
}

func (BlogPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", BlogComment.Type).
			StorageKey(edge.Symbol("blogComment_blogPostId_fkey")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("author", User.Type).
			Ref("blog_posts").
			Field("author_user_id").
			Required().
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (BlogPost) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("author_user_id").StorageKey("blogPost_authorUserId_idx"),
		index.Fields("featured").StorageKey("blogPost_featured_idx"),
		index.Fields("publish_date").StorageKey("blogPost_publishDate_idx"),
		index.Fields("slug").Unique().StorageKey("blogPost_slug_key"),
		index.Fields("status", "featured", "publish_date").StorageKey("blogPost_status_featured_publishDate_idx"),
		index.Fields("status").StorageKey("blogPost_status_idx"),
	}
}

// Certificate holds the schema definition for the certificate table.
type Certificate struct{ ent.Schema }

func (Certificate) Annotations() []entschema.Annotation { return table("certificate") }

func (Certificate) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		textColumn("name"),
		textColumn("issuer"),
		textColumn("year"),
		camelText("verification_link", "verificationLink"),
		optionalCamelText("credential_id", "credentialId"),
		field.Bool("featured").Default(false),
		field.Int("sort_order").StorageKey("sortOrder").Default(0),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
	}
}

func (Certificate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("featured").StorageKey("certificate_featured_idx"),
		index.Fields("sort_order").StorageKey("certificate_sortOrder_idx"),
	}
}

// CVDownloadLog holds the schema definition for the cvDownloadLog table.
type CVDownloadLog struct{ ent.Schema }

func (CVDownloadLog) Annotations() []entschema.Annotation { return table("cvDownloadLog") }

func (CVDownloadLog) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		camelText("ip_hash", "ipHash"),
		camelText("user_agent", "userAgent"),
		field.String("referrer").Default("direct").NotEmpty(),
		defaultTime("downloaded_at", "downloadedAt"),
	}
}

func (CVDownloadLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("downloaded_at").StorageKey("cvDownloadLog_downloadedAt_idx"),
		index.Fields("ip_hash", "downloaded_at").StorageKey("cvDownloadLog_ipHash_downloadedAt_idx"),
	}
}

// Education holds the schema definition for the education table.
type Education struct{ ent.Schema }

func (Education) Annotations() []entschema.Annotation { return table("education") }

func (Education) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		textColumn("degree"),
		textColumn("school"),
		textColumn("period"),
		textColumn("description"),
		textArray("highlights"),
		field.Int("sort_order").StorageKey("sortOrder").Default(0),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
	}
}

func (Education) Indexes() []ent.Index {
	return []ent.Index{index.Fields("sort_order").StorageKey("education_sortOrder_idx")}
}

// Experience holds the schema definition for the experience table.
type Experience struct{ ent.Schema }

func (Experience) Annotations() []entschema.Annotation { return table("experience") }

func (Experience) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		textColumn("role"),
		textColumn("company"),
		textColumn("period"),
		textColumn("location"),
		textColumn("summary"),
		textArray("achievements"),
		field.Int("sort_order").StorageKey("sortOrder").Default(0),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
	}
}

func (Experience) Indexes() []ent.Index {
	return []ent.Index{index.Fields("sort_order").StorageKey("experience_sortOrder_idx")}
}

// Message holds the schema definition for the message table.
type Message struct{ ent.Schema }

func (Message) Annotations() []entschema.Annotation { return table("message") }

func (Message) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		camelText("sender_name", "senderName"),
		camelText("sender_email", "senderEmail"),
		textColumn("subject"),
		textColumn("body"),
		pgEnum("status", "MessageStatus", "UNREAD", "UNREAD", "READ", "ARCHIVED"),
		optionalTime("read_at", "readAt"),
		optionalTime("archived_at", "archivedAt"),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
	}
}

func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at").StorageKey("message_createdAt_idx"),
		index.Fields("status", "created_at").StorageKey("message_status_createdAt_idx"),
		index.Fields("status").StorageKey("message_status_idx"),
	}
}

// ProfileContent holds the schema definition for the profileContent table.
type ProfileContent struct{ ent.Schema }

func (ProfileContent) Annotations() []entschema.Annotation { return table("profileContent") }

func (ProfileContent) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		field.String("storage_key").StorageKey("storageKey").Default("primary").NotEmpty(),
		camelText("full_name", "fullName"),
		textColumn("headline"),
		camelText("short_intro", "shortIntro"),
		textColumn("about"),
		textColumn("location"),
		textColumn("email"),
		textColumn("phone"),
		textColumn("availability"),
		camelText("primary_cta", "primaryCta"),
		field.JSON("social_links", map[string]any{}).StorageKey("socialLinks").SchemaType(postgresJSONB),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
		optionalCamelText("profile_photo_url", "profilePhotoUrl"),
		field.Strings("focus").SchemaType(postgresTextArray).Default([]string{}).Annotations(entsql.Default("ARRAY[]::text[]")),
		field.JSON("stats", []any{}).SchemaType(postgresJSONB).Default([]any{}).Annotations(entsql.Default("'[]'::jsonb")),
	}
}

func (ProfileContent) Indexes() []ent.Index {
	return []ent.Index{index.Fields("storage_key").Unique().StorageKey("profileContent_storageKey_key")}
}

// Project holds the schema definition for the project table.
type Project struct{ ent.Schema }

func (Project) Annotations() []entschema.Annotation { return table("project") }

func (Project) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		textColumn("title"),
		textColumn("slug"),
		textColumn("summary"),
		textColumn("category"),
		textColumn("year"),
		camelText("client_or_company", "clientOrCompany"),
		textColumn("role"),
		textColumn("duration"),
		camelText("thumbnail_placeholder", "thumbnailPlaceholder"),
		optionalCamelText("project_url", "projectUrl"),
		optionalCamelText("github_url", "githubUrl"),
		camelText("impact_summary", "impactSummary"),
		textArray("tags"),
		camelTextArray("impact_bullets", "impactBullets"),
		camelTextArray("tech_stack", "techStack"),
		textArray("process"),
		field.Bool("featured").Default(false),
		pgEnum("status", "ProjectStatus", "DRAFT", "DRAFT", "PUBLISHED", "ARCHIVED"),
		field.Int("sort_order").StorageKey("sortOrder").Optional().Nillable(),
		pgEnum("accent", "ProjectAccent", "RED", "RED", "BLUE", "CREAM"),
		textColumn("challenge"),
		textColumn("outcome"),
		field.JSON("metrics", map[string]any{}).SchemaType(postgresJSONB),
		field.JSON("gallery", []any{}).SchemaType(postgresJSONB),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
	}
}

func (Project) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("featured").StorageKey("project_featured_idx"),
		index.Fields("slug").Unique().StorageKey("project_slug_key"),
		index.Fields("sort_order").StorageKey("project_sortOrder_idx"),
		index.Fields("status", "featured", "sort_order").StorageKey("project_status_featured_sortOrder_idx"),
		index.Fields("status").StorageKey("project_status_idx"),
		index.Fields("updated_at").StorageKey("project_updatedAt_idx"),
	}
}

// ResumeAsset holds the schema definition for the resumeAsset table.
type ResumeAsset struct{ ent.Schema }

func (ResumeAsset) Annotations() []entschema.Annotation { return table("resumeAsset") }

func (ResumeAsset) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		field.String("storage_key").StorageKey("storageKey").Default("primary").NotEmpty(),
		camelText("download_url", "downloadUrl"),
		optionalCamelText("file_name", "fileName"),
		field.Int("file_size_bytes").StorageKey("fileSizeBytes").Optional().Nillable(),
		field.String("mime_type").StorageKey("mimeType").Default("application/pdf").NotEmpty(),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
	}
}

func (ResumeAsset) Indexes() []ent.Index {
	return []ent.Index{index.Fields("storage_key").Unique().StorageKey("resumeAsset_storageKey_key")}
}

// Session holds the schema definition for the session table.
type Session struct{ ent.Schema }

func (Session) Annotations() []entschema.Annotation { return table("session") }

func (Session) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		requiredTime("expires_at", "expiresAt"),
		textColumn("token"),
		requiredTime("created_at", "createdAt"),
		requiredTime("updated_at", "updatedAt"),
		optionalCamelText("ip_address", "ipAddress"),
		optionalCamelText("user_agent", "userAgent"),
		camelText("user_id", "userId"),
		optionalCamelText("impersonated_by", "impersonatedBy"),
	}
}

func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("sessions").
			Field("user_id").
			Required().
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (Session) Indexes() []ent.Index {
	return []ent.Index{index.Fields("token").Unique().StorageKey("session_token_key")}
}

// Skill holds the schema definition for the skill table.
type Skill struct{ ent.Schema }

func (Skill) Annotations() []entschema.Annotation { return table("skill") }

func (Skill) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		textColumn("name"),
		textColumn("category"),
		textColumn("level"),
		field.Bool("featured").Default(false),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
	}
}

func (Skill) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("category", "featured").StorageKey("skill_category_featured_idx"),
		index.Fields("category").StorageKey("skill_category_idx"),
		index.Fields("featured").StorageKey("skill_featured_idx"),
	}
}

// Testimonial holds the schema definition for the testimonial table.
type Testimonial struct{ ent.Schema }

func (Testimonial) Annotations() []entschema.Annotation { return table("testimonial") }

func (Testimonial) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		textColumn("name"),
		textColumn("role"),
		optionalText("company"),
		textColumn("message"),
		field.Int("rating"),
		pgEnum("relation", "TestimonialRelation", "", "CLIENT", "COLLEAGUE", "MENTOR", "OTHER"),
		pgEnum("status", "TestimonialStatus", "PENDING", "PENDING", "APPROVED", "REJECTED"),
		field.Bool("featured").Default(false),
		optionalTime("reviewed_at", "reviewedAt"),
		defaultTime("created_at", "createdAt"),
		updateTime("updated_at", "updatedAt"),
	}
}

func (Testimonial) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at").StorageKey("testimonial_createdAt_idx"),
		index.Fields("featured").StorageKey("testimonial_featured_idx"),
		index.Fields("status", "featured").StorageKey("testimonial_status_featured_idx"),
		index.Fields("status").StorageKey("testimonial_status_idx"),
	}
}

// TwoFactor holds the schema definition for the twoFactor table.
type TwoFactor struct{ ent.Schema }

func (TwoFactor) Annotations() []entschema.Annotation { return table("twoFactor") }

func (TwoFactor) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		textColumn("secret"),
		camelText("backup_codes", "backupCodes"),
		camelText("user_id", "userId"),
	}
}

func (TwoFactor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("two_factors").
			Field("user_id").
			Required().
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

// User holds the schema definition for the user table.
type User struct{ ent.Schema }

func (User) Annotations() []entschema.Annotation { return table("user") }

func (User) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		textColumn("name"),
		textColumn("email"),
		field.Bool("email_verified").StorageKey("emailVerified"),
		optionalText("image"),
		requiredTime("created_at", "createdAt"),
		requiredTime("updated_at", "updatedAt"),
		field.Bool("two_factor_enabled").StorageKey("twoFactorEnabled").Optional().Nillable(),
		optionalTime("ban_expires", "banExpires"),
		optionalCamelText("ban_reason", "banReason"),
		field.Bool("banned").Optional().Nillable(),
		optionalCamelText("display_username", "displayUsername"),
		field.Bool("is_anonymous").StorageKey("isAnonymous").Optional().Nillable(),
		optionalText("role"),
		optionalText("username"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique().StorageKey("user_email_key"),
		index.Fields("username").Unique().StorageKey("user_username_key"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type).
			StorageKey(edge.Symbol("account_userId_fkey")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("blog_posts", BlogPost.Type).
			StorageKey(edge.Symbol("blogPost_authorUserId_fkey")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("reviewed_comments", BlogComment.Type).
			StorageKey(edge.Symbol("blogComment_reviewedByUserId_fkey")).
			Annotations(entsql.OnDelete(entsql.SetNull)),
		edge.To("sessions", Session.Type).
			StorageKey(edge.Symbol("session_userId_fkey")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("two_factors", TwoFactor.Type).
			StorageKey(edge.Symbol("twoFactor_userId_fkey")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

// Verification holds the schema definition for the verification table.
type Verification struct{ ent.Schema }

func (Verification) Annotations() []entschema.Annotation { return table("verification") }

func (Verification) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		textColumn("identifier"),
		textColumn("value"),
		requiredTime("expires_at", "expiresAt"),
		requiredTime("created_at", "createdAt"),
		requiredTime("updated_at", "updatedAt"),
	}
}

// VisitorLog holds the schema definition for the visitorLog table.
type VisitorLog struct{ ent.Schema }

func (VisitorLog) Annotations() []entschema.Annotation { return table("visitorLog") }

func (VisitorLog) Fields() []ent.Field {
	return []ent.Field{
		textID(),
		camelText("visitor_id", "visitorId"),
		camelText("ip_hash", "ipHash"),
		textColumn("path"),
		field.String("referrer").Default("direct").NotEmpty(),
		field.String("referrer_source").StorageKey("referrerSource").Default("direct").NotEmpty(),
		camelText("user_agent", "userAgent"),
		field.Bool("is_unique_daily_visitor").StorageKey("isUniqueDailyVisitor").Default(false),
		defaultTime("visited_at", "visitedAt"),
	}
}

func (VisitorLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("path", "visited_at").StorageKey("visitorLog_path_visitedAt_idx"),
		index.Fields("referrer_source", "visited_at").StorageKey("visitorLog_referrerSource_visitedAt_idx"),
		index.Fields("visited_at").StorageKey("visitorLog_visitedAt_idx"),
		index.Fields("visitor_id", "visited_at").StorageKey("visitorLog_visitorId_visitedAt_idx"),
	}
}

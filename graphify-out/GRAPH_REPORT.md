# Graph Report - portofolio-lightweight  (2026-05-03)

## Corpus Check
- 257 files · ~341,748 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 8558 nodes · 13312 edges · 68 communities detected
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1274 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Community 0|Community 0]]
- [[_COMMUNITY_Community 1|Community 1]]
- [[_COMMUNITY_Community 2|Community 2]]
- [[_COMMUNITY_Community 3|Community 3]]
- [[_COMMUNITY_Community 4|Community 4]]
- [[_COMMUNITY_Community 5|Community 5]]
- [[_COMMUNITY_Community 7|Community 7]]
- [[_COMMUNITY_Community 8|Community 8]]
- [[_COMMUNITY_Community 9|Community 9]]
- [[_COMMUNITY_Community 11|Community 11]]
- [[_COMMUNITY_Community 12|Community 12]]
- [[_COMMUNITY_Community 13|Community 13]]
- [[_COMMUNITY_Community 14|Community 14]]
- [[_COMMUNITY_Community 15|Community 15]]
- [[_COMMUNITY_Community 16|Community 16]]
- [[_COMMUNITY_Community 17|Community 17]]
- [[_COMMUNITY_Community 22|Community 22]]
- [[_COMMUNITY_Community 25|Community 25]]
- [[_COMMUNITY_Community 26|Community 26]]
- [[_COMMUNITY_Community 28|Community 28]]
- [[_COMMUNITY_Community 30|Community 30]]
- [[_COMMUNITY_Community 31|Community 31]]
- [[_COMMUNITY_Community 32|Community 32]]
- [[_COMMUNITY_Community 33|Community 33]]
- [[_COMMUNITY_Community 34|Community 34]]
- [[_COMMUNITY_Community 37|Community 37]]
- [[_COMMUNITY_Community 38|Community 38]]
- [[_COMMUNITY_Community 39|Community 39]]
- [[_COMMUNITY_Community 40|Community 40]]
- [[_COMMUNITY_Community 42|Community 42]]
- [[_COMMUNITY_Community 43|Community 43]]
- [[_COMMUNITY_Community 44|Community 44]]
- [[_COMMUNITY_Community 45|Community 45]]
- [[_COMMUNITY_Community 46|Community 46]]
- [[_COMMUNITY_Community 47|Community 47]]
- [[_COMMUNITY_Community 48|Community 48]]
- [[_COMMUNITY_Community 49|Community 49]]
- [[_COMMUNITY_Community 50|Community 50]]
- [[_COMMUNITY_Community 51|Community 51]]
- [[_COMMUNITY_Community 52|Community 52]]
- [[_COMMUNITY_Community 53|Community 53]]
- [[_COMMUNITY_Community 54|Community 54]]
- [[_COMMUNITY_Community 55|Community 55]]
- [[_COMMUNITY_Community 56|Community 56]]
- [[_COMMUNITY_Community 57|Community 57]]
- [[_COMMUNITY_Community 58|Community 58]]
- [[_COMMUNITY_Community 59|Community 59]]
- [[_COMMUNITY_Community 60|Community 60]]
- [[_COMMUNITY_Community 61|Community 61]]
- [[_COMMUNITY_Community 62|Community 62]]
- [[_COMMUNITY_Community 63|Community 63]]
- [[_COMMUNITY_Community 64|Community 64]]
- [[_COMMUNITY_Community 65|Community 65]]
- [[_COMMUNITY_Community 66|Community 66]]
- [[_COMMUNITY_Community 67|Community 67]]
- [[_COMMUNITY_Community 68|Community 68]]
- [[_COMMUNITY_Community 69|Community 69]]
- [[_COMMUNITY_Community 70|Community 70]]
- [[_COMMUNITY_Community 71|Community 71]]
- [[_COMMUNITY_Community 73|Community 73]]
- [[_COMMUNITY_Community 74|Community 74]]
- [[_COMMUNITY_Community 76|Community 76]]
- [[_COMMUNITY_Community 77|Community 77]]
- [[_COMMUNITY_Community 78|Community 78]]
- [[_COMMUNITY_Community 80|Community 80]]
- [[_COMMUNITY_Community 81|Community 81]]
- [[_COMMUNITY_Community 83|Community 83]]
- [[_COMMUNITY_Community 85|Community 85]]

## God Nodes (most connected - your core abstractions)
1. `setContextOp()` - 181 edges
2. `ProjectMutation` - 160 edges
3. `UserMutation` - 139 edges
4. `BlogPostMutation` - 123 edges
5. `hooks` - 109 edges
6. `BlogCommentMutation` - 105 edges
7. `ProfileContentMutation` - 100 edges
8. `AccountMutation` - 96 edges
9. `IsConstraintError()` - 92 edges
10. `TestimonialMutation` - 80 edges

## Surprising Connections (you probably didn't know these)
- `CmsApp()` --calls--> `renderPage()`  [INFERRED]
  ui-sample/pages/CmsApp.jsx → apps/frontend/admin/src/components/dashboard/AdminDashboard.tsx
- `newHandler()` --calls--> `NewLinuxMetricsProvider()`  [INFERRED]
  apps/backend/cmd/api/main.go → apps/backend/internal/infrastructure/monitoring/linux.go
- `GuestRoute()` --calls--> `useAuth()`  [INFERRED]
  apps/frontend/admin/src/components/GuestRoute.tsx → apps/frontend/admin/src/hooks/useAuth.ts
- `ProtectedRoute()` --calls--> `useAuth()`  [INFERRED]
  apps/frontend/admin/src/components/ProtectedRoute.tsx → apps/frontend/admin/src/hooks/useAuth.ts
- `DashboardPage()` --calls--> `useAuth()`  [INFERRED]
  apps/frontend/admin/src/pages/DashboardPage.tsx → apps/frontend/admin/src/hooks/useAuth.ts

## Communities

### Community 0 - "Community 0"
Cohesion: 0.01
Nodes (54): AccountGroupBy, AccountQuery, AccountSelect, BlogPostGroupBy, BlogPostQuery, BlogPostSelect, CertificateGroupBy, CertificateQuery (+46 more)

### Community 1 - "Community 1"
Cohesion: 0.0
Nodes (9): BlogPostMutation, CertificateMutation, EducationMutation, ExperienceMutation, ProjectMutation, SessionMutation, SkillMutation, TestimonialMutation (+1 more)

### Community 2 - "Community 2"
Cohesion: 0.01
Nodes (89): AccountClient, accountOption, BlogCommentClient, blogcommentOption, BlogPostClient, blogpostOption, CertificateClient, certificateOption (+81 more)

### Community 3 - "Community 3"
Cohesion: 0.01
Nodes (52): AccountDelete, AccountDeleteOne, BlogCommentDelete, BlogCommentDeleteOne, BlogPostDelete, BlogPostDeleteOne, CertificateDelete, CertificateDeleteOne (+44 more)

### Community 4 - "Community 4"
Cohesion: 0.01
Nodes (39): PasswordNotNil(), newAuthService(), DisableTOTPCommand, emailVerificationLimitState, EnableTOTPCommand, EntCredentialRepository, EntEmailVerificationLimiter, EntSessionRepository (+31 more)

### Community 5 - "Community 5"
Cohesion: 0.01
Nodes (85): Account, AccountEdges, Accounts, BlogComment, BlogCommentEdges, BlogComments, BlogPost, BlogPostEdges (+77 more)

### Community 7 - "Community 7"
Cohesion: 0.01
Nodes (9): ByAuthorField(), ByComments(), ByCommentsCount(), newAuthorStep(), newCommentsStep(), OrderOption, Status, HasAuthorWith() (+1 more)

### Community 8 - "Community 8"
Cohesion: 0.01
Nodes (21): OrderOption, ByAccounts(), ByAccountsCount(), ByBlogPosts(), ByBlogPostsCount(), ByReviewedComments(), ByReviewedCommentsCount(), BySessions() (+13 more)

### Community 9 - "Community 9"
Cohesion: 0.01
Nodes (15): ByBlogPostField(), ByChildren(), ByChildrenCount(), ByParentField(), ByReviewedByField(), newBlogPostStep(), newChildrenStep(), newParentStep() (+7 more)

### Community 11 - "Community 11"
Cohesion: 0.02
Nodes (2): BlogCommentMutation, MessageMutation

### Community 12 - "Community 12"
Cohesion: 0.03
Nodes (62): selector, Account, BlogComment, BlogPost, Certificate, CVDownloadLog, Education, Experience (+54 more)

### Community 13 - "Community 13"
Cohesion: 0.02
Nodes (2): UserUpdate, UserUpdateOne

### Community 14 - "Community 14"
Cohesion: 0.02
Nodes (2): AccountMutation, TwoFactorMutation

### Community 15 - "Community 15"
Cohesion: 0.02
Nodes (2): ProfileContentMutation, ResumeAssetMutation

### Community 16 - "Community 16"
Cohesion: 0.02
Nodes (2): ProjectUpdate, ProjectUpdateOne

### Community 17 - "Community 17"
Cohesion: 0.03
Nodes (8): TestimonialCreate, TestimonialCreateBulk, TestimonialUpdate, TestimonialUpdateOne, OrderOption, Relation, Status, RelationValidator()

### Community 22 - "Community 22"
Cohesion: 0.03
Nodes (2): BlogPostUpdate, BlogPostUpdateOne

### Community 25 - "Community 25"
Cohesion: 0.03
Nodes (75): appEnvironment(), appStoragePath(), main(), newDatabaseClient(), newHandler(), optionalDatabaseClient(), stringEnvFallback(), TestRootRouteReturnsOK() (+67 more)

### Community 26 - "Community 26"
Cohesion: 0.03
Nodes (2): CVDownloadLogMutation, VisitorLogMutation

### Community 28 - "Community 28"
Cohesion: 0.03
Nodes (2): BlogCommentUpdate, BlogCommentUpdateOne

### Community 30 - "Community 30"
Cohesion: 0.03
Nodes (6): AccountCreate, AccountCreateBulk, CertificateCreate, CertificateCreateBulk, SessionCreate, SessionCreateBulk

### Community 31 - "Community 31"
Cohesion: 0.03
Nodes (52): argon2Params, Argon2PasswordVerifier, sameToken(), validateNewPassword(), authEmailTemplateData, AuthenticatedUser, BetterAuthScryptPasswordVerifier, ChangePasswordCommand (+44 more)

### Community 32 - "Community 32"
Cohesion: 0.04
Nodes (2): AccountUpdate, AccountUpdateOne

### Community 33 - "Community 33"
Cohesion: 0.04
Nodes (2): ProfileContentUpdate, ProfileContentUpdateOne

### Community 34 - "Community 34"
Cohesion: 0.03
Nodes (6): ProjectCreate, ProjectCreateBulk, Accent, OrderOption, AccentValidator(), Status

### Community 37 - "Community 37"
Cohesion: 0.05
Nodes (2): SessionUpdate, SessionUpdateOne

### Community 38 - "Community 38"
Cohesion: 0.06
Nodes (2): MessageUpdate, MessageUpdateOne

### Community 39 - "Community 39"
Cohesion: 0.06
Nodes (2): ExperienceUpdate, ExperienceUpdateOne

### Community 40 - "Community 40"
Cohesion: 0.06
Nodes (33): APIEndpointDetail, APIIndexResponse, authErrorMessage(), clientIP(), emailVerificationErrorMessage(), firstQueryValue(), passwordErrorMessage(), sessionTokenFromCookie() (+25 more)

### Community 42 - "Community 42"
Cohesion: 0.05
Nodes (35): NewService(), TestChangePasswordRequiresCurrentPassword(), TestChangePasswordUpdatesPassword(), TestForgotPasswordReturnsDeliveryFailure(), TestForgotPasswordSendsResetEmailViaSMTP(), TestLoginUsesDefaultTTLWithoutRememberMe(), TestLoginUsesRememberMeTTL(), TestLogoutRevokesCurrentSessionToken() (+27 more)

### Community 43 - "Community 43"
Cohesion: 0.06
Nodes (2): EducationUpdate, EducationUpdateOne

### Community 44 - "Community 44"
Cohesion: 0.06
Nodes (2): ResumeAssetUpdate, ResumeAssetUpdateOne

### Community 45 - "Community 45"
Cohesion: 0.04
Nodes (28): AccountFunc, BlogCommentFunc, BlogPostFunc, CertificateFunc, Chain, Condition, CVDownloadLogFunc, EducationFunc (+20 more)

### Community 46 - "Community 46"
Cohesion: 0.08
Nodes (3): UserGroupBy, UserQuery, UserSelect

### Community 47 - "Community 47"
Cohesion: 0.06
Nodes (2): VisitorLogUpdate, VisitorLogUpdateOne

### Community 48 - "Community 48"
Cohesion: 0.08
Nodes (3): BlogCommentGroupBy, BlogCommentQuery, BlogCommentSelect

### Community 49 - "Community 49"
Cohesion: 0.05
Nodes (31): AggregateFunc, clientCtxKey, ConstraintError, As(), Asc(), checkColumn(), Desc(), IsNotLoaded() (+23 more)

### Community 50 - "Community 50"
Cohesion: 0.06
Nodes (2): UserCreate, UserCreateBulk

### Community 51 - "Community 51"
Cohesion: 0.07
Nodes (2): BlogPostCreate, BlogPostCreateBulk

### Community 52 - "Community 52"
Cohesion: 0.08
Nodes (2): SkillUpdate, SkillUpdateOne

### Community 53 - "Community 53"
Cohesion: 0.05
Nodes (2): renderPage(), CmsApp()

### Community 54 - "Community 54"
Cohesion: 0.08
Nodes (2): BlogCommentCreate, BlogCommentCreateBulk

### Community 55 - "Community 55"
Cohesion: 0.09
Nodes (2): ProfileContentCreate, ProfileContentCreateBulk

### Community 56 - "Community 56"
Cohesion: 0.11
Nodes (2): MessageCreate, MessageCreateBulk

### Community 57 - "Community 57"
Cohesion: 0.12
Nodes (2): VisitorLogCreate, VisitorLogCreateBulk

### Community 58 - "Community 58"
Cohesion: 0.12
Nodes (2): ResumeAssetCreate, ResumeAssetCreateBulk

### Community 59 - "Community 59"
Cohesion: 0.11
Nodes (2): ExperienceCreate, ExperienceCreateBulk

### Community 60 - "Community 60"
Cohesion: 0.13
Nodes (2): SkillCreate, SkillCreateBulk

### Community 61 - "Community 61"
Cohesion: 0.15
Nodes (2): CVDownloadLogCreate, CVDownloadLogCreateBulk

### Community 62 - "Community 62"
Cohesion: 0.11
Nodes (18): Account, BlogComment, BlogPost, Certificate, CVDownloadLog, Education, Experience, Message (+10 more)

### Community 63 - "Community 63"
Cohesion: 0.12
Nodes (3): ByUserField(), newUserStep(), OrderOption

### Community 64 - "Community 64"
Cohesion: 0.12
Nodes (1): OrderOption

### Community 65 - "Community 65"
Cohesion: 0.12
Nodes (2): OrderOption, Status

### Community 66 - "Community 66"
Cohesion: 0.15
Nodes (6): GuestRoute(), ProtectedRoute(), useAuth(), formatAbbreviatedName(), TopBar(), DashboardPage()

### Community 67 - "Community 67"
Cohesion: 0.15
Nodes (3): OrderOption, ByUserField(), newUserStep()

### Community 68 - "Community 68"
Cohesion: 0.15
Nodes (1): OrderOption

### Community 69 - "Community 69"
Cohesion: 0.29
Nodes (6): handleDisableTwoFactor(), handleEnableTwoFactor(), handleRegenerateBackupCodes(), handleSetupTwoFactor(), resetTwoFactorFeedback(), twoFactorErrorMessage()

### Community 70 - "Community 70"
Cohesion: 0.17
Nodes (1): OrderOption

### Community 71 - "Community 71"
Cohesion: 0.17
Nodes (1): OrderOption

### Community 73 - "Community 73"
Cohesion: 0.18
Nodes (1): OrderOption

### Community 74 - "Community 74"
Cohesion: 0.18
Nodes (1): OrderOption

### Community 76 - "Community 76"
Cohesion: 0.24
Nodes (7): AuthClaims(), NewJWTAuthentication(), TestJWTAuthenticationIgnoresBearerToken(), TestJWTAuthenticationUsesSessionCookie(), Authenticator, contextKey, stubAuthenticator

### Community 77 - "Community 77"
Cohesion: 0.2
Nodes (1): OrderOption

### Community 78 - "Community 78"
Cohesion: 0.29
Nodes (9): migrateSchema(), NewClient(), newOptions(), Open(), WithMigrateOptions(), WithOptions(), Option, options (+1 more)

### Community 80 - "Community 80"
Cohesion: 0.25
Nodes (3): OrderOption, ByUserField(), newUserStep()

### Community 81 - "Community 81"
Cohesion: 0.22
Nodes (1): OrderOption

### Community 83 - "Community 83"
Cohesion: 0.25
Nodes (1): OrderOption

### Community 85 - "Community 85"
Cohesion: 0.33
Nodes (1): RouteErrorBoundary

## Knowledge Gaps
- **159 isolated node(s):** `AuthConfig`, `AuthCookieConfig`, `RootResponse`, `APIIndexResponse`, `APIEndpointDetail` (+154 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **Thin community `Community 11`** (162 nodes): `BlogCommentMutation`, `.AddChildIDs()`, `.AddedEdges()`, `.AddedField()`, `.AddedFields()`, `.AddedIDs()`, `.AddField()`, `.BlogPostCleared()`, `.BlogPostID()`, `.BlogPostIDs()`, `.Body()`, `.ChildrenCleared()`, `.ChildrenIDs()`, `.ClearBlogPost()`, `.ClearChildren()`, `.ClearedEdges()`, `.ClearEdge()`, `.ClearField()`, `.ClearParent()`, `.ClearParentID()`, `.ClearReviewedAt()`, `.ClearReviewedBy()`, `.ClearReviewedByUserID()`, `.Client()`, `.CreatedAt()`, `.DisplayName()`, `.EdgeCleared()`, `.Email()`, `.Field()`, `.FieldCleared()`, `.Fields()`, `.Fingerprint()`, `.ID()`, `.IPHash()`, `.OldBody()`, `.OldCreatedAt()`, `.OldEmail()`, `.OldIPHash()`, `.OldReviewedAt()`, `.OldStatus()`, `.OldUpdatedAt()`, `.Op()`, `.ParentCleared()`, `.ParentID()`, `.ParentIDCleared()`, `.ParentIDs()`, `.RemoveChildIDs()`, `.RemovedChildrenIDs()`, `.RemovedEdges()`, `.RemovedIDs()`, `.ResetBlogPost()`, `.ResetBlogPostID()`, `.ResetBody()`, `.ResetChildren()`, `.ResetCreatedAt()`, `.ResetDisplayName()`, `.ResetEdge()`, `.ResetEmail()`, `.ResetField()`, `.ResetFingerprint()`, `.ResetIPHash()`, `.ResetParent()`, `.ResetParentID()`, `.ResetReviewedAt()`, `.ResetReviewedBy()`, `.ResetReviewedByUserID()`, `.ResetStatus()`, `.ResetUpdatedAt()`, `.ReviewedAt()`, `.ReviewedAtCleared()`, `.ReviewedByCleared()`, `.ReviewedByID()`, `.ReviewedByIDs()`, `.ReviewedByUserID()`, `.ReviewedByUserIDCleared()`, `.SetBlogPostID()`, `.SetBody()`, `.SetCreatedAt()`, `.SetDisplayName()`, `.SetEmail()`, `.SetField()`, `.SetFingerprint()`, `.SetID()`, `.SetIPHash()`, `.SetOp()`, `.SetParentID()`, `.SetReviewedAt()`, `.SetReviewedByID()`, `.SetReviewedByUserID()`, `.SetStatus()`, `.SetUpdatedAt()`, `.Status()`, `.Tx()`, `.Type()`, `.UpdatedAt()`, `.Where()`, `.WhereP()`, `MessageMutation`, `.AddedEdges()`, `.AddedField()`, `.AddedFields()`, `.AddedIDs()`, `.AddField()`, `.ArchivedAt()`, `.ArchivedAtCleared()`, `.Body()`, `.ClearArchivedAt()`, `.ClearedEdges()`, `.ClearEdge()`, `.ClearField()`, `.ClearReadAt()`, `.Client()`, `.CreatedAt()`, `.EdgeCleared()`, `.Field()`, `.FieldCleared()`, `.Fields()`, `.ID()`, `.OldCreatedAt()`, `.OldStatus()`, `.OldUpdatedAt()`, `.Op()`, `.ReadAt()`, `.ReadAtCleared()`, `.RemovedEdges()`, `.RemovedIDs()`, `.ResetArchivedAt()`, `.ResetBody()`, `.ResetCreatedAt()`, `.ResetEdge()`, `.ResetField()`, `.ResetReadAt()`, `.ResetSenderEmail()`, `.ResetSenderName()`, `.ResetStatus()`, `.ResetSubject()`, `.ResetUpdatedAt()`, `.SenderEmail()`, `.SenderName()`, `.SetArchivedAt()`, `.SetBody()`, `.SetCreatedAt()`, `.SetField()`, `.SetID()`, `.SetOp()`, `.SetReadAt()`, `.SetSenderEmail()`, `.SetSenderName()`, `.SetStatus()`, `.SetSubject()`, `.SetUpdatedAt()`, `.Status()`, `.Subject()`, `.Tx()`, `.Type()`, `.UpdatedAt()`, `.Where()`, `.WhereP()`, `.SetStatus()`, `.Email()`, `.ResetEmail()`, `.SetEmail()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 13`** (143 nodes): `user_update.go`, `UserUpdate`, `.AddAccountIDs()`, `.AddAccounts()`, `.AddBlogPostIDs()`, `.AddBlogPosts()`, `.AddReviewedCommentIDs()`, `.AddReviewedComments()`, `.AddSessionIDs()`, `.AddSessions()`, `.AddTwoFactorIDs()`, `.AddTwoFactors()`, `.check()`, `.ClearAccounts()`, `.ClearBanExpires()`, `.ClearBanned()`, `.ClearBanReason()`, `.ClearBlogPosts()`, `.ClearDisplayUsername()`, `.ClearImage()`, `.ClearIsAnonymous()`, `.ClearReviewedComments()`, `.ClearRole()`, `.ClearSessions()`, `.ClearTwoFactorEnabled()`, `.ClearTwoFactors()`, `.ClearUsername()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.RemoveAccountIDs()`, `.RemoveAccounts()`, `.RemoveBlogPostIDs()`, `.RemoveBlogPosts()`, `.RemoveReviewedCommentIDs()`, `.RemoveReviewedComments()`, `.RemoveSessionIDs()`, `.RemoveSessions()`, `.RemoveTwoFactorIDs()`, `.RemoveTwoFactors()`, `.SaveX()`, `.SetBanExpires()`, `.SetBanned()`, `.SetBanReason()`, `.SetCreatedAt()`, `.SetDisplayUsername()`, `.SetEmail()`, `.SetEmailVerified()`, `.SetImage()`, `.SetIsAnonymous()`, `.SetName()`, `.SetNillableBanExpires()`, `.SetNillableBanned()`, `.SetNillableBanReason()`, `.SetNillableCreatedAt()`, `.SetNillableDisplayUsername()`, `.SetNillableEmail()`, `.SetNillableEmailVerified()`, `.SetNillableImage()`, `.SetNillableIsAnonymous()`, `.SetNillableName()`, `.SetNillableRole()`, `.SetNillableTwoFactorEnabled()`, `.SetNillableUpdatedAt()`, `.SetNillableUsername()`, `.SetRole()`, `.SetTwoFactorEnabled()`, `.SetUpdatedAt()`, `.SetUsername()`, `.sqlSave()`, `.Where()`, `UserUpdateOne`, `.AddAccountIDs()`, `.AddAccounts()`, `.AddBlogPostIDs()`, `.AddBlogPosts()`, `.AddReviewedCommentIDs()`, `.AddReviewedComments()`, `.AddSessionIDs()`, `.AddSessions()`, `.AddTwoFactorIDs()`, `.AddTwoFactors()`, `.check()`, `.ClearAccounts()`, `.ClearBanExpires()`, `.ClearBanned()`, `.ClearBanReason()`, `.ClearBlogPosts()`, `.ClearDisplayUsername()`, `.ClearImage()`, `.ClearIsAnonymous()`, `.ClearReviewedComments()`, `.ClearRole()`, `.ClearSessions()`, `.ClearTwoFactorEnabled()`, `.ClearTwoFactors()`, `.ClearUsername()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.RemoveAccountIDs()`, `.RemoveAccounts()`, `.RemoveBlogPostIDs()`, `.RemoveBlogPosts()`, `.RemoveReviewedCommentIDs()`, `.RemoveReviewedComments()`, `.RemoveSessionIDs()`, `.RemoveSessions()`, `.RemoveTwoFactorIDs()`, `.RemoveTwoFactors()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetBanExpires()`, `.SetBanned()`, `.SetBanReason()`, `.SetCreatedAt()`, `.SetDisplayUsername()`, `.SetEmail()`, `.SetEmailVerified()`, `.SetImage()`, `.SetIsAnonymous()`, `.SetName()`, `.SetNillableBanExpires()`, `.SetNillableBanned()`, `.SetNillableBanReason()`, `.SetNillableCreatedAt()`, `.SetNillableDisplayUsername()`, `.SetNillableEmail()`, `.SetNillableEmailVerified()`, `.SetNillableImage()`, `.SetNillableIsAnonymous()`, `.SetNillableName()`, `.SetNillableRole()`, `.SetNillableTwoFactorEnabled()`, `.SetNillableUpdatedAt()`, `.SetNillableUsername()`, `.SetRole()`, `.SetTwoFactorEnabled()`, `.SetUpdatedAt()`, `.SetUsername()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 14`** (141 nodes): `AccountMutation`, `.AccessToken()`, `.AccessTokenCleared()`, `.AccessTokenExpiresAt()`, `.AccessTokenExpiresAtCleared()`, `.AccountID()`, `.AddedEdges()`, `.AddedField()`, `.AddedFields()`, `.AddedIDs()`, `.AddField()`, `.ClearAccessToken()`, `.ClearAccessTokenExpiresAt()`, `.ClearedEdges()`, `.ClearEdge()`, `.ClearField()`, `.ClearIDToken()`, `.ClearPassword()`, `.ClearRefreshToken()`, `.ClearRefreshTokenExpiresAt()`, `.ClearScope()`, `.ClearUser()`, `.Client()`, `.CreatedAt()`, `.EdgeCleared()`, `.Field()`, `.FieldCleared()`, `.Fields()`, `.ID()`, `.IDToken()`, `.IDTokenCleared()`, `.OldAccessToken()`, `.OldAccessTokenExpiresAt()`, `.OldAccountID()`, `.OldCreatedAt()`, `.OldField()`, `.OldIDToken()`, `.OldPassword()`, `.OldProviderID()`, `.OldRefreshToken()`, `.OldRefreshTokenExpiresAt()`, `.OldScope()`, `.OldUpdatedAt()`, `.OldUserID()`, `.Op()`, `.Password()`, `.PasswordCleared()`, `.ProviderID()`, `.RefreshToken()`, `.RefreshTokenCleared()`, `.RefreshTokenExpiresAt()`, `.RefreshTokenExpiresAtCleared()`, `.RemovedEdges()`, `.RemovedIDs()`, `.ResetAccessToken()`, `.ResetAccessTokenExpiresAt()`, `.ResetAccountID()`, `.ResetCreatedAt()`, `.ResetEdge()`, `.ResetField()`, `.ResetIDToken()`, `.ResetPassword()`, `.ResetProviderID()`, `.ResetRefreshToken()`, `.ResetRefreshTokenExpiresAt()`, `.ResetScope()`, `.ResetUpdatedAt()`, `.ResetUser()`, `.ResetUserID()`, `.Scope()`, `.ScopeCleared()`, `.SetAccessToken()`, `.SetAccessTokenExpiresAt()`, `.SetAccountID()`, `.SetCreatedAt()`, `.SetField()`, `.SetID()`, `.SetIDToken()`, `.SetOp()`, `.SetPassword()`, `.SetProviderID()`, `.SetRefreshToken()`, `.SetRefreshTokenExpiresAt()`, `.SetScope()`, `.SetUpdatedAt()`, `.SetUserID()`, `.Tx()`, `.Type()`, `.UpdatedAt()`, `.UserCleared()`, `.UserID()`, `.UserIDs()`, `.Where()`, `.WhereP()`, `.ClearEdge()`, `.ResetEdge()`, `TwoFactorMutation`, `.AddedEdges()`, `.AddedField()`, `.AddedFields()`, `.AddedIDs()`, `.AddField()`, `.BackupCodes()`, `.ClearedEdges()`, `.ClearedFields()`, `.ClearEdge()`, `.ClearField()`, `.ClearUser()`, `.Client()`, `.EdgeCleared()`, `.Field()`, `.FieldCleared()`, `.Fields()`, `.ID()`, `.OldBackupCodes()`, `.OldField()`, `.OldSecret()`, `.OldUserID()`, `.Op()`, `.RemovedEdges()`, `.RemovedIDs()`, `.ResetBackupCodes()`, `.ResetEdge()`, `.ResetField()`, `.ResetSecret()`, `.ResetUser()`, `.ResetUserID()`, `.Secret()`, `.SetBackupCodes()`, `.SetField()`, `.SetID()`, `.SetOp()`, `.SetSecret()`, `.SetUserID()`, `.Tx()`, `.Type()`, `.UserCleared()`, `.UserID()`, `.UserIDs()`, `.Where()`, `.WhereP()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 15`** (138 nodes): `ProfileContentMutation`, `.About()`, `.AddedEdges()`, `.AddedField()`, `.AddedFields()`, `.AddedIDs()`, `.AddField()`, `.AppendedFocus()`, `.AppendedStats()`, `.AppendFocus()`, `.AppendStats()`, `.Availability()`, `.ClearedEdges()`, `.ClearEdge()`, `.ClearField()`, `.ClearProfilePhotoURL()`, `.Client()`, `.CreatedAt()`, `.EdgeCleared()`, `.Email()`, `.Field()`, `.FieldCleared()`, `.Fields()`, `.Focus()`, `.FullName()`, `.Headline()`, `.ID()`, `.OldCreatedAt()`, `.OldEmail()`, `.OldStorageKey()`, `.OldUpdatedAt()`, `.Op()`, `.Phone()`, `.PrimaryCta()`, `.ProfilePhotoURL()`, `.ProfilePhotoURLCleared()`, `.RemovedEdges()`, `.RemovedIDs()`, `.ResetAbout()`, `.ResetAvailability()`, `.ResetCreatedAt()`, `.ResetEdge()`, `.ResetEmail()`, `.ResetField()`, `.ResetFocus()`, `.ResetFullName()`, `.ResetHeadline()`, `.ResetPhone()`, `.ResetPrimaryCta()`, `.ResetProfilePhotoURL()`, `.ResetShortIntro()`, `.ResetSocialLinks()`, `.ResetStats()`, `.ResetStorageKey()`, `.ResetUpdatedAt()`, `.SetAbout()`, `.SetAvailability()`, `.SetCreatedAt()`, `.SetEmail()`, `.SetField()`, `.SetFocus()`, `.SetFullName()`, `.SetHeadline()`, `.SetID()`, `.SetOp()`, `.SetPhone()`, `.SetPrimaryCta()`, `.SetProfilePhotoURL()`, `.SetShortIntro()`, `.SetSocialLinks()`, `.SetStats()`, `.SetStorageKey()`, `.SetUpdatedAt()`, `.ShortIntro()`, `.SocialLinks()`, `.Stats()`, `.StorageKey()`, `.Tx()`, `.Type()`, `.UpdatedAt()`, `.Where()`, `.WhereP()`, `ResumeAssetMutation`, `.AddedEdges()`, `.AddedField()`, `.AddedFields()`, `.AddedFileSizeBytes()`, `.AddedIDs()`, `.AddField()`, `.AddFileSizeBytes()`, `.ClearedEdges()`, `.ClearEdge()`, `.ClearField()`, `.ClearFileName()`, `.ClearFileSizeBytes()`, `.Client()`, `.CreatedAt()`, `.DownloadURL()`, `.EdgeCleared()`, `.Field()`, `.FieldCleared()`, `.Fields()`, `.FileName()`, `.FileNameCleared()`, `.FileSizeBytes()`, `.FileSizeBytesCleared()`, `.ID()`, `.MimeType()`, `.OldCreatedAt()`, `.OldUpdatedAt()`, `.Op()`, `.RemovedEdges()`, `.RemovedIDs()`, `.ResetCreatedAt()`, `.ResetDownloadURL()`, `.ResetEdge()`, `.ResetField()`, `.ResetFileName()`, `.ResetFileSizeBytes()`, `.ResetMimeType()`, `.ResetStorageKey()`, `.ResetUpdatedAt()`, `.SetCreatedAt()`, `.SetDownloadURL()`, `.SetField()`, `.SetFileName()`, `.SetFileSizeBytes()`, `.SetID()`, `.SetMimeType()`, `.SetOp()`, `.SetStorageKey()`, `.SetUpdatedAt()`, `.StorageKey()`, `.Tx()`, `.Type()`, `.UpdatedAt()`, `.Where()`, `.WhereP()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 16`** (137 nodes): `project_update.go`, `ProjectUpdate`, `.AddSortOrder()`, `.AppendGallery()`, `.AppendImpactBullets()`, `.AppendProcess()`, `.AppendTags()`, `.AppendTechStack()`, `.ClearGithubURL()`, `.ClearImpactBullets()`, `.ClearProcess()`, `.ClearProjectURL()`, `.ClearSortOrder()`, `.ClearTags()`, `.ClearTechStack()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetAccent()`, `.SetCategory()`, `.SetChallenge()`, `.SetClientOrCompany()`, `.SetCreatedAt()`, `.SetDuration()`, `.SetFeatured()`, `.SetGallery()`, `.SetGithubURL()`, `.SetImpactBullets()`, `.SetImpactSummary()`, `.SetMetrics()`, `.SetNillableAccent()`, `.SetNillableCategory()`, `.SetNillableChallenge()`, `.SetNillableClientOrCompany()`, `.SetNillableCreatedAt()`, `.SetNillableDuration()`, `.SetNillableFeatured()`, `.SetNillableGithubURL()`, `.SetNillableImpactSummary()`, `.SetNillableOutcome()`, `.SetNillableProjectURL()`, `.SetNillableRole()`, `.SetNillableSlug()`, `.SetNillableSortOrder()`, `.SetNillableStatus()`, `.SetNillableSummary()`, `.SetNillableThumbnailPlaceholder()`, `.SetNillableTitle()`, `.SetNillableYear()`, `.SetOutcome()`, `.SetProcess()`, `.SetProjectURL()`, `.SetRole()`, `.SetSlug()`, `.SetSortOrder()`, `.SetStatus()`, `.SetSummary()`, `.SetTags()`, `.SetTechStack()`, `.SetThumbnailPlaceholder()`, `.SetTitle()`, `.SetUpdatedAt()`, `.SetYear()`, `.sqlSave()`, `.Where()`, `ProjectUpdateOne`, `.AddSortOrder()`, `.AppendGallery()`, `.AppendImpactBullets()`, `.AppendProcess()`, `.AppendTags()`, `.AppendTechStack()`, `.check()`, `.ClearGithubURL()`, `.ClearImpactBullets()`, `.ClearProcess()`, `.ClearProjectURL()`, `.ClearSortOrder()`, `.ClearTags()`, `.ClearTechStack()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetAccent()`, `.SetCategory()`, `.SetChallenge()`, `.SetClientOrCompany()`, `.SetCreatedAt()`, `.SetDuration()`, `.SetFeatured()`, `.SetGallery()`, `.SetGithubURL()`, `.SetImpactBullets()`, `.SetImpactSummary()`, `.SetMetrics()`, `.SetNillableAccent()`, `.SetNillableCategory()`, `.SetNillableChallenge()`, `.SetNillableClientOrCompany()`, `.SetNillableCreatedAt()`, `.SetNillableDuration()`, `.SetNillableFeatured()`, `.SetNillableGithubURL()`, `.SetNillableImpactSummary()`, `.SetNillableOutcome()`, `.SetNillableProjectURL()`, `.SetNillableRole()`, `.SetNillableSlug()`, `.SetNillableSortOrder()`, `.SetNillableStatus()`, `.SetNillableSummary()`, `.SetNillableThumbnailPlaceholder()`, `.SetNillableTitle()`, `.SetNillableYear()`, `.SetOutcome()`, `.SetProcess()`, `.SetProjectURL()`, `.SetRole()`, `.SetSlug()`, `.SetSortOrder()`, `.SetStatus()`, `.SetSummary()`, `.SetTags()`, `.SetTechStack()`, `.SetThumbnailPlaceholder()`, `.SetTitle()`, `.SetUpdatedAt()`, `.SetYear()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 22`** (114 nodes): `blogpost_update.go`, `BlogPostUpdate`, `.AddCommentIDs()`, `.AddComments()`, `.AppendTags()`, `.check()`, `.ClearAuthor()`, `.ClearComments()`, `.ClearCoverImagePlaceholder()`, `.ClearPublishDate()`, `.ClearSeoDescription()`, `.ClearSeoTitle()`, `.ClearTags()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.RemoveCommentIDs()`, `.RemoveComments()`, `.Save()`, `.SaveX()`, `.SetAuthor()`, `.SetAuthorID()`, `.SetAuthorName()`, `.SetAuthorUserID()`, `.SetCategory()`, `.SetContent()`, `.SetCoverImagePlaceholder()`, `.SetCreatedAt()`, `.SetExcerpt()`, `.SetFeatured()`, `.SetNillableAuthorName()`, `.SetNillableAuthorUserID()`, `.SetNillableCategory()`, `.SetNillableContent()`, `.SetNillableCoverImagePlaceholder()`, `.SetNillableCreatedAt()`, `.SetNillableExcerpt()`, `.SetNillableFeatured()`, `.SetNillablePublishDate()`, `.SetNillableReadingTime()`, `.SetNillableSeoDescription()`, `.SetNillableSeoTitle()`, `.SetNillableSlug()`, `.SetNillableStatus()`, `.SetNillableTitle()`, `.SetPublishDate()`, `.SetReadingTime()`, `.SetSeoDescription()`, `.SetSeoTitle()`, `.SetSlug()`, `.SetStatus()`, `.SetTags()`, `.SetTitle()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`, `BlogPostUpdateOne`, `.AddCommentIDs()`, `.AddComments()`, `.AppendTags()`, `.check()`, `.ClearAuthor()`, `.ClearComments()`, `.ClearCoverImagePlaceholder()`, `.ClearPublishDate()`, `.ClearSeoDescription()`, `.ClearSeoTitle()`, `.ClearTags()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.RemoveCommentIDs()`, `.RemoveComments()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetAuthor()`, `.SetAuthorID()`, `.SetAuthorName()`, `.SetAuthorUserID()`, `.SetCategory()`, `.SetContent()`, `.SetCoverImagePlaceholder()`, `.SetCreatedAt()`, `.SetExcerpt()`, `.SetFeatured()`, `.SetNillableAuthorName()`, `.SetNillableAuthorUserID()`, `.SetNillableCategory()`, `.SetNillableContent()`, `.SetNillableCoverImagePlaceholder()`, `.SetNillableCreatedAt()`, `.SetNillableExcerpt()`, `.SetNillableFeatured()`, `.SetNillablePublishDate()`, `.SetNillableReadingTime()`, `.SetNillableSeoDescription()`, `.SetNillableSeoTitle()`, `.SetNillableSlug()`, `.SetNillableStatus()`, `.SetNillableTitle()`, `.SetPublishDate()`, `.SetReadingTime()`, `.SetSeoDescription()`, `.SetSeoTitle()`, `.SetSlug()`, `.SetStatus()`, `.SetTags()`, `.SetTitle()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 26`** (102 nodes): `CVDownloadLogMutation`, `.AddedEdges()`, `.AddedField()`, `.AddedFields()`, `.AddedIDs()`, `.AddField()`, `.ClearedEdges()`, `.ClearedFields()`, `.ClearEdge()`, `.ClearField()`, `.Client()`, `.DownloadedAt()`, `.EdgeCleared()`, `.Field()`, `.FieldCleared()`, `.Fields()`, `.ID()`, `.IPHash()`, `.OldDownloadedAt()`, `.OldField()`, `.OldIPHash()`, `.OldReferrer()`, `.OldUserAgent()`, `.Op()`, `.Referrer()`, `.RemovedEdges()`, `.RemovedIDs()`, `.ResetDownloadedAt()`, `.ResetEdge()`, `.ResetField()`, `.ResetIPHash()`, `.ResetReferrer()`, `.ResetUserAgent()`, `.SetDownloadedAt()`, `.SetField()`, `.SetID()`, `.SetIPHash()`, `.SetOp()`, `.SetReferrer()`, `.SetUserAgent()`, `.Tx()`, `.Type()`, `.UserAgent()`, `.Where()`, `.WhereP()`, `VisitorLogMutation`, `.AddedEdges()`, `.AddedField()`, `.AddedFields()`, `.AddedIDs()`, `.AddField()`, `.ClearedEdges()`, `.ClearedFields()`, `.ClearEdge()`, `.ClearField()`, `.EdgeCleared()`, `.Field()`, `.Fields()`, `.IPHash()`, `.IsUniqueDailyVisitor()`, `.OldField()`, `.OldIPHash()`, `.OldIsUniqueDailyVisitor()`, `.OldPath()`, `.OldReferrer()`, `.OldReferrerSource()`, `.OldUserAgent()`, `.OldVisitedAt()`, `.OldVisitorID()`, `.Op()`, `.Path()`, `.Referrer()`, `.ReferrerSource()`, `.RemovedEdges()`, `.RemovedIDs()`, `.ResetEdge()`, `.ResetField()`, `.ResetIPHash()`, `.ResetIsUniqueDailyVisitor()`, `.ResetPath()`, `.ResetReferrer()`, `.ResetReferrerSource()`, `.ResetUserAgent()`, `.ResetVisitedAt()`, `.ResetVisitorID()`, `.SetField()`, `.SetID()`, `.SetIPHash()`, `.SetIsUniqueDailyVisitor()`, `.SetOp()`, `.SetPath()`, `.SetReferrer()`, `.SetReferrerSource()`, `.SetUserAgent()`, `.SetVisitedAt()`, `.SetVisitorID()`, `.Tx()`, `.Type()`, `.UserAgent()`, `.VisitedAt()`, `.VisitorID()`, `.WhereP()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 28`** (100 nodes): `blogcomment_update.go`, `BlogCommentUpdate`, `.AddChildIDs()`, `.AddChildren()`, `.check()`, `.ClearBlogPost()`, `.ClearChildren()`, `.ClearParent()`, `.ClearParentID()`, `.ClearReviewedAt()`, `.ClearReviewedBy()`, `.ClearReviewedByUserID()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.RemoveChildIDs()`, `.RemoveChildren()`, `.Save()`, `.SaveX()`, `.SetBlogPost()`, `.SetBlogPostID()`, `.SetBody()`, `.SetCreatedAt()`, `.SetDisplayName()`, `.SetEmail()`, `.SetFingerprint()`, `.SetIPHash()`, `.SetNillableBlogPostID()`, `.SetNillableBody()`, `.SetNillableCreatedAt()`, `.SetNillableDisplayName()`, `.SetNillableEmail()`, `.SetNillableFingerprint()`, `.SetNillableIPHash()`, `.SetNillableParentID()`, `.SetNillableReviewedAt()`, `.SetNillableReviewedByID()`, `.SetNillableReviewedByUserID()`, `.SetNillableStatus()`, `.SetParent()`, `.SetParentID()`, `.SetReviewedAt()`, `.SetReviewedBy()`, `.SetReviewedByID()`, `.SetReviewedByUserID()`, `.SetStatus()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`, `BlogCommentUpdateOne`, `.AddChildIDs()`, `.AddChildren()`, `.check()`, `.ClearBlogPost()`, `.ClearChildren()`, `.ClearParent()`, `.ClearParentID()`, `.ClearReviewedAt()`, `.ClearReviewedBy()`, `.ClearReviewedByUserID()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.RemoveChildIDs()`, `.RemoveChildren()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetBlogPost()`, `.SetBlogPostID()`, `.SetBody()`, `.SetCreatedAt()`, `.SetDisplayName()`, `.SetEmail()`, `.SetFingerprint()`, `.SetIPHash()`, `.SetNillableBlogPostID()`, `.SetNillableBody()`, `.SetNillableCreatedAt()`, `.SetNillableDisplayName()`, `.SetNillableEmail()`, `.SetNillableFingerprint()`, `.SetNillableIPHash()`, `.SetNillableParentID()`, `.SetNillableReviewedAt()`, `.SetNillableReviewedByID()`, `.SetNillableReviewedByUserID()`, `.SetNillableStatus()`, `.SetParent()`, `.SetParentID()`, `.SetReviewedAt()`, `.SetReviewedBy()`, `.SetReviewedByID()`, `.SetReviewedByUserID()`, `.SetStatus()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 32`** (86 nodes): `account_update.go`, `AccountUpdate`, `.check()`, `.ClearAccessToken()`, `.ClearAccessTokenExpiresAt()`, `.ClearIDToken()`, `.ClearPassword()`, `.ClearRefreshToken()`, `.ClearRefreshTokenExpiresAt()`, `.ClearScope()`, `.ClearUser()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetAccessToken()`, `.SetAccessTokenExpiresAt()`, `.SetAccountID()`, `.SetCreatedAt()`, `.SetIDToken()`, `.SetNillableAccessToken()`, `.SetNillableAccessTokenExpiresAt()`, `.SetNillableAccountID()`, `.SetNillableCreatedAt()`, `.SetNillableIDToken()`, `.SetNillablePassword()`, `.SetNillableProviderID()`, `.SetNillableRefreshToken()`, `.SetNillableRefreshTokenExpiresAt()`, `.SetNillableScope()`, `.SetNillableUpdatedAt()`, `.SetNillableUserID()`, `.SetPassword()`, `.SetProviderID()`, `.SetRefreshToken()`, `.SetRefreshTokenExpiresAt()`, `.SetScope()`, `.SetUpdatedAt()`, `.SetUser()`, `.SetUserID()`, `.sqlSave()`, `.Where()`, `AccountUpdateOne`, `.check()`, `.ClearAccessToken()`, `.ClearAccessTokenExpiresAt()`, `.ClearIDToken()`, `.ClearPassword()`, `.ClearRefreshToken()`, `.ClearRefreshTokenExpiresAt()`, `.ClearScope()`, `.ClearUser()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetAccessToken()`, `.SetAccessTokenExpiresAt()`, `.SetAccountID()`, `.SetCreatedAt()`, `.SetIDToken()`, `.SetNillableAccessToken()`, `.SetNillableAccessTokenExpiresAt()`, `.SetNillableAccountID()`, `.SetNillableCreatedAt()`, `.SetNillableIDToken()`, `.SetNillablePassword()`, `.SetNillableProviderID()`, `.SetNillableRefreshToken()`, `.SetNillableRefreshTokenExpiresAt()`, `.SetNillableScope()`, `.SetNillableUpdatedAt()`, `.SetNillableUserID()`, `.SetPassword()`, `.SetProviderID()`, `.SetRefreshToken()`, `.SetRefreshTokenExpiresAt()`, `.SetScope()`, `.SetUpdatedAt()`, `.SetUser()`, `.SetUserID()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 33`** (84 nodes): `profilecontent_update.go`, `ProfileContentUpdate`, `.AppendFocus()`, `.AppendStats()`, `.check()`, `.ClearProfilePhotoURL()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetAbout()`, `.SetAvailability()`, `.SetCreatedAt()`, `.SetEmail()`, `.SetFocus()`, `.SetFullName()`, `.SetHeadline()`, `.SetLocation()`, `.SetNillableAbout()`, `.SetNillableAvailability()`, `.SetNillableCreatedAt()`, `.SetNillableEmail()`, `.SetNillableFullName()`, `.SetNillableHeadline()`, `.SetNillableLocation()`, `.SetNillablePhone()`, `.SetNillablePrimaryCta()`, `.SetNillableProfilePhotoURL()`, `.SetNillableShortIntro()`, `.SetNillableStorageKey()`, `.SetPhone()`, `.SetPrimaryCta()`, `.SetProfilePhotoURL()`, `.SetShortIntro()`, `.SetSocialLinks()`, `.SetStats()`, `.SetStorageKey()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`, `ProfileContentUpdateOne`, `.AppendFocus()`, `.AppendStats()`, `.check()`, `.ClearProfilePhotoURL()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetAbout()`, `.SetAvailability()`, `.SetCreatedAt()`, `.SetEmail()`, `.SetFocus()`, `.SetFullName()`, `.SetHeadline()`, `.SetLocation()`, `.SetNillableAbout()`, `.SetNillableAvailability()`, `.SetNillableCreatedAt()`, `.SetNillableEmail()`, `.SetNillableFullName()`, `.SetNillableHeadline()`, `.SetNillableLocation()`, `.SetNillablePhone()`, `.SetNillablePrimaryCta()`, `.SetNillableProfilePhotoURL()`, `.SetNillableShortIntro()`, `.SetNillableStorageKey()`, `.SetPhone()`, `.SetPrimaryCta()`, `.SetProfilePhotoURL()`, `.SetShortIntro()`, `.SetSocialLinks()`, `.SetStats()`, `.SetStorageKey()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 37`** (62 nodes): `session_update.go`, `SessionUpdate`, `.check()`, `.ClearImpersonatedBy()`, `.ClearIPAddress()`, `.ClearUser()`, `.ClearUserAgent()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetCreatedAt()`, `.SetExpiresAt()`, `.SetImpersonatedBy()`, `.SetIPAddress()`, `.SetNillableCreatedAt()`, `.SetNillableExpiresAt()`, `.SetNillableImpersonatedBy()`, `.SetNillableIPAddress()`, `.SetNillableToken()`, `.SetNillableUpdatedAt()`, `.SetNillableUserAgent()`, `.SetNillableUserID()`, `.SetToken()`, `.SetUpdatedAt()`, `.SetUser()`, `.SetUserAgent()`, `.SetUserID()`, `.sqlSave()`, `.Where()`, `SessionUpdateOne`, `.check()`, `.ClearImpersonatedBy()`, `.ClearIPAddress()`, `.ClearUser()`, `.ClearUserAgent()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetCreatedAt()`, `.SetExpiresAt()`, `.SetImpersonatedBy()`, `.SetIPAddress()`, `.SetNillableCreatedAt()`, `.SetNillableExpiresAt()`, `.SetNillableImpersonatedBy()`, `.SetNillableIPAddress()`, `.SetNillableToken()`, `.SetNillableUpdatedAt()`, `.SetNillableUserAgent()`, `.SetNillableUserID()`, `.SetToken()`, `.SetUpdatedAt()`, `.SetUser()`, `.SetUserAgent()`, `.SetUserID()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 38`** (60 nodes): `message_update.go`, `MessageUpdate`, `.check()`, `.ClearArchivedAt()`, `.ClearReadAt()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetArchivedAt()`, `.SetBody()`, `.SetCreatedAt()`, `.SetNillableArchivedAt()`, `.SetNillableBody()`, `.SetNillableCreatedAt()`, `.SetNillableReadAt()`, `.SetNillableSenderEmail()`, `.SetNillableSenderName()`, `.SetNillableStatus()`, `.SetNillableSubject()`, `.SetReadAt()`, `.SetSenderEmail()`, `.SetSenderName()`, `.SetStatus()`, `.SetSubject()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`, `MessageUpdateOne`, `.check()`, `.ClearArchivedAt()`, `.ClearReadAt()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetArchivedAt()`, `.SetBody()`, `.SetCreatedAt()`, `.SetNillableArchivedAt()`, `.SetNillableBody()`, `.SetNillableCreatedAt()`, `.SetNillableReadAt()`, `.SetNillableSenderEmail()`, `.SetNillableSenderName()`, `.SetNillableStatus()`, `.SetNillableSubject()`, `.SetReadAt()`, `.SetSenderEmail()`, `.SetSenderName()`, `.SetStatus()`, `.SetSubject()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 39`** (60 nodes): `experience_update.go`, `ExperienceUpdate`, `.AddSortOrder()`, `.AppendAchievements()`, `.check()`, `.ClearAchievements()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetAchievements()`, `.SetCompany()`, `.SetCreatedAt()`, `.SetLocation()`, `.SetNillableCompany()`, `.SetNillableCreatedAt()`, `.SetNillableLocation()`, `.SetNillablePeriod()`, `.SetNillableRole()`, `.SetNillableSortOrder()`, `.SetNillableSummary()`, `.SetPeriod()`, `.SetRole()`, `.SetSortOrder()`, `.SetSummary()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`, `ExperienceUpdateOne`, `.AddSortOrder()`, `.AppendAchievements()`, `.check()`, `.ClearAchievements()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetAchievements()`, `.SetCompany()`, `.SetCreatedAt()`, `.SetLocation()`, `.SetNillableCompany()`, `.SetNillableCreatedAt()`, `.SetNillableLocation()`, `.SetNillablePeriod()`, `.SetNillableRole()`, `.SetNillableSortOrder()`, `.SetNillableSummary()`, `.SetPeriod()`, `.SetRole()`, `.SetSortOrder()`, `.SetSummary()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 43`** (56 nodes): `education_update.go`, `EducationUpdate`, `.AddSortOrder()`, `.AppendHighlights()`, `.check()`, `.ClearHighlights()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetCreatedAt()`, `.SetDegree()`, `.SetDescription()`, `.SetHighlights()`, `.SetNillableCreatedAt()`, `.SetNillableDegree()`, `.SetNillableDescription()`, `.SetNillablePeriod()`, `.SetNillableSchool()`, `.SetNillableSortOrder()`, `.SetPeriod()`, `.SetSchool()`, `.SetSortOrder()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`, `EducationUpdateOne`, `.AddSortOrder()`, `.AppendHighlights()`, `.check()`, `.ClearHighlights()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetCreatedAt()`, `.SetDegree()`, `.SetDescription()`, `.SetHighlights()`, `.SetNillableCreatedAt()`, `.SetNillableDegree()`, `.SetNillableDescription()`, `.SetNillablePeriod()`, `.SetNillableSchool()`, `.SetNillableSortOrder()`, `.SetPeriod()`, `.SetSchool()`, `.SetSortOrder()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 44`** (54 nodes): `resumeasset_update.go`, `ResumeAssetUpdate`, `.AddFileSizeBytes()`, `.check()`, `.ClearFileName()`, `.ClearFileSizeBytes()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetCreatedAt()`, `.SetDownloadURL()`, `.SetFileName()`, `.SetFileSizeBytes()`, `.SetMimeType()`, `.SetNillableCreatedAt()`, `.SetNillableDownloadURL()`, `.SetNillableFileName()`, `.SetNillableFileSizeBytes()`, `.SetNillableMimeType()`, `.SetNillableStorageKey()`, `.SetStorageKey()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`, `ResumeAssetUpdateOne`, `.AddFileSizeBytes()`, `.check()`, `.ClearFileName()`, `.ClearFileSizeBytes()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetCreatedAt()`, `.SetDownloadURL()`, `.SetFileName()`, `.SetFileSizeBytes()`, `.SetMimeType()`, `.SetNillableCreatedAt()`, `.SetNillableDownloadURL()`, `.SetNillableFileName()`, `.SetNillableFileSizeBytes()`, `.SetNillableMimeType()`, `.SetNillableStorageKey()`, `.SetStorageKey()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 47`** (52 nodes): `visitorlog_update.go`, `VisitorLogUpdate`, `.check()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetIPHash()`, `.SetIsUniqueDailyVisitor()`, `.SetNillableIPHash()`, `.SetNillableIsUniqueDailyVisitor()`, `.SetNillablePath()`, `.SetNillableReferrer()`, `.SetNillableReferrerSource()`, `.SetNillableUserAgent()`, `.SetNillableVisitedAt()`, `.SetNillableVisitorID()`, `.SetPath()`, `.SetReferrer()`, `.SetReferrerSource()`, `.SetUserAgent()`, `.SetVisitedAt()`, `.SetVisitorID()`, `.sqlSave()`, `.Where()`, `VisitorLogUpdateOne`, `.check()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetIPHash()`, `.SetIsUniqueDailyVisitor()`, `.SetNillableIPHash()`, `.SetNillableIsUniqueDailyVisitor()`, `.SetNillablePath()`, `.SetNillableReferrer()`, `.SetNillableReferrerSource()`, `.SetNillableUserAgent()`, `.SetNillableVisitedAt()`, `.SetNillableVisitorID()`, `.SetPath()`, `.SetReferrer()`, `.SetReferrerSource()`, `.SetUserAgent()`, `.SetVisitedAt()`, `.SetVisitorID()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 50`** (49 nodes): `user_create.go`, `UserCreate`, `.AddAccountIDs()`, `.AddAccounts()`, `.AddBlogPostIDs()`, `.AddBlogPosts()`, `.AddReviewedCommentIDs()`, `.AddReviewedComments()`, `.AddSessionIDs()`, `.AddSessions()`, `.AddTwoFactorIDs()`, `.AddTwoFactors()`, `.check()`, `.createSpec()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetBanExpires()`, `.SetBanned()`, `.SetBanReason()`, `.SetCreatedAt()`, `.SetDisplayUsername()`, `.SetEmail()`, `.SetEmailVerified()`, `.SetID()`, `.SetImage()`, `.SetIsAnonymous()`, `.SetName()`, `.SetNillableBanExpires()`, `.SetNillableBanned()`, `.SetNillableBanReason()`, `.SetNillableDisplayUsername()`, `.SetNillableImage()`, `.SetNillableIsAnonymous()`, `.SetNillableRole()`, `.SetNillableTwoFactorEnabled()`, `.SetNillableUsername()`, `.SetRole()`, `.SetTwoFactorEnabled()`, `.SetUpdatedAt()`, `.SetUsername()`, `.sqlSave()`, `UserCreateBulk`, `.Exec()`, `.ExecX()`, `.Save()`, `.SaveX()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 51`** (45 nodes): `blogpost_create.go`, `BlogPostCreate`, `.AddCommentIDs()`, `.AddComments()`, `.check()`, `.createSpec()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetAuthor()`, `.SetAuthorID()`, `.SetAuthorName()`, `.SetAuthorUserID()`, `.SetCategory()`, `.SetContent()`, `.SetCoverImagePlaceholder()`, `.SetCreatedAt()`, `.SetExcerpt()`, `.SetFeatured()`, `.SetID()`, `.SetNillableCoverImagePlaceholder()`, `.SetNillableCreatedAt()`, `.SetNillableFeatured()`, `.SetNillablePublishDate()`, `.SetNillableSeoDescription()`, `.SetNillableSeoTitle()`, `.SetNillableStatus()`, `.SetPublishDate()`, `.SetReadingTime()`, `.SetSeoDescription()`, `.SetSeoTitle()`, `.SetSlug()`, `.SetStatus()`, `.SetTags()`, `.SetTitle()`, `.SetUpdatedAt()`, `.sqlSave()`, `BlogPostCreateBulk`, `.Exec()`, `.ExecX()`, `.Save()`, `.SaveX()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 52`** (44 nodes): `skill_update.go`, `SkillUpdate`, `.check()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetCategory()`, `.SetCreatedAt()`, `.SetFeatured()`, `.SetLevel()`, `.SetName()`, `.SetNillableCategory()`, `.SetNillableCreatedAt()`, `.SetNillableFeatured()`, `.SetNillableLevel()`, `.SetNillableName()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`, `SkillUpdateOne`, `.check()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.Select()`, `.SetCategory()`, `.SetCreatedAt()`, `.SetFeatured()`, `.SetLevel()`, `.SetName()`, `.SetNillableCategory()`, `.SetNillableCreatedAt()`, `.SetNillableFeatured()`, `.SetNillableLevel()`, `.SetNillableName()`, `.SetUpdatedAt()`, `.sqlSave()`, `.Where()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 53`** (42 nodes): `AdminDashboard.tsx`, `handleNavigate()`, `Icon()`, `IconChart()`, `IconEdit()`, `IconEye()`, `IconFolder()`, `IconMail()`, `IconPlus()`, `IconStar()`, `IconTrash()`, `IsoChest()`, `renderPage()`, `StatTile()`, `statusTag()`, `Analytics()`, `CmsApp()`, `Dashboard()`, `Icon()`, `IconBell()`, `IconChart()`, `IconCog()`, `IconDashboard()`, `IconEdit()`, `IconEye()`, `IconFolder()`, `IconMail()`, `IconPlus()`, `IconPost()`, `IconSearch()`, `IconStar()`, `IconTrash()`, `IsoChest()`, `Messages()`, `Posts()`, `Projects()`, `ProjectsTable()`, `px()`, `Settings()`, `Skills()`, `StatTile()`, `CmsApp.jsx`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 54`** (41 nodes): `blogcomment_create.go`, `BlogCommentCreate`, `.AddChildIDs()`, `.AddChildren()`, `.check()`, `.createSpec()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetBlogPost()`, `.SetBlogPostID()`, `.SetBody()`, `.SetCreatedAt()`, `.SetDisplayName()`, `.SetEmail()`, `.SetFingerprint()`, `.SetID()`, `.SetIPHash()`, `.SetNillableCreatedAt()`, `.SetNillableParentID()`, `.SetNillableReviewedAt()`, `.SetNillableReviewedByID()`, `.SetNillableReviewedByUserID()`, `.SetNillableStatus()`, `.SetParent()`, `.SetParentID()`, `.SetReviewedAt()`, `.SetReviewedBy()`, `.SetReviewedByID()`, `.SetReviewedByUserID()`, `.SetStatus()`, `.SetUpdatedAt()`, `.sqlSave()`, `BlogCommentCreateBulk`, `.Exec()`, `.ExecX()`, `.Save()`, `.SaveX()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 55`** (36 nodes): `profilecontent_create.go`, `ProfileContentCreate`, `.check()`, `.createSpec()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetAbout()`, `.SetAvailability()`, `.SetCreatedAt()`, `.SetEmail()`, `.SetFocus()`, `.SetFullName()`, `.SetHeadline()`, `.SetID()`, `.SetLocation()`, `.SetNillableCreatedAt()`, `.SetNillableProfilePhotoURL()`, `.SetNillableStorageKey()`, `.SetPhone()`, `.SetPrimaryCta()`, `.SetProfilePhotoURL()`, `.SetShortIntro()`, `.SetSocialLinks()`, `.SetStats()`, `.SetStorageKey()`, `.SetUpdatedAt()`, `.sqlSave()`, `ProfileContentCreateBulk`, `.Exec()`, `.ExecX()`, `.Save()`, `.SaveX()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 56`** (30 nodes): `message_create.go`, `MessageCreate`, `.check()`, `.createSpec()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetArchivedAt()`, `.SetBody()`, `.SetCreatedAt()`, `.SetID()`, `.SetNillableArchivedAt()`, `.SetNillableCreatedAt()`, `.SetNillableReadAt()`, `.SetNillableStatus()`, `.SetReadAt()`, `.SetSenderEmail()`, `.SetSenderName()`, `.SetStatus()`, `.SetSubject()`, `.SetUpdatedAt()`, `.sqlSave()`, `MessageCreateBulk`, `.Exec()`, `.ExecX()`, `.Save()`, `.SaveX()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 57`** (29 nodes): `visitorlog_create.go`, `VisitorLogCreate`, `.check()`, `.createSpec()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetID()`, `.SetIPHash()`, `.SetIsUniqueDailyVisitor()`, `.SetNillableIsUniqueDailyVisitor()`, `.SetNillableReferrer()`, `.SetNillableReferrerSource()`, `.SetNillableVisitedAt()`, `.SetPath()`, `.SetReferrer()`, `.SetReferrerSource()`, `.SetUserAgent()`, `.SetVisitedAt()`, `.SetVisitorID()`, `.sqlSave()`, `VisitorLogCreateBulk`, `.Exec()`, `.ExecX()`, `.Save()`, `.SaveX()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 58`** (29 nodes): `resumeasset_create.go`, `ResumeAssetCreate`, `.check()`, `.createSpec()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetCreatedAt()`, `.SetDownloadURL()`, `.SetFileName()`, `.SetFileSizeBytes()`, `.SetID()`, `.SetMimeType()`, `.SetNillableCreatedAt()`, `.SetNillableFileName()`, `.SetNillableFileSizeBytes()`, `.SetNillableMimeType()`, `.SetNillableStorageKey()`, `.SetStorageKey()`, `.SetUpdatedAt()`, `.sqlSave()`, `ResumeAssetCreateBulk`, `.Exec()`, `.ExecX()`, `.Save()`, `.SaveX()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 59`** (28 nodes): `experience_create.go`, `ExperienceCreate`, `.check()`, `.createSpec()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetAchievements()`, `.SetCompany()`, `.SetCreatedAt()`, `.SetID()`, `.SetLocation()`, `.SetNillableCreatedAt()`, `.SetNillableSortOrder()`, `.SetPeriod()`, `.SetRole()`, `.SetSortOrder()`, `.SetSummary()`, `.SetUpdatedAt()`, `.sqlSave()`, `ExperienceCreateBulk`, `.Exec()`, `.ExecX()`, `.Save()`, `.SaveX()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 60`** (25 nodes): `skill_create.go`, `SkillCreate`, `.check()`, `.createSpec()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetCategory()`, `.SetCreatedAt()`, `.SetFeatured()`, `.SetID()`, `.SetLevel()`, `.SetName()`, `.SetNillableCreatedAt()`, `.SetNillableFeatured()`, `.SetUpdatedAt()`, `.sqlSave()`, `SkillCreateBulk`, `.Exec()`, `.ExecX()`, `.Save()`, `.SaveX()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 61`** (23 nodes): `cvdownloadlog_create.go`, `CVDownloadLogCreate`, `.check()`, `.createSpec()`, `.defaults()`, `.Exec()`, `.ExecX()`, `.Mutation()`, `.Save()`, `.SaveX()`, `.SetDownloadedAt()`, `.SetID()`, `.SetIPHash()`, `.SetNillableDownloadedAt()`, `.SetNillableReferrer()`, `.SetReferrer()`, `.SetUserAgent()`, `.sqlSave()`, `CVDownloadLogCreateBulk`, `.Exec()`, `.ExecX()`, `.Save()`, `.SaveX()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 64`** (17 nodes): `profilecontent.go`, `OrderOption`, `ByAbout()`, `ByAvailability()`, `ByCreatedAt()`, `ByEmail()`, `ByFullName()`, `ByHeadline()`, `ByID()`, `ByLocation()`, `ByPhone()`, `ByPrimaryCta()`, `ByProfilePhotoURL()`, `ByShortIntro()`, `ByStorageKey()`, `ByUpdatedAt()`, `ValidColumn()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 65`** (16 nodes): `message.go`, `ByArchivedAt()`, `ByBody()`, `ByCreatedAt()`, `ByID()`, `ByReadAt()`, `BySenderEmail()`, `BySenderName()`, `ByStatus()`, `BySubject()`, `ByUpdatedAt()`, `StatusValidator()`, `ValidColumn()`, `OrderOption`, `Status`, `.String()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 68`** (13 nodes): `certificate.go`, `ByCreatedAt()`, `ByCredentialID()`, `ByFeatured()`, `ByID()`, `ByIssuer()`, `ByName()`, `BySortOrder()`, `ByUpdatedAt()`, `ByVerificationLink()`, `ByYear()`, `ValidColumn()`, `OrderOption`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 70`** (12 nodes): `experience.go`, `ByCompany()`, `ByCreatedAt()`, `ByID()`, `ByLocation()`, `ByPeriod()`, `ByRole()`, `BySortOrder()`, `BySummary()`, `ByUpdatedAt()`, `ValidColumn()`, `OrderOption`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 71`** (12 nodes): `visitorlog.go`, `OrderOption`, `ByID()`, `ByIPHash()`, `ByIsUniqueDailyVisitor()`, `ByPath()`, `ByReferrer()`, `ByReferrerSource()`, `ByUserAgent()`, `ByVisitedAt()`, `ByVisitorID()`, `ValidColumn()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 73`** (11 nodes): `resumeasset.go`, `OrderOption`, `ByCreatedAt()`, `ByDownloadURL()`, `ByFileName()`, `ByFileSizeBytes()`, `ByID()`, `ByMimeType()`, `ByStorageKey()`, `ByUpdatedAt()`, `ValidColumn()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 74`** (11 nodes): `education.go`, `ByCreatedAt()`, `ByDegree()`, `ByDescription()`, `ByID()`, `ByPeriod()`, `BySchool()`, `BySortOrder()`, `ByUpdatedAt()`, `ValidColumn()`, `OrderOption`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 77`** (10 nodes): `skill.go`, `OrderOption`, `ByCategory()`, `ByCreatedAt()`, `ByFeatured()`, `ByID()`, `ByLevel()`, `ByName()`, `ByUpdatedAt()`, `ValidColumn()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 81`** (9 nodes): `verification.go`, `OrderOption`, `ByCreatedAt()`, `ByExpiresAt()`, `ByID()`, `ByIdentifier()`, `ByUpdatedAt()`, `ByValue()`, `ValidColumn()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 83`** (8 nodes): `cvdownloadlog.go`, `ByDownloadedAt()`, `ByID()`, `ByIPHash()`, `ByReferrer()`, `ByUserAgent()`, `ValidColumn()`, `OrderOption`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Community 85`** (6 nodes): `RouteErrorBoundary.tsx`, `RouteErrorBoundary`, `.componentDidCatch()`, `.componentDidUpdate()`, `.getDerivedStateFromError()`, `.render()`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `IsConstraintError()` connect `Community 3` to `Community 13`, `Community 16`, `Community 17`, `Community 22`, `Community 28`, `Community 30`, `Community 32`, `Community 33`, `Community 34`, `Community 37`, `Community 38`, `Community 39`, `Community 43`, `Community 44`, `Community 47`, `Community 49`, `Community 50`, `Community 51`, `Community 52`, `Community 54`, `Community 55`, `Community 56`, `Community 57`, `Community 58`, `Community 59`, `Community 60`, `Community 61`?**
  _High betweenness centrality (0.105) - this node is a cross-community bridge._
- **Why does `Open()` connect `Community 5` to `Community 0`?**
  _High betweenness centrality (0.044) - this node is a cross-community bridge._
- **Are the 180 inferred relationships involving `setContextOp()` (e.g. with `.First()` and `.FirstID()`) actually correct?**
  _`setContextOp()` has 180 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AuthConfig`, `AuthCookieConfig`, `RootResponse` to the rest of the system?**
  _159 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Community 0` be split into smaller, more focused modules?**
  _Cohesion score 0.01 - nodes in this community are weakly interconnected._
- **Should `Community 1` be split into smaller, more focused modules?**
  _Cohesion score 0.0 - nodes in this community are weakly interconnected._
- **Should `Community 2` be split into smaller, more focused modules?**
  _Cohesion score 0.01 - nodes in this community are weakly interconnected._
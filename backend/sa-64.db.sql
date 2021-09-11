BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "users" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "videos" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"url"	text,
	"owner_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_users_videos" FOREIGN KEY("owner_id") REFERENCES "users"("id")
);
CREATE TABLE IF NOT EXISTS "playlists" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"title"	text,
	"owner_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_users_playlists" FOREIGN KEY("owner_id") REFERENCES "users"("id")
);
CREATE TABLE IF NOT EXISTS "resolutions" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"value"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "watch_videos" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"watched_time"	datetime,
	"resolution_id"	integer,
	"playlist_id"	integer,
	"video_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_videos_watch_videos" FOREIGN KEY("video_id") REFERENCES "videos"("id"),
	CONSTRAINT "fk_resolutions_watch_videos" FOREIGN KEY("resolution_id") REFERENCES "resolutions"("id"),
	CONSTRAINT "fk_playlists_watch_videos" FOREIGN KEY("playlist_id") REFERENCES "playlists"("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_email" ON "users" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_videos_deleted_at" ON "videos" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_videos_url" ON "videos" (
	"url"
);
CREATE INDEX IF NOT EXISTS "idx_playlists_deleted_at" ON "playlists" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_resolutions_deleted_at" ON "resolutions" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_watch_videos_deleted_at" ON "watch_videos" (
	"deleted_at"
);
COMMIT;

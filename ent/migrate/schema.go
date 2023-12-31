// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlockedUsersRelationshipsColumns holds the columns for the "blocked_users_relationships" table.
	BlockedUsersRelationshipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
		{Name: "blocker_id", Type: field.TypeInt, Nullable: true},
	}
	// BlockedUsersRelationshipsTable holds the schema information for the "blocked_users_relationships" table.
	BlockedUsersRelationshipsTable = &schema.Table{
		Name:       "blocked_users_relationships",
		Columns:    BlockedUsersRelationshipsColumns,
		PrimaryKey: []*schema.Column{BlockedUsersRelationshipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blocked_users_relationships_user_account_blocked_by",
				Columns:    []*schema.Column{BlockedUsersRelationshipsColumns[3]},
				RefColumns: []*schema.Column{UserAccountColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "blocked_users_relationships_user_account_blocked_users",
				Columns:    []*schema.Column{BlockedUsersRelationshipsColumns[4]},
				RefColumns: []*schema.Column{UserAccountColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MediaColumns holds the columns for the "media" table.
	MediaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "path", Type: field.TypeString},
		{Name: "category", Type: field.TypeEnum, Enums: []string{"IMAGE", "DOCUMENT", "OTHER"}},
	}
	// MediaTable holds the schema information for the "media" table.
	MediaTable = &schema.Table{
		Name:       "media",
		Columns:    MediaColumns,
		PrimaryKey: []*schema.Column{MediaColumns[0]},
	}
	// ThreadsColumns holds the columns for the "threads" table.
	ThreadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "is_comment_disabled", Type: field.TypeBool, Default: false},
		{Name: "visibility", Type: field.TypeEnum, Enums: []string{"PRIVATE", "PUBLIC"}, Default: "PUBLIC"},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PUBLISHED", "DRAFT"}, Default: "PUBLISHED"},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
		{Name: "repost_thread_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "author_id", Type: field.TypeInt, Nullable: true},
	}
	// ThreadsTable holds the schema information for the "threads" table.
	ThreadsTable = &schema.Table{
		Name:       "threads",
		Columns:    ThreadsColumns,
		PrimaryKey: []*schema.Column{ThreadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "threads_threads_children",
				Columns:    []*schema.Column{ThreadsColumns[8]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "threads_threads_repost",
				Columns:    []*schema.Column{ThreadsColumns[9]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "threads_user_account_threads",
				Columns:    []*schema.Column{ThreadsColumns[10]},
				RefColumns: []*schema.Column{UserAccountColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ThreadAccountColumns holds the columns for the "thread_account" table.
	ThreadAccountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "reply_count", Type: field.TypeInt, Default: 0},
		{Name: "like_count", Type: field.TypeInt, Default: 0},
		{Name: "thread_thread_count", Type: field.TypeInt},
	}
	// ThreadAccountTable holds the schema information for the "thread_account" table.
	ThreadAccountTable = &schema.Table{
		Name:       "thread_account",
		Columns:    ThreadAccountColumns,
		PrimaryKey: []*schema.Column{ThreadAccountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "thread_account_threads_thread_count",
				Columns:    []*schema.Column{ThreadAccountColumns[3]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ThreadLikeUsersColumns holds the columns for the "thread_like_users" table.
	ThreadLikeUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ThreadLikeUsersTable holds the schema information for the "thread_like_users" table.
	ThreadLikeUsersTable = &schema.Table{
		Name:       "thread_like_users",
		Columns:    ThreadLikeUsersColumns,
		PrimaryKey: []*schema.Column{ThreadLikeUsersColumns[0]},
	}
	// UserAccountColumns holds the columns for the "user_account" table.
	UserAccountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeBytes},
		{Name: "is_verified", Type: field.TypeBool, Default: false},
		{Name: "is_private", Type: field.TypeBool, Default: false},
		{Name: "is_email_verified", Type: field.TypeBool, Default: false},
	}
	// UserAccountTable holds the schema information for the "user_account" table.
	UserAccountTable = &schema.Table{
		Name:       "user_account",
		Columns:    UserAccountColumns,
		PrimaryKey: []*schema.Column{UserAccountColumns[0]},
	}
	// UserCountColumns holds the columns for the "user_count" table.
	UserCountColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "follower_count", Type: field.TypeInt, Default: 0},
		{Name: "followings_count", Type: field.TypeInt, Default: 0},
		{Name: "user_account_user_count", Type: field.TypeInt, Unique: true},
	}
	// UserCountTable holds the schema information for the "user_count" table.
	UserCountTable = &schema.Table{
		Name:       "user_count",
		Columns:    UserCountColumns,
		PrimaryKey: []*schema.Column{UserCountColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_count_user_account_user_count",
				Columns:    []*schema.Column{UserCountColumns[3]},
				RefColumns: []*schema.Column{UserAccountColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserFollowerRelationshipsColumns holds the columns for the "user_follower_relationships" table.
	UserFollowerRelationshipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
		{Name: "follower_id", Type: field.TypeInt, Nullable: true},
	}
	// UserFollowerRelationshipsTable holds the schema information for the "user_follower_relationships" table.
	UserFollowerRelationshipsTable = &schema.Table{
		Name:       "user_follower_relationships",
		Columns:    UserFollowerRelationshipsColumns,
		PrimaryKey: []*schema.Column{UserFollowerRelationshipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_follower_relationships_user_account_followers",
				Columns:    []*schema.Column{UserFollowerRelationshipsColumns[3]},
				RefColumns: []*schema.Column{UserAccountColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "user_follower_relationships_user_account_followings",
				Columns:    []*schema.Column{UserFollowerRelationshipsColumns[4]},
				RefColumns: []*schema.Column{UserAccountColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserProfileColumns holds the columns for the "user_profile" table.
	UserProfileColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "display_name", Type: field.TypeString},
		{Name: "bio", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "user_account_profile", Type: field.TypeInt, Unique: true},
		{Name: "profile_picture_id", Type: field.TypeInt, Nullable: true},
		{Name: "banner_id", Type: field.TypeInt, Nullable: true},
	}
	// UserProfileTable holds the schema information for the "user_profile" table.
	UserProfileTable = &schema.Table{
		Name:       "user_profile",
		Columns:    UserProfileColumns,
		PrimaryKey: []*schema.Column{UserProfileColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_profile_user_account_profile",
				Columns:    []*schema.Column{UserProfileColumns[3]},
				RefColumns: []*schema.Column{UserAccountColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_profile_media_profile_picture",
				Columns:    []*schema.Column{UserProfileColumns[4]},
				RefColumns: []*schema.Column{MediaColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "user_profile_media_banner",
				Columns:    []*schema.Column{UserProfileColumns[5]},
				RefColumns: []*schema.Column{MediaColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ThreadImagesColumns holds the columns for the "thread_images" table.
	ThreadImagesColumns = []*schema.Column{
		{Name: "thread_id", Type: field.TypeInt},
		{Name: "media_id", Type: field.TypeInt},
	}
	// ThreadImagesTable holds the schema information for the "thread_images" table.
	ThreadImagesTable = &schema.Table{
		Name:       "thread_images",
		Columns:    ThreadImagesColumns,
		PrimaryKey: []*schema.Column{ThreadImagesColumns[0], ThreadImagesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "thread_images_thread_id",
				Columns:    []*schema.Column{ThreadImagesColumns[0]},
				RefColumns: []*schema.Column{ThreadsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "thread_images_media_id",
				Columns:    []*schema.Column{ThreadImagesColumns[1]},
				RefColumns: []*schema.Column{MediaColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlockedUsersRelationshipsTable,
		MediaTable,
		ThreadsTable,
		ThreadAccountTable,
		ThreadLikeUsersTable,
		UserAccountTable,
		UserCountTable,
		UserFollowerRelationshipsTable,
		UserProfileTable,
		ThreadImagesTable,
	}
)

func init() {
	BlockedUsersRelationshipsTable.ForeignKeys[0].RefTable = UserAccountTable
	BlockedUsersRelationshipsTable.ForeignKeys[1].RefTable = UserAccountTable
	ThreadsTable.ForeignKeys[0].RefTable = ThreadsTable
	ThreadsTable.ForeignKeys[1].RefTable = ThreadsTable
	ThreadsTable.ForeignKeys[2].RefTable = UserAccountTable
	ThreadAccountTable.ForeignKeys[0].RefTable = ThreadsTable
	ThreadAccountTable.Annotation = &entsql.Annotation{
		Table: "thread_account",
	}
	UserAccountTable.Annotation = &entsql.Annotation{
		Table: "user_account",
	}
	UserCountTable.ForeignKeys[0].RefTable = UserAccountTable
	UserCountTable.Annotation = &entsql.Annotation{
		Table: "user_count",
	}
	UserFollowerRelationshipsTable.ForeignKeys[0].RefTable = UserAccountTable
	UserFollowerRelationshipsTable.ForeignKeys[1].RefTable = UserAccountTable
	UserProfileTable.ForeignKeys[0].RefTable = UserAccountTable
	UserProfileTable.ForeignKeys[1].RefTable = MediaTable
	UserProfileTable.ForeignKeys[2].RefTable = MediaTable
	UserProfileTable.Annotation = &entsql.Annotation{
		Table: "user_profile",
	}
	ThreadImagesTable.ForeignKeys[0].RefTable = ThreadsTable
	ThreadImagesTable.ForeignKeys[1].RefTable = MediaTable
}

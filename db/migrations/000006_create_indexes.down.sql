-- Users テーブルのインデックス削除
DROP INDEX IF EXISTS idx_users_id;

-- Groups テーブルのインデックス削除
DROP INDEX IF EXISTS idx_groups_name;

-- GroupMembers テーブルのインデックス削除
DROP INDEX IF EXISTS idx_group_members_group_id;
DROP INDEX IF EXISTS idx_group_members_user_id;
DROP INDEX IF EXISTS idx_group_members_admin;
DROP INDEX IF EXISTS idx_groups_id;

-- UserEvents テーブルのインデックス削除
DROP INDEX IF EXISTS idx_user_events_user_id;
DROP INDEX IF EXISTS idx_user_events_created_at;
DROP INDEX IF EXISTS idx_user_events_event_name;
DROP INDEX IF EXISTS idx_user_events_emotion;
DROP INDEX IF EXISTS idx_user_events_user_created;

-- Reactions テーブルのインデックス削除
DROP INDEX IF EXISTS idx_reactions_event_id;
DROP INDEX IF EXISTS idx_reactions_user_id;
DROP INDEX IF EXISTS idx_reactions_created_at;
DROP INDEX IF EXISTS idx_reactions_reaction_type;
DROP INDEX IF EXISTS idx_reactions_event_user;

-- Users テーブルのインデックス削除
DROP INDEX IF EXISTS idx_users_id;

-- Groups テーブルのインデックス削除
DROP INDEX IF EXISTS idx_groups_id;

-- GroupMembers テーブルのインデックス削除
DROP INDEX IF EXISTS idx_group_members_group_id;
DROP INDEX IF EXISTS idx_group_members_user_id;

-- UserEvents テーブルのインデックス削除
DROP INDEX IF EXISTS idx_user_events_user_id;

-- Reactions テーブルのインデックス削除
DROP INDEX IF EXISTS idx_reactions_event_id;
DROP INDEX IF EXISTS idx_reactions_user_id;
DROP INDEX IF EXISTS idx_reactions_event_user;

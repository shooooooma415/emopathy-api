-- Users テーブルのインデックス
CREATE INDEX idx_users_id ON users(id);

-- Groups テーブルのインデックス
CREATE INDEX idx_groups_id ON groups(id);

-- GroupMembers テーブルのインデックス
CREATE INDEX idx_group_members_group_id ON group_members(group_id);
CREATE INDEX idx_group_members_user_id ON group_members(user_id);

-- UserEvents テーブルのインデックス
CREATE INDEX idx_user_events_user_id ON user_events(user_id);
CREATE INDEX idx_user_events_user_created ON user_events(group_id, created_at);

-- Reactions テーブルのインデックス
CREATE INDEX idx_reactions_event_id ON reactions(event_id);
CREATE INDEX idx_reactions_user_id ON reactions(user_id);
CREATE INDEX idx_reactions_event_user ON reactions(event_id, user_id);

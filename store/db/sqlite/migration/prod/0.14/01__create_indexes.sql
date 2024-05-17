CREATE INDEX IF NOT EXISTS idx_user_username ON user (username);
CREATE INDEX IF NOT EXISTS idx_locket_creator_id ON locket (creator_id);
CREATE INDEX IF NOT EXISTS idx_locket_content ON locket (content);
CREATE INDEX IF NOT EXISTS idx_locket_visibility ON locket (visibility);
CREATE INDEX IF NOT EXISTS idx_resource_creator_id ON resource (creator_id);

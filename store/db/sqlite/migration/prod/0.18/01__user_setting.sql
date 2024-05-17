UPDATE user_setting SET key = 'USER_SETTING_LOCALE', value = REPLACE(value, '"', '') WHERE key = 'locale';
UPDATE user_setting SET key = 'USER_SETTING_APPEARANCE', value = REPLACE(value, '"', '') WHERE key = 'appearance';
UPDATE user_setting SET key = 'USER_SETTING_LOCKET_VISIBILITY', value = REPLACE(value, '"', '') WHERE key = 'locket-visibility';
UPDATE user_setting SET key = 'USER_SETTING_TELEGRAM_USER_ID', value = REPLACE(value, '"', '') WHERE key = 'telegram-user-id';

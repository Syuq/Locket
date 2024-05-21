INSERT INTO
  user (
    `id`,
    `username`,
    `role`,
    `email`,
    `nickname`,
    `password_hash`,
    `description`
  )
VALUES
  (
    101,
    'lockets-demo',
    'HOST',
    'demo@uselockets.com',
    'Demo',
    -- raw password: secret
    '$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK',
    '👋 Welcome to lockets.'
  );

INSERT INTO
  user (
    `id`,
    `username`,
    `role`,
    `email`,
    `nickname`,
    `password_hash`,
    `description`
  )
VALUES
  (
    102,
    'jack',
    'USER',
    'faker@lockets.com',
    'Faker',
    -- raw password: secret
    '$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK',
    'The REAL Faker.'
  );

INSERT INTO
  user (
    `id`,
    `row_status`,
    `username`,
    `role`,
    `email`,
    `nickname`,
    `password_hash`,
    `description`
  )
VALUES
  (
    103,
    'ARCHIVED',
    'bob',
    'USER',
    'locket@lockets.com',
    'Locket',
    -- raw password: secret
    '$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK',
    'Sorry.'
  );
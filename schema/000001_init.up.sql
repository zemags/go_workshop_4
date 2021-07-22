create table memo (
  id integer not null auto_increment primary key,
  title varchar(255) not null,
  content text not null,
  created datetime not null,
  expires datetime not null
);
create index idx_memo_created on memo(created);
INSERT INTO
  memo (title, content, created, expires)
VALUES
  (
    'Apple',
    'Buy apple',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
  );
INSERT INTO
  memo (title, content, created, expires)
VALUES
  (
    'Orange',
    'Buy two oranges',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
  );
INSERT INTO
  memo (title, content, created, expires)
VALUES
  (
    'Banana',
    'Find banana',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
  );
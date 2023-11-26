CREATE TABLE `TodoList` (
  `todo_id`     INTEGER  NOT NULL PRIMARY KEY AUTOINCREMENT,
  `name`        TEXT     NOT NULL,
  `kind`        INTEGER  NOT NULL DEFAULT 0,
  `state`       INTEGER  NOT NULL DEFAULT 0
)

INSERT INTO TodoList
VALUES (1, 'finish todo app', 0, 0)

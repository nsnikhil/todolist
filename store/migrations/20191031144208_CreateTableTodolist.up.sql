CREATE TABLE todolist (
id uuid PRIMARY KEY NOT NULL,
description text NOT NULL,
status boolean NOT NULL DEFAULT 't'
);
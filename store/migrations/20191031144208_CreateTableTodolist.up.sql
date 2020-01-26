CREATE TABLE todolist (
id uuid PRIMARY KEY NOT NULL,
title text NOT NULL,
description text ,
status boolean DEFAULT 'f',
tags text[]
);
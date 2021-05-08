-- +migrate Up
insert into users values(1, "2021-05-01 12:45:35", "2021-05-01 12:45:35", null, "user1", "", "08011112222", 0, "user1@example.com", "password");

insert into hosts values(1, "2021-05-01 11:45:35", "2021-05-01 11:45:35", null, "host1", "tanaka", "", "", "012000001111", "tokyo", "", true, false);

insert into plans values(1, "2021-05-01 11:45:35", "2021-05-01 11:45:35", null, 1, "planA", "description: planA", 8000, "share", "", 24, "", "", "", null);
insert into plans values(2, "2021-05-01 11:45:35", "2021-05-01 11:45:35", null, 2, "planB", "description: planB", 12000, "share", "", 12, "", "", "", null);
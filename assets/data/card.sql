create table cards (
  id int AUTO_INCREMENT PRIMARY KEY ,
  type int(10) not null, /* 1 for vocab, 2 for code */
  front text not null,
  back text not null,
  known int(10) default 0
) default charset=utf8mb4;
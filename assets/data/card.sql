create table cards (
  id int AUTO_INCREMENT PRIMARY KEY ,
  front text not null,
  back text not null,
  known int(10) default 0
) default charset=utf8mb4;
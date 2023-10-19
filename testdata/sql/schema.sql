create table user
(
    id        varchar(36)  not null
        primary key
        unique,
    name      varchar(128) not null,
    firstname varchar(255) not null
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;

insert into user (id, name, firstname) values ('ff3fefb0-e617-4e52-9f80-96769451520f', 'John', 'Doe');
insert into user (id, name, firstname) values ('9eeb4a62-c4cc-4be1-ac63-ca392154b65b', 'Gweneth', 'Vasyaev');
insert into user (id, name, firstname) values ('c26f311f-a8c0-4678-81f0-cd1b56ccdcf2', 'Ripley', 'Myhill');
insert into user (id, name, firstname) values ('58c8e7aa-2ff2-4ecf-8c4d-ce3cda7fcc80', 'Mandel', 'Guilayn');
insert into user (id, name, firstname) values ('453207d5-b316-46ac-9fe0-78635fe42530', 'Carina', 'Larmuth');
insert into user (id, name, firstname) values ('1bb0715a-b7b3-4ab7-9f55-42a06fb0b675', 'Finlay', 'Lilleyman');
insert into user (id, name, firstname) values ('2fe8a7ca-58e0-43d1-bc19-9c9ff29a2565', 'Judd', 'Cuerdale');
insert into user (id, name, firstname) values ('4f7deb74-d545-4d81-9a55-3dca70fd59f5', 'Erminia', 'Farnworth');
insert into user (id, name, firstname) values ('cd978e7e-1484-4be5-9e9d-ea365676284c', 'Joan', 'Sowle');
insert into user (id, name, firstname) values ('e8c58bf5-532d-458b-8c6d-b2f828f48b6c', 'Regina', 'McLennan');

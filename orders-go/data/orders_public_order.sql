create table "order"
(
    id              serial                              not null
        constraint order_pkey
            primary key,
    user_id         text                                not null,
    total_units     integer   default 0                 not null,
    total_nanos     integer   default 0                 not null,
    contact_name    text                                not null,
    phone           text                                not null,
    city            text                                not null,
    address1        text                                not null,
    address2        text,
    order_status_id integer                             not null
        constraint order_order_status_id_fk
            references order_status
            on update cascade on delete cascade,
    created_at      timestamp default CURRENT_TIMESTAMP not null
);

alter table "order"
    owner to postgres;

INSERT INTO public."order" (id, user_id, total_units, total_nanos, contact_name, phone, city, address1, address2, order_status_id, created_at) VALUES (default, '2', 200, 40, 'Contact Name 2', '+380687333231', 'Odessa', 'NP #40', null, 2, '2020-12-22 22:04:45.489884');
INSERT INTO public."order" (id, user_id, total_units, total_nanos, contact_name, phone, city, address1, address2, order_status_id, created_at) VALUES (default, '1', 1, 0, 'Contact Name 1', '+380687333230', 'Odessa', 'NP #40', null, 1, '2020-12-21 22:04:45.489884');

ALTER SEQUENCE order_id_seq RESTART WITH 1000;
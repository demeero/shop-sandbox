create table order_item
(
    id           serial            not null
        constraint order_item_pkey
            primary key,
    quantity     integer           not null,
    total_units  integer default 0 not null,
    total_nanos  integer default 0 not null,
    product_id   text              not null,
    product_name text              not null,
    order_id     integer           not null
        constraint order_item_order_fk
            references "order"
);

alter table order_item
    owner to postgres;

INSERT INTO public.order_item (id, quantity, total_units, total_nanos, product_id, product_name, order_id) VALUES (1, 1, 1, 0, '1', 'Product Name 1', 1);
INSERT INTO public.order_item (id, quantity, total_units, total_nanos, product_id, product_name, order_id) VALUES (2, 2, 100, 20, '2', 'Product Name 2', 2);
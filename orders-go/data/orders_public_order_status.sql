create table order_status
(
    id   serial not null
        constraint order_status_pkey
            primary key,
    name text   not null
        constraint order_status_name_key
            unique
);

alter table order_status
    owner to postgres;

INSERT INTO public.order_status (id, name) VALUES (1, 'pending');
INSERT INTO public.order_status (id, name) VALUES (2, 'processing');
INSERT INTO public.order_status (id, name) VALUES (3, 'canceled');
INSERT INTO public.order_status (id, name) VALUES (4, 'completed');
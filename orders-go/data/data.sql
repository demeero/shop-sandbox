DROP TABLE IF EXISTS order_item;
DROP TABLE IF EXISTS "order";
DROP TABLE IF EXISTS "order_status";
-- =====================================================================================================================

CREATE TABLE order_status
(
    id   SERIAL NOT NULL,
    name TEXT   NOT NULL,
    CONSTRAINT order_status_pk PRIMARY KEY (id),
    CONSTRAINT order_status_name_key UNIQUE (name)
);

INSERT INTO order_status (id, name)
VALUES (1, 'pending');
INSERT INTO order_status (id, name)
VALUES (2, 'processing');
INSERT INTO order_status (id, name)
VALUES (3, 'canceled');
INSERT INTO order_status (id, name)
VALUES (4, 'completed');

ALTER SEQUENCE order_status_id_seq RESTART WITH 5;
-- =====================================================================================================================

CREATE TABLE "order"
(
    id              UUID      NOT NULL DEFAULT gen_random_uuid(),
    user_id         TEXT      NOT NULL,
    total_units     INTEGER   NOT NULL DEFAULT 0,
    total_nanos     INTEGER   NOT NULL DEFAULT 0,
    contact_name    TEXT      NOT NULL,
    phone           TEXT      NOT NULL,
    city            TEXT      NOT NULL,
    address1        TEXT      NOT NULL,
    address2        TEXT,
    order_status_id INTEGER   NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    CONSTRAINT order_pk PRIMARY KEY (id),
    CONSTRAINT order_order_status_id_fk FOREIGN KEY (order_status_id) REFERENCES order_status (id)
);
CREATE INDEX idx_order_pagination ON "order" (created_at, id);

INSERT INTO "order" (id, user_id, total_units, total_nanos, contact_name, phone, city, address1, address2,
                     order_status_id, created_at)
VALUES ('c9767077-2aa2-4e1a-ad35-87b5d223ef5f', '2', 200, 40, 'Contact Name 21', '+380687333231', 'Odessa', 'NP #40',
        null, 2, '2020-12-21 22:04:45.489884'),
       ('ed6bf8bd-afdc-41b2-b265-abfb4088dee1', '2', 100, 0, 'Contact Name 22', '+380687333232', 'Odessa', 'NP #41',
        null, 1, '2020-12-22 22:04:45.489884'),
       ('dc93a7cd-b3ad-4a65-a8d5-83248a8af324', '1', 300, 0, 'Contact Name 11', '+380687333233', 'Odessa', 'NP #42',
        null, 1, '2020-12-22 22:04:45.489884'),
       ('f149cfdd-13ae-4060-9960-ae2ff23d5ae4', '1', 50, 0, 'Contact Name 11', '+380687333233', 'Odessa', 'NP #42',
        null, 1, '2020-12-23 22:23:45.489884'),
       ('378d275f-b537-4c7c-af7e-1b0cd9372307', '3', 44, 0, 'Contact Name 31', '+380687333234', 'Odessa', 'NP #43',
        null, 1, '2020-12-24 22:24:45.489884');
-- =====================================================================================================================

CREATE TABLE order_item
(
    id           UUID              NOT NULL DEFAULT gen_random_uuid(),
    quantity     INTEGER           NOT NULL,
    total_units  INTEGER DEFAULT 0 NOT NULL,
    total_nanos  INTEGER DEFAULT 0 NOT NULL,
    product_id   TEXT              NOT NULL,
    product_name TEXT              NOT NULL,
    order_id     UUID              NOT NULL,
    CONSTRAINT order_item_pk PRIMARY KEY (id),
    CONSTRAINT order_item_order_fk FOREIGN KEY (order_id) REFERENCES "order"
);

INSERT INTO order_item (id, quantity, total_units, total_nanos, product_id, product_name, order_id)
VALUES ('c2f1e96c-828a-42a5-b541-dde5da62c431', 1, 100, 40, '1', 'Product Name 1',
        'c9767077-2aa2-4e1a-ad35-87b5d223ef5f'),
       ('9e39e4b1-8e95-4be7-8b54-5a5eff27c7c6', 2, 100, 0, '2', 'Product Name 2',
        'c9767077-2aa2-4e1a-ad35-87b5d223ef5f'),

       ('9135e3f2-ac08-4084-87c8-2a95cab8bfac', 1, 25, 0, '3', 'Product Name 3',
        'ed6bf8bd-afdc-41b2-b265-abfb4088dee1'),
       ('83db1c9b-ea2f-47f1-a73d-8f5bfa584be4', 3, 75, 0, '4', 'Product Name 4',
        'ed6bf8bd-afdc-41b2-b265-abfb4088dee1'),

       ('3cc7fabc-5486-4c38-a088-5fa6410fceda', 1, 300, 0, '5', 'Product Name 5',
        'dc93a7cd-b3ad-4a65-a8d5-83248a8af324'),

       ('e16d532a-c4c0-447f-8779-cf4e3b5f280e', 1, 25, 0, '3', 'Product Name 3',
        'f149cfdd-13ae-4060-9960-ae2ff23d5ae4'),
       ('644c07b1-af70-4a99-bee0-1cdbdb55e54d', 1, 10, 0, '6', 'Product Name 6',
        'f149cfdd-13ae-4060-9960-ae2ff23d5ae4'),
       ('73ce14c0-e4ed-4886-aa1c-949e71828217', 1, 15, 0, '7', 'Product Name 7',
        'f149cfdd-13ae-4060-9960-ae2ff23d5ae4'),

       ('c2a4f171-313b-4534-b983-3a0ee43911f3', 44, 44, 0, '111', 'Product Name 111',
        '378d275f-b537-4c7c-af7e-1b0cd9372307');
-- =====================================================================================================================
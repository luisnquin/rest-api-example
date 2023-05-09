

/*
    PRODUCTS
*/

INSERT INTO products (product_id, name, price)
    VALUES ('f7afdebd-2ab7-46e6-832c-229dc6b5d94b', 'Cookie', 1.25),
           ('e7f5bfcf-8292-4655-ace7-f7d77c073c98', 'Orange', 1.55),
           ('f2a8a7a5-fa93-4f0a-a0b3-4bcbfba644c3', 'Milk', 2.25),
           ('087af08d-f66f-4679-9746-d59a39a3640a', 'Strawberry', 2.00);

/*
    CUSTOMERS
*/

INSERT INTO customers (customer_id, first_name, last_name, email, phone_number)
    VALUES ('fce1b59d-910c-4ac3-9e49-7653506fb697', 'John', 'Doe', 'john@doe.com', '1234'),
           ('0b06a186-3853-47fb-bd70-29236c56c183', 'Johanna', 'Doe', 'johanna@doe.com', '2345'),
           ('c7a52f6e-4c10-4b1c-bd45-97f1e2d61b47', 'Jane', 'Doe', 'jane@doe.com', '5678'),
           ('ab3e45bf-7b01-4323-8e3d-1a2b3c4d5e6f', 'Bob', 'Smith', 'bob@smith.com', '9101');

/*
    ORDERS
*/

INSERT INTO orders (order_id, customer_id, created_at, updated_at)
    VALUES ('f5b36c11-7cf3-4d14-90f2-21655abd6126', 'fce1b59d-910c-4ac3-9e49-7653506fb697', NOW(), NOW()),
           ('77926bba-aedf-45f5-8ce3-55107a78c982', 'fce1b59d-910c-4ac3-9e49-7653506fb697', NOW(), NOW()),

           ('385ceac8-d272-48b9-bea9-a2c30611ac65', '0b06a186-3853-47fb-bd70-29236c56c183', NOW(), NOW()),
           ('de815a6b-211e-445c-a9d7-3db93eb71264', '0b06a186-3853-47fb-bd70-29236c56c183', NOW(), NOW()),

           ('ce6db1b3-77c9-4e2d-9745-96d9997d80c5', 'c7a52f6e-4c10-4b1c-bd45-97f1e2d61b47', NOW(), NOW()),
           ('a6508588-f894-4562-8112-54ab70a1c6e5', 'c7a52f6e-4c10-4b1c-bd45-97f1e2d61b47', NOW(), NOW()),

           ('726d283b-4444-430d-aa32-372af265e2e9', 'ab3e45bf-7b01-4323-8e3d-1a2b3c4d5e6f', NOW(), NOW()),
           ('08e2e947-0bad-4f54-98eb-5b4bbc4863f8', 'ab3e45bf-7b01-4323-8e3d-1a2b3c4d5e6f', NOW(), NOW());

/*
    ORDER ITEMS
*/

INSERT INTO order_items (order_item_id, order_id, product_id, quantity)
    VALUES ('776421c5-fc5c-466c-a534-98e2a3aca8f2', 'f5b36c11-7cf3-4d14-90f2-21655abd6126', 'e7f5bfcf-8292-4655-ace7-f7d77c073c98', 50),
           ('cd90f146-e84b-4b2e-936d-e433e58335c2', 'f5b36c11-7cf3-4d14-90f2-21655abd6126', 'f2a8a7a5-fa93-4f0a-a0b3-4bcbfba644c3', 15),
           ('93b9d615-77ec-40c0-9e30-f24d3ea5eac7', 'f5b36c11-7cf3-4d14-90f2-21655abd6126', '087af08d-f66f-4679-9746-d59a39a3640a', 12),

           ('1f470b8a-e7ca-4c68-9d11-cec3248c0a04', '77926bba-aedf-45f5-8ce3-55107a78c982', 'f7afdebd-2ab7-46e6-832c-229dc6b5d94b', 1),
           ('838b5d80-cf2c-415d-9a3e-2dc0c96d9ff8', '77926bba-aedf-45f5-8ce3-55107a78c982', 'e7f5bfcf-8292-4655-ace7-f7d77c073c98', 5),
           ('64a39e6a-3f8d-48c4-a78d-c33835b9f9be', '77926bba-aedf-45f5-8ce3-55107a78c982', 'f2a8a7a5-fa93-4f0a-a0b3-4bcbfba644c3', 12),
           ('846ae65f-d3b3-48ad-992f-7be40adcdfcd', '77926bba-aedf-45f5-8ce3-55107a78c982', '087af08d-f66f-4679-9746-d59a39a3640a', 40),

           ('2b58a405-1ffc-4319-b1ec-ec771a0cfbb0', '385ceac8-d272-48b9-bea9-a2c30611ac65', 'f2a8a7a5-fa93-4f0a-a0b3-4bcbfba644c3', 1),
           ('1086148f-f6cf-4de9-a3c9-25623e8a9e55', '385ceac8-d272-48b9-bea9-a2c30611ac65', 'f7afdebd-2ab7-46e6-832c-229dc6b5d94b', 3),

           ('3487fd5a-dbfc-4132-8433-304bab6759ba', 'de815a6b-211e-445c-a9d7-3db93eb71264', 'f7afdebd-2ab7-46e6-832c-229dc6b5d94b', 7),
           ('df7ed781-38f9-432a-ae20-e578858a23d5', 'de815a6b-211e-445c-a9d7-3db93eb71264', 'f2a8a7a5-fa93-4f0a-a0b3-4bcbfba644c3', 5),
           ('f6249473-e4f3-47ef-83f8-36aefaccd42f', 'de815a6b-211e-445c-a9d7-3db93eb71264', '087af08d-f66f-4679-9746-d59a39a3640a', 9),

           ('4ad08d13-1cbf-4032-b3b2-992a63fab46c', 'ce6db1b3-77c9-4e2d-9745-96d9997d80c5', '087af08d-f66f-4679-9746-d59a39a3640a', 10),
           ('c589f4c3-2130-4a32-9423-d833d1bbf6f6', 'ce6db1b3-77c9-4e2d-9745-96d9997d80c5', 'f2a8a7a5-fa93-4f0a-a0b3-4bcbfba644c3', 10),

           ('282bbf83-9458-41cd-b87e-7c813c0f6574', 'a6508588-f894-4562-8112-54ab70a1c6e5', '087af08d-f66f-4679-9746-d59a39a3640a', 5),
           ('afe1a6c4-0368-4964-951b-c7cdb0f9722b', 'a6508588-f894-4562-8112-54ab70a1c6e5', 'e7f5bfcf-8292-4655-ace7-f7d77c073c98', 30),
           ('9a528975-7b68-42a4-918a-f578eb221b1f', 'a6508588-f894-4562-8112-54ab70a1c6e5', 'f2a8a7a5-fa93-4f0a-a0b3-4bcbfba644c3', 45),

           ('86036a14-12c8-47bc-8511-c3a29908eaa0', '726d283b-4444-430d-aa32-372af265e2e9', 'f7afdebd-2ab7-46e6-832c-229dc6b5d94b', 5),
           ('b00a973e-b9b0-4e5b-bb18-fe51c612cd09', '726d283b-4444-430d-aa32-372af265e2e9', 'e7f5bfcf-8292-4655-ace7-f7d77c073c98', 2),
           ('d01b22af-909e-4057-93cf-a61f62bac7fb', '726d283b-4444-430d-aa32-372af265e2e9', 'f2a8a7a5-fa93-4f0a-a0b3-4bcbfba644c3', 3),
           ('967a9c04-6b9c-4253-9448-537a252d2790', '726d283b-4444-430d-aa32-372af265e2e9', '087af08d-f66f-4679-9746-d59a39a3640a', 4),

           ('b1c9ae6e-03fd-46cf-a1c8-8c531545bf88', '08e2e947-0bad-4f54-98eb-5b4bbc4863f8', '087af08d-f66f-4679-9746-d59a39a3640a', 10),
           ('e31db058-550a-4250-874e-fdbe58b358fa', '08e2e947-0bad-4f54-98eb-5b4bbc4863f8', 'e7f5bfcf-8292-4655-ace7-f7d77c073c98', 5);

create table products (
id serial Primary key, -- Уникальный идентификатор товара
name Text not null,
description text ,
price INT,
category text ,
picture text 
) 



INSERT INTO products(      -- куда будем вставлять данные
  name,
  price,
  category-- имена полей 
)
VALUES                -- перечисляем вставляемые строки
  ('Латте', 220, 'Кофе'),
  ('Капучино', 230, 'Кофе'),
  ('Зеленый чай ', 180, 'чай'),
  ('Пуэр', 300, 'Чай'),
  ('Лимонад', 300, 'прочее')
 
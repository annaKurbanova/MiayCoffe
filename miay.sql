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
 
//создаем таблицу корзина 
CREATE TABLE unified_cart (
    id SERIAL PRIMARY KEY,         -- Уникальный идентификатор корзины
    user_id INT NOT NULL,          -- ID пользователя
    products JSONB,                -- Массив товаров в корзине
    notes TEXT,                    -- Дополнительные заметки
    status VARCHAR(50) DEFAULT 'Active', -- Значение по умолчанию  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата создания
    total DECIMAL(10, 2)           -- Общая сумма корзины
);


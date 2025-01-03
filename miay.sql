CREATE USER user_name with PASSWORD 'user_pass' SUPERUSER;



create table products (
id serial Primary key, 
name Text not null,
description text ,
price INT,
category text ,
picture text 
);



INSERT INTO products( 
  name,
  price,
  picture,
  category
)
VALUES     
  
    ( 'Капучино', 
    200,
   'https://img.freepik.com/premium-photo/italian-cappuccino-with-decoration_159805-806.jpg?w=826',  
   'coffee'
    ),
    ('Латте',
    220,
    'https://img.freepik.com/premium-photo/italian-cappuccino-with-decoration_159805-806.jpg?w=826',  
    'coffee'
     ),
     (
     'Эспрессо',
     100 ,
    'https://img.freepik.com/premium-photo/italian-cappuccino-with-decoration_159805-806.jpg?w=826',  
     'coffee'
     ),
    ( 
     'Раф',
     300 ,
    'https://img.freepik.com/premium-photo/italian-cappuccino-with-decoration_159805-806.jpg?w=826',  
     'coffee'
    ),
    (
     'Ассам',
     100 ,
    'https://www.chay.info/upload/medialibrary/146/3i03anmkw7cwyvo0v96xi857k1pxqi8a.jpeg',
     'tea'
    ),
    (
     'Сенча',
     100 ,
    'https://www.chay.info/upload/medialibrary/146/3i03anmkw7cwyvo0v96xi857k1pxqi8a.jpeg',
     'tea'
    ),
    (
     'Травянной',
     100 ,
    'https://www.chay.info/upload/medialibrary/146/3i03anmkw7cwyvo0v96xi857k1pxqi8a.jpeg',
      'tea'
    ),
  (
     'Ягодный',
     150 ,
    'https://www.chay.info/upload/medialibrary/146/3i03anmkw7cwyvo0v96xi857k1pxqi8a.jpeg',
      'tea'
  ),
  (
     'Облепиха',
     150 ,
 'https://www.chay.info/upload/medialibrary/146/3i03anmkw7cwyvo0v96xi857k1pxqi8a.jpeg',
      'tea'
  ),
  ( 
     'Лайм - мята',
     200 ,
    'https://www.ethnomir.ru/upload/medialibrary/08e/limonad_1140.jpg',
     'lemonades'
  ),
  ( 
     'Малина - базилик',
     200 ,
 'https://www.ethnomir.ru/upload/medialibrary/08e/limonad_1140.jpg',
      'lemonades'
  ),
  ( 
     'Личи - апельсин',
     200 ,
 'https://www.ethnomir.ru/upload/medialibrary/08e/limonad_1140.jpg', 
     'lemonades'
  ),
  ( 
     'Манго - маракуйа',
     200 ,
 'https://www.ethnomir.ru/upload/medialibrary/08e/limonad_1140.jpg',
      'lemonades'
  ),
  ( 
     'Чизкейк',
     250 ,
    'https://www.calend.ru/img/Articles/2019/desert.jpg',
     'deserts'
  ),
  (
     'Тирамису',
     230 ,
 'https://www.calend.ru/img/Articles/2019/desert.jpg',
      'deserts'
  ),
  (
    ' Наполеон',
     210 ,
 'https://www.calend.ru/img/Articles/2019/desert.jpg',
      'deserts'
  ),
  (
     'Анна-Павлова',
     260 ,
 'https://www.calend.ru/img/Articles/2019/desert.jpg',
      'deserts'
  );

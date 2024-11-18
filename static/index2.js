// Получаем ссылки на элементы модального окна регистрации и входа
const button_modal_registration = document.querySelector('.button_modal_registration'); // Кнопка для открытия модалки регистрации
const modal_window_registration = document.querySelector('.modal_window_registration'); // Модальное окно регистрации
const button_registration = document.querySelector('.button_registration'); // Кнопка подтверждения регистрации

const button_modal_entrance = document.querySelector('.button_modal_entrance'); // Кнопка для открытия модалки входа
const modal_window_entrance = document.querySelector('.modal_window_entrance'); // Модальное окно входа
const button_entrance = document.querySelector('.button_entrance'); // Кнопка подтверждения входа

const name_user = document.querySelector('#name_user'); // Поле ввода имени пользователя регистрации
const telephone_user = document.querySelector('#telephone_user'); // Поле ввода номера телефона пользователя регистрации
const name_user2 = document.querySelector('#name_user2');// Поле ввода имени пользователя входа
const telephone_user2 = document.querySelector('#telephone_user2');// Поле ввода номера телефона пользователя входа

const button_menu = document.querySelector(".button_menu");
const fast_menu_modal = document.querySelector(".fast_menu_modal");
const close_modal = document.querySelector('.close_modal');


button_menu.onclick = function(){
    fast_menu_modal.classList.add('active_modal');
}
close_modal.onclick = function(){
    fast_menu_modal.classList.remove('active_modal');
}
// Функция для открытия и закрытия модального окна регистрации
button_modal_registration.onclick = function () {
    button_registration.style.color = "black";
    name_user.value=""
    telephone_user.value=""
    modal_window_registration.classList.toggle('active'); // Показать/скрыть модалку регистрации
    modal_window_entrance.classList.remove('active'); // Закрыть модалку входа, если она открыта
};

// Функция для открытия и закрытия модального окна входа
button_modal_entrance.onclick = function () {
    button_entrance.style.color = "black";
    name_user2.value=""
    telephone_user2.value=""
    modal_window_entrance.classList.toggle('active'); // Показать/скрыть модалку входа
    modal_window_registration.classList.remove('active'); // Закрыть модалку регистрации, если она открыта

};

// Асинхронная функция для отправки POST-запроса на сервер
async function sendRequest(url, data) {
    try {
        const response = await fetch(url, {
            method: "POST", // Указываем, что используем метод POST
            headers: {
                "Content-Type": "application/json", // Устанавливаем заголовок, чтобы сервер знал, что данные в формате JSON
            },
            body: JSON.stringify(data), // Преобразуем объект data в строку JSON и добавляем в тело запроса
        });
        
        console.log("Сырой ответ:", JSON.stringify(response)); // Логируем ответ в сыром виде
        console.log("Статус успешности:", response.ok); // Проверяем, прошел ли запрос успешно

        if (response.ok) { // Проверяем, вернул ли сервер успешный статус
            const responseData = JSON.stringify(response); // Преобразуем ответ в JSON, если он успешный
            console.log("Запрос успешно отправлен", responseData); // Логируем данные, полученные от сервера
        } else {
            console.error("Ошибка при отправке запроса:", response.statusText); // Логируем ошибку, если запрос не успешен
        }
    } catch (error) {
        console.error("Ошибка соединения или отправки запроса:", error); // Логируем ошибку при соединении
    }
}

// Обработчик кнопки регистрации
button_registration.addEventListener("click", () => {
    const bodyData = {
        Name: name_user.value, // Считываем имя пользователя из поля ввода
        PhoneNumber: telephone_user.value, // Считываем номер телефона из поля ввода
    };
    if (!name_user.value || !telephone_user.value) {//Проверка на заполнение полей
        button_registration.style.color = "red";
        return;
    }else{
        modal_window_registration.classList.remove('active');//Закрываем модалку после отправки запроса
    }
    sendRequest("http://localhost:1324/registration", bodyData); // Отправляем запрос с данными на сервер
});

// Обработчик кнопки входа
button_entrance.addEventListener("click", () => {
    const bodyData = {
        Name: name_user2.value, // Считываем имя пользователя из поля ввода
        PhoneNumber: telephone_user2.value, // Считываем номер телефона из поля ввода
    };
    if (!name_user2.value || !telephone_user2.value) {//Проверка на заполнение полей
        button_entrance.style.color="red";
        return;
    }else{
        modal_window_entrance.classList.remove('active');//Закрываем модалку после отправки запроса
    }
    sendRequest("http://localhost:1324/Vhod", bodyData); // Отправляем запрос с данными на сервер
    
   
    
});

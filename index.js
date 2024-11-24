document.addEventListener('DOMContentLoaded', function () {
    const modalRegistration = document.querySelector('.modal_window_registration');
    const modalEntrance = document.querySelector('.modal_window_entrance');
    const btnRegistration = document.querySelector('.button_modal_registration');
    const btnEntrance = document.querySelector('.button_modal_entrance');
    const closeButtons = document.querySelectorAll('.close-modal');

    // Открытие модального окна регистрации
    btnRegistration.addEventListener('click', () => {
        modalRegistration.classList.add('active');
    });

    // Открытие модального окна входа
    btnEntrance.addEventListener('click', () => {
        modalEntrance.classList.add('active');
    });

    // Закрытие модальных окон при нажатии на кнопки "закрыть"
    closeButtons.forEach(button => {
        button.addEventListener('click', () => {
            modalRegistration.classList.remove('active');
            modalEntrance.classList.remove('active');
        });
    });

    // Закрытие модальных окон при клике за пределами окна
    window.addEventListener('click', (e) => {
        if (e.target === modalRegistration) {
            modalRegistration.classList.remove('active');
        }
        if (e.target === modalEntrance) {
            modalEntrance.classList.remove('active');
        }
    });
});

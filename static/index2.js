const button_modal_registration = document.querySelector('.button_modal_registration')
const modal_window_registration = document.querySelector('.modal_window_registration')
const button_registration = document.querySelector('.button_registration')
const button_modal_entrance = document.querySelector('.button_modal_entrance')
const modal_window_entrance = document.querySelector('.modal_window_entrance')
const button_entrance = document.querySelector('.button_entrance')
const name_user = document.querySelector('#name_user')
const telephone_user = document.querySelector('#telephone_user')

button_modal_registration.onclick = function(){
modal_window_registration.classList.toggle('active')
modal_window_entrance.classList.remove('active')
}

button_modal_entrance.onclick = function(){
  modal_window_entrance.classList.toggle('active')
  modal_window_registration.classList.remove('active')
  }

  async function sendRequest(url, data) {
    try {
        const response = await fetch(url, {
          mode:"no-cors",
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });
 
        if (response.ok) {
            const responseData = await response.json(); 
            console.log("Запрос успешно отправлен", responseData);
        } else {
            console.error("Ошибка при отправке запроса:", response.statusText);
        }
    } catch (error) {
        console.error("Ошибка соединения или отправки запроса:", error);
    }
 }
 
 button_registration.addEventListener("click", () => {
    const bodyData = {
        Name: name_user.value,
        PhoneNumber: telephone_user.value,
    };
    sendRequest("http://localhost:1324/registration", bodyData);
 });
 
 button_entrance.addEventListener("click", () => {
    const bodyData = {
        Name: name_user.value,
        PhoneNumber: telephone_user.value,
    };
    sendRequest("http://localhost:1324/registration", bodyData);
 });
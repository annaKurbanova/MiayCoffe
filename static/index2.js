const button_modal = document.querySelector('.button_modal')
const modal_window = document.querySelector('.modal_window')
const button_registration = document.querySelector('.button_registration')

console.log("registration handle")

button_modal.onclick = function(){
modal_window.style.display='block'
}

button_registration.addEventListener("click", async () => {
  
  const bodyData = JSON.stringify({
    Name: name_user.value,
    PhoneNumber: telephone_user.value,  
  });

  try {
    const responce = await fetch("http://localhost:1324/registration", {
      mode: "no-cors",
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: bodyData,
    });

    if (responce.ok) {
      console.log("Запрос успешно отправлен");
      console.log(responce)
    } else {
      console.error("Ошибка при отправке запроса:", responce.statusText);
    }
  } catch (error) {
    console.error("Ошибка соединения или отправки запроса:", error);
  }
});
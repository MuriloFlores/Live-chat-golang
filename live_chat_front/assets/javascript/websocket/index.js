import createMessageContainer from './../elements/message_container.js';
import renderConectedUserList from '../elements/connected_users_list.js';
const chat = document.getElementById("chat_record");
let nickname = localStorage.getItem("nickname")
const connection = new WebSocket("ws://localhost:8080/chat/" + nickname)
document.title = "Olá, " + nickname + " Bem-Vindo"

connection.onmessage = (e) => {
    let data = JSON.parse(e.data)
    chat.appendChild(createMessageContainer(nickname, data))
    renderConectedUserList()
}

connection.onopen = (e) => {
    renderConectedUserList()
}

document.getElementById("sender").addEventListener("click", () => {
    let content = document.getElementById("message_content");

    let message = {
        "from": nickname,
        "content": content.value,
    }

    connection.send(JSON.stringify(message))
    content.value = ""
    renderConectedUserList() 
})

document.querySelector(".logout").addEventListener("click", () => {
    localStorage.clear()
    window.location.href = './index.html'
})


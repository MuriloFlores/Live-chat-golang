let nickname = ""

document.querySelector(".btn").addEventListener("click", () => {
    nickname = document.querySelector(".nickname").value
    localStorage.setItem("nickname", nickname)
})

document.querySelector('form').addEventListener('submit', function(event) {
    event.preventDefault();
    window.location.href = './chat.html?nickname=' + encodeURIComponent(nickname)
});



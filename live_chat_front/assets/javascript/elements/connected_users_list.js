export default function renderConectedUserList() {
    fetch("http://localhost:8080/listUsers")
        .then(response => response.json())
        .then(data => {
            // Encontre o elemento 'ul' no qual você deseja adicionar os itens da lista
            const ul = document.querySelector('.connected_users');

            ul.innerHTML = '';

            for (const key in data) {
                let li = document.createElement('li');
                li.classList.add("p-3", "flex", "flex-row", "justify-start", "items-center");

                const div = document.createElement('div');
                div.classList.add("bg-green-500", "w-2", "h-2", "rounded-full", "m-2");

                const p = document.createElement("p");
                p.innerHTML = data[key].nickname;

                li.appendChild(div);
                li.appendChild(p);

                // Adicione o elemento 'li' ao elemento 'ul'
                ul.appendChild(li);
            }
        })
        .catch(error => console.error('Erro:', error));
}

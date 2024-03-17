export default function createMessageContainer(nickname, data) {
    const icon = document.createElement("div")
    let containerMessage;

    icon.classList.add("bg-zinc-300", "w-9", "h-9", "rounded-full", "text-center", "pt-1", "font-semibold")

    const message = document.createElement("p")
    message.classList.add("bg-zinc-300", "rounded-lg", "p-2", "text-start", "max-w-lg")

    if (data.from == nickname) {
        containerMessage = document.createElement("div")
        containerMessage.classList.add("flex",
            "justify-start", "items-center", "gap-5", "w-full", "pt-2", "flex-row-reverse")
    } else {
        containerMessage = document.createElement("div")
        containerMessage.classList.add("flex",
            "justify-start", "items-center", "gap-5", "w-full", "pt-2")
    }

    let iconString = data.from.substring(0, 2).toUpperCase()
    icon.innerHTML = iconString


    message.innerHTML = data.content
    
    containerMessage.appendChild(icon)
    containerMessage.appendChild(message)

    return containerMessage

}
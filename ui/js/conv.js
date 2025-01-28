document.addEventListener("DOMContentLoaded", () => {
    const userList = document.getElementById("userList");
    const conversationModal = document.getElementById("conversationModal");
    const closeModal = document.querySelector(".close-modal");
    const messagesDiv = document.getElementById("messages");
    const messageForm = document.getElementById("messageForm");
    const messageInput = document.getElementById("messageInput");
  
    // Dummy user data
    const users = [
      { id: 1, name: "Alice", status: "online" },
      { id: 2, name: "Bob", status: "offline" },
      { id: 3, name: "Charlie", status: "online" },
    ];
  
    // Display users
    function renderUsers() {
      userList.innerHTML = "";
      users.forEach((user) => {
        const userItem = document.createElement("li");
        userItem.textContent = user.name;
        userItem.classList.add(user.status);
        userItem.dataset.userId = user.id;
        userList.appendChild(userItem);
      });
    }
  
    // Open conversation
    userList.addEventListener("click", (e) => {
      if (e.target.tagName === "LI") {
        const userId = e.target.dataset.userId;
        const user = users.find((u) => u.id == userId);
        if (user) {
          openConversation(user);
        }
      }
    });
  
    function openConversation(user) {
      conversationModal.style.display = "flex";
      messagesDiv.innerHTML = `<p>Conversation with ${user.name}</p>`;
    }
  
    closeModal.addEventListener("click", () => {
      conversationModal.style.display = "none";
    });
  
    // Send message
    messageForm.addEventListener("submit", (e) => {
      e.preventDefault();
      const message = messageInput.value.trim();
      if (message) {
        const newMessage = document.createElement("p");
        newMessage.textContent = `You: ${message}`;
        messagesDiv.appendChild(newMessage);
        messageInput.value = "";
      }
    });
  
    renderUsers();
  });
  
function createcomment(Comment) {
    // Create the main coment container
    const coment = document.createElement('div');
    coment.classList.add('coment');
  
    // Create the user info section
    const userInfo = document.createElement('div');
    userInfo.classList.add('user-info');
  
    const avatar = document.createElement('img');
    avatar.src = '/ui/css/default-profile.jpg';
    avatar.alt = 'User Avatar';
    avatar.classList.add('avatar');
  
    const userDetails = document.createElement('div');
    userDetails.classList.add('user-details');
  
    const username = document.createElement('h4');
    username.classList.add('username');
    username.textContent = Comment.author;
  
    const timestamp = document.createElement('p');
    timestamp.classList.add('timestamp');
    timestamp.textContent = Comment.date;
  
    userDetails.appendChild(username);
    userDetails.appendChild(timestamp);
    userInfo.appendChild(avatar);
    userInfo.appendChild(userDetails);
  
    // Create the coment content
    const comentContent = document.createElement('p');
    comentContent.classList.add('coment-content');
    comentContent.textContent = Comment.content;
  
    // Create the coment actions section
    const comentActions = document.createElement('div');
    comentActions.classList.add('coment-actions');
  
    // Like button and notification
    const like = document.createElement('div');
    like.classList.add('like');
  
    const likeButton = document.createElement('button');
    likeButton.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="#707C97"><path d="M720-144H264v-480l288-288 32 22q17 12 26 30.5t5 38.5l-1 5-38 192h264q30 0 51 21t21 51v57q0 8-1.5 14.5T906-467L786.93-187.8Q778-168 760-156t-40 12Zm-384-72h384l120-279v-57H488l49-243-201 201v378Zm0-378v378-378Zm-72-30v72H120v336h144v72H48v-480h216Z"/></svg>`;
    const likeNotification = document.createElement('span');
    likeNotification.classList.add('notification-icon');
    likeNotification.textContent = Comment.likes;
  
    like.appendChild(likeButton);
    like.appendChild(likeNotification);
  
    // Dislike button and notification
    const dislike = document.createElement('div');
    dislike.classList.add('dislike');
    const dislikeButton = document.createElement('button');
    dislikeButton.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" height="20px" viewBox="0 -960 960 960" width="20px" fill="#707C97"><path d="M240-816h456v480L408-48l-32-22q-17-12-26-30.5t-5-38.5l1-5 38-192H120q-30 0-51-21t-21-51v-57q0-8 1.5-14.5T54-493l119-279q8-20 26.5-32t40.5-12Zm384 72H240L120-465v57h352l-49 243 201-201v-378Zm0 378v-378 378Zm72 30v-72h144v-336H696v-72h216v480H696Z"/></svg>`;
    const dislikeNotification = document.createElement('span');
    dislikeNotification.classList.add('notification-icon');
    dislikeNotification.textContent = Comment.dislikes;
  
    dislike.appendChild(dislikeButton);
    dislike.appendChild(dislikeNotification);
  
    // Append all coment actions to the coment-actions container
    comentActions.appendChild(like);
    comentActions.appendChild(dislike);
    // Append all sections to the coment
    coment.appendChild(userInfo);
    coment.appendChild(comentContent);
    coment.appendChild(comentActions);
    coment.id = Comment.id
    let container = document.querySelector(".comments-section")
    container.appendChild(coment)
}
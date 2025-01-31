var info = {}
await getInfoData().then(i =>{
  info = i 
})
let num = 0
let loading = false; 
let isSubmitting = false;
function createSidebar(user) {
    // Create the main sidebar container
    const sidebar = document.createElement('div');
    sidebar.classList.add('sidebar');
  
    // Create the profile section
    const profileSection = document.createElement('div');
    profileSection.classList.add('profile-section');
  
    const profilePic = document.createElement('img');
    profilePic.id = 'profile-pic';
    profilePic.src = '/ui/css/default-profile.jpg';
    profilePic.alt = 'Profile Picture';
  
    const userName = document.createElement('h3');
    userName.id = 'user-name';
    userName.textContent = user.username;
  
    profileSection.appendChild(profilePic);
    profileSection.appendChild(userName);
  
    // Create the menu
    const menu = document.createElement('nav');
    menu.classList.add('menu');
  
    // Menu items data
    const menuItems = [
      {
        text: 'Home',
        svgPath: 'M240-200h120v-240h240v240h120v-360L480-740 240-560v360Zm-80 80v-480l320-240 320 240v480H520v-240h-80v240H160Zm320-350Z',
      },
      {
        text: 'Chat',
        svgPath: 'M240-400h320v-80H240v80Zm0-120h480v-80H240v80Zm0-120h480v-80H240v80ZM80-80v-720q0-33 23.5-56.5T160-880h640q33 0 56.5 23.5T880-800v480q0 33-23.5 56.5T800-240H240L80-80Zm126-240h594v-480H160v525l46-45Zm-46 0v-480 480Z',
      },
      {
        text: 'Profile',
        svgPath: 'M480-480q-66 0-113-47t-47-113q0-66 47-113t113-47q66 0 113 47t47 113q0 66-47 113t-113 47ZM160-160v-112q0-34 17.5-62.5T224-378q62-31 126-46.5T480-440q66 0 130 15.5T736-378q29 15 46.5 43.5T800-272v112H160Zm80-80h480v-32q0-11-5.5-20T700-306q-54-27-109-40.5T480-360q-56 0-111 13.5T260-306q-9 5-14.5 14t-5.5 20v32Zm240-320q33 0 56.5-23.5T560-640q0-33-23.5-56.5T480-720q-33 0-56.5 23.5T400-640q0 33 23.5 56.5T480-560Zm0-80Zm0 400Z',
      },
      {
        text: 'Log out',
        svgPath: 'M479.88-478.67q-14.21 0-23.71-9.58t-9.5-23.75v-337.33q0-14.17 9.61-23.75 9.62-9.59 23.84-9.59 14.21 0 23.71 9.59 9.5 9.58 9.5 23.75V-512q0 14.17-9.61 23.75-9.62 9.58-23.84 9.58Zm.12 360q-75 0-140.5-28.5t-114-77q-48.5-48.5-77-114T120-478.67q0-63 21.67-121.83 21.66-58.83 62.33-106.83 9.67-11.34 24-11.5 14.33-.17 25.17 10.66 9.16 9.17 7.83 22.84-1.33 13.66-10 25-31.67 38-48 85.15-16.33 47.16-16.33 96.51 0 122.57 85.38 207.96 85.38 85.38 207.95 85.38t207.95-85.38q85.38-85.39 85.38-207.96 0-50.66-16.16-97.5-16.17-46.83-48.5-85.5-8.89-11.03-9.78-24.18Q698-699 707-708q10.67-10.67 25.67-10.17 15 .5 24.66 12.17 41 48 61.84 106.33 20.83 58.34 20.83 121 0 75-28.5 140.5t-77 114q-48.5 48.5-114 77T480-118.67Z',
        className: 'LO',
      },
    ];
    
    menuItems.forEach((item) => {
      const menuItem = document.createElement('a');
      menuItem.href = '#';
      menuItem.classList.add('menu-item');
      if (item.className) menuItem.classList.add(item.className);
  
      const button = document.createElement('button');
      button.classList.add('btn');
  
      const svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
      svg.setAttribute('xmlns', 'http://www.w3.org/2000/svg');
      svg.setAttribute('height', '40');
      svg.setAttribute('viewBox', '0 -960 960 960');
      svg.setAttribute('width', '40');
      svg.setAttribute('fill', '#707C97');
  
      const path = document.createElementNS('http://www.w3.org/2000/svg', 'path');
      path.setAttribute('d', item.svgPath);
  
      const linearGradient = document.createElementNS('http://www.w3.org/2000/svg', 'linearGradient');
      linearGradient.setAttribute('id', 'hoverGradient');
      linearGradient.setAttribute('x1', '0%');
      linearGradient.setAttribute('y1', '0%');
      linearGradient.setAttribute('x2', '100%');
      linearGradient.setAttribute('y2', '100%');
  
      const stop1 = document.createElementNS('http://www.w3.org/2000/svg', 'stop');
      stop1.setAttribute('offset', '0%');
      stop1.setAttribute('stop-color', '#7CB8F7');
  
      const stop2 = document.createElementNS('http://www.w3.org/2000/svg', 'stop');
      stop2.setAttribute('offset', '93%');
      stop2.setAttribute('stop-color', '#2A8BF2');
  
      linearGradient.appendChild(stop1);
      linearGradient.appendChild(stop2);
      svg.appendChild(path);
      svg.appendChild(linearGradient);
      button.appendChild(svg);
      menuItem.appendChild(button);
      menuItem.append(` ${item.text}`);
      menu.appendChild(menuItem);
    });
  
    // Append profile section and menu to the sidebar
    sidebar.appendChild(profileSection);
    sidebar.appendChild(menu);
  
    // Append the sidebar to the body
    document.querySelector(".container").appendChild(sidebar);
    let logoutbtn = document.querySelector(".LO")
    logoutbtn.addEventListener('click',function(){
  logout()
})
  };
  document.querySelectorAll('.like, .dislike, .comment').forEach((element) => {
    element.addEventListener('click', () => {
      element.classList.toggle('active');
    });
  });
  function Removesidebar(){
            let sidebar = document.querySelector(".sidebar")
            sidebar.remove()
  }
function Removecard(){
  let card = document.querySelector(".Form") || null
  if (card != null){
    card.remove()
  }
}
function createCard() {
  // Create the main card container
  const card = document.createElement('div');
  card.classList.add('card');

  // Create the front side (Login Form)
  const frontSide = document.createElement('div');
  frontSide.classList.add('card-side', 'front');

  const loginHeading = document.createElement('h2');
  loginHeading.textContent = 'Login';

  const loginForm = document.createElement('form');
  loginForm.id = 'login-form';

  const loginIdLabel = document.createElement('label');
  loginIdLabel.setAttribute('for', 'login-id');
  loginIdLabel.textContent = 'Nickname or E-mail';

  const loginIdInput = document.createElement('input');
  loginIdInput.type = 'text';
  loginIdInput.id = 'login-id';
  loginIdInput.name = 'login-id';
  loginIdInput.required = true;

  const loginPasswordLabel = document.createElement('label');
  loginPasswordLabel.setAttribute('for', 'login-password');
  loginPasswordLabel.textContent = 'Password';

  const loginPasswordInput = document.createElement('input');
  loginPasswordInput.type = 'password';
  loginPasswordInput.id = 'login-password';
  loginPasswordInput.name = 'login-password';
  loginPasswordInput.required = true;

  const loginButton = document.createElement('button');
  loginButton.type = 'submit';
  loginButton.classList.add('btn');
  loginButton.textContent = 'Login';

  const switchToRegister = document.createElement('p');
  switchToRegister.classList.add('switch');
  switchToRegister.innerHTML = 'Don\'t have an account? <span id="switch-to-register">Register</span>';

  loginForm.appendChild(loginIdLabel);
  loginForm.appendChild(loginIdInput);
  loginForm.appendChild(loginPasswordLabel);
  loginForm.appendChild(loginPasswordInput);
  loginForm.appendChild(loginButton);
  loginForm.appendChild(switchToRegister);

  frontSide.appendChild(loginHeading);
  frontSide.appendChild(loginForm);

  // Create the back side (Register Form)
  const backSide = document.createElement('div');
  backSide.classList.add('card-side', 'back');

  const registerHeading = document.createElement('h2');
  registerHeading.textContent = 'Register';

  const registerForm = document.createElement('form');
  registerForm.classList.add('register-form');

  const nicknameLabel = document.createElement('label');
  nicknameLabel.setAttribute('for', 'nickname');
  nicknameLabel.textContent = 'Nickname';

  const nicknameInput = document.createElement('input');
  nicknameInput.type = 'text';
  nicknameInput.id = 'nickname';
  nicknameInput.name = 'nickname';
  nicknameInput.required = true;

  const ageLabel = document.createElement('label');
  ageLabel.setAttribute('for', 'age');
  ageLabel.textContent = 'Age';

  const ageInput = document.createElement('input');
  ageInput.type = 'number';
  ageInput.id = 'age';
  ageInput.name = 'age';
  ageInput.required = true;

  const genderLabel = document.createElement('label');
  genderLabel.setAttribute('for', 'gender');
  genderLabel.textContent = 'Gender';

  const genderSelect = document.createElement('select');
  genderSelect.id = 'gender';
  genderSelect.name = 'gender';
  genderSelect.required = true;

  const maleOption = document.createElement('option');
  maleOption.value = 'male';
  maleOption.textContent = 'Male';

  const femaleOption = document.createElement('option');
  femaleOption.value = 'female';
  femaleOption.textContent = 'Female';

  genderSelect.appendChild(maleOption);
  genderSelect.appendChild(femaleOption);

  const firstNameLabel = document.createElement('label');
  firstNameLabel.setAttribute('for', 'first-name');
  firstNameLabel.textContent = 'First Name';

  const firstNameInput = document.createElement('input');
  firstNameInput.type = 'text';
  firstNameInput.id = 'first-name';
  firstNameInput.name = 'first-name';
  firstNameInput.required = true;

  const lastNameLabel = document.createElement('label');
  lastNameLabel.setAttribute('for', 'last-name');
  lastNameLabel.textContent = 'Last Name';

  const lastNameInput = document.createElement('input');
  lastNameInput.type = 'text';
  lastNameInput.id = 'last-name';
  lastNameInput.name = 'last-name';
  lastNameInput.required = true;

  const emailLabel = document.createElement('label');
  emailLabel.setAttribute('for', 'email');
  emailLabel.textContent = 'E-mail';

  const emailInput = document.createElement('input');
  emailInput.type = 'email';
  emailInput.id = 'email';
  emailInput.name = 'email';
  emailInput.required = true;

  const passwordLabel = document.createElement('label');
  passwordLabel.setAttribute('for', 'password');
  passwordLabel.textContent = 'Password';

  const passwordInput = document.createElement('input');
  passwordInput.type = 'password';
  passwordInput.id = 'password';
  passwordInput.name = 'password';
  passwordInput.required = true;

  const registerButton = document.createElement('button');
  registerButton.type = 'submit';
  registerButton.classList.add('btn');
  registerButton.textContent = 'Register';

  const switchToLogin = document.createElement('p');
  switchToLogin.classList.add('switch');
  switchToLogin.innerHTML = 'Already have an account? <span id="switch-to-login">Login</span>';

  registerForm.appendChild(nicknameLabel);
  registerForm.appendChild(nicknameInput);
  registerForm.appendChild(ageLabel);
  registerForm.appendChild(ageInput);
  registerForm.appendChild(genderLabel);
  registerForm.appendChild(genderSelect);
  registerForm.appendChild(firstNameLabel);
  registerForm.appendChild(firstNameInput);
  registerForm.appendChild(lastNameLabel);
  registerForm.appendChild(lastNameInput);
  registerForm.appendChild(emailLabel);
  registerForm.appendChild(emailInput);
  registerForm.appendChild(passwordLabel);
  registerForm.appendChild(passwordInput);
  registerForm.appendChild(registerButton);
  registerForm.appendChild(switchToLogin);

  backSide.appendChild(registerHeading);
  backSide.appendChild(registerForm);

  // Append front and back sides to the card
  card.appendChild(frontSide);
  card.appendChild(backSide);
  let container = document.querySelector(".container")
  // Append the card to the page
  let form = document.createElement("div")
  form.classList.add("Form")
  form.appendChild(card)
  container.appendChild(form);
switchToRegister.addEventListener('click', () => {
  card.classList.add('flipped');
});

switchToLogin.addEventListener('click', () => {
  card.classList.remove('flipped');
});
  return card;
}
function createcomment(Comment, container) {
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
  container.appendChild(coment)
}
function postin(){
  const postDiv = document.createElement('div');
  postDiv.classList.add('post', 'beta');

  const userInfoDiv = document.createElement('div');
  userInfoDiv.classList.add('user-info');
  
  const avatarImg = document.createElement('img');
  avatarImg.src ='/ui/css/default-profile.jpg';
  avatarImg.alt = 'User Avatar';
  avatarImg.classList.add('avatar');

  const userDetailsDiv = document.createElement('div');
  userDetailsDiv.classList.add('user-details');
  
  const usernameH4 = document.createElement('h4');
  usernameH4.classList.add('username');
  usernameH4.textContent = "Add Your Own Post!";

  userDetailsDiv.appendChild(usernameH4);
  userInfoDiv.appendChild(avatarImg);
  userInfoDiv.appendChild(userDetailsDiv);

  const postContentTextarea = document.createElement('textarea');
  postContentTextarea.classList.add('post-content', 'input');
  postContentTextarea.placeholder = 'Enter your text here...';

  const categoriesDiv = document.createElement('div');
  categoriesDiv.classList.add('categories');

  const categories = ['football', 'cars', 'ronaldo'];
  categories.forEach(category => {
    const categoryLabel = document.createElement('label');
    categoryLabel.classList.add('categorie');
    
    const categoryInput = document.createElement('input');
    categoryInput.type = 'checkbox';
    categoryInput.name = category;
    categoryInput.value = category;

    categoryLabel.appendChild(categoryInput);
    categoryLabel.appendChild(document.createTextNode(category.charAt(0).toUpperCase() + category.slice(1)));

    categoriesDiv.appendChild(categoryLabel);
  });

  const addPostDiv = document.createElement('div');
  addPostDiv.classList.add('addpost');

  const addPostButton = document.createElement('button');
  addPostButton.textContent = 'Post-It!';

  addPostDiv.appendChild(addPostButton);

  postDiv.appendChild(userInfoDiv);
  postDiv.appendChild(postContentTextarea);
  postDiv.appendChild(categoriesDiv);
  postDiv.appendChild(addPostDiv);

  return postDiv;
}
function createPost(Post) {
  // Create the main post container
  const post = document.createElement('div');
  post.classList.add('post');

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
  username.textContent = Post.author;

  const timestamp = document.createElement('p');
  timestamp.classList.add('timestamp');
  timestamp.textContent = Post.date;

  userDetails.appendChild(username);
  userDetails.appendChild(timestamp);
  userInfo.appendChild(avatar);
  userInfo.appendChild(userDetails);

  // Create the post content
  const postContent = document.createElement('p');
  postContent.classList.add('post-content');
  postContent.textContent = Post.content;

  // Create the post actions section
  const postActions = document.createElement('div');
  postActions.classList.add('post-actions');

  // Like button and notification
  const like = document.createElement('div');
  like.classList.add('like');

  const likeButton = document.createElement('button');
  const likeSvg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
  likeSvg.setAttribute('xmlns', 'http://www.w3.org/2000/svg');
  likeSvg.setAttribute('height', '20px');
  likeSvg.setAttribute('viewBox', '0 -960 960 960');
  likeSvg.setAttribute('width', '20px');
  likeSvg.setAttribute('fill', '#707C97');
  const likePath = document.createElementNS('http://www.w3.org/2000/svg', 'path');
  likePath.setAttribute('d', 'M720-144H264v-480l288-288 32 22q17 12 26 30.5t5 38.5l-1 5-38 192h264q30 0 51 21t21 51v57q0 8-1.5 14.5T906-467L786.93-187.8Q778-168 760-156t-40 12Zm-384-72h384l120-279v-57H488l49-243-201 201v378Zm0-378v378-378Zm-72-30v72H120v336h144v72H48v-480h216Z');
  likeSvg.appendChild(likePath);
  likeButton.appendChild(likeSvg);

  const likeNotification = document.createElement('span');
  likeNotification.classList.add('notification-icon');
  likeNotification.textContent = Post.likes;

  like.appendChild(likeButton);
  like.appendChild(likeNotification);

  // Dislike button and notification
  const dislike = document.createElement('div');
  dislike.classList.add('dislike');

  const dislikeButton = document.createElement('button');
  const dislikeSvg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
  dislikeSvg.setAttribute('xmlns', 'http://www.w3.org/2000/svg');
  dislikeSvg.setAttribute('height', '20px');
  dislikeSvg.setAttribute('viewBox', '0 -960 960 960');
  dislikeSvg.setAttribute('width', '20px');
  dislikeSvg.setAttribute('fill', '#707C97');
  const dislikePath = document.createElementNS('http://www.w3.org/2000/svg', 'path');
  dislikePath.setAttribute('d', 'M240-816h456v480L408-48l-32-22q-17-12-26-30.5t-5-38.5l1-5 38-192H120q-30 0-51-21t-21-51v-57q0-8 1.5-14.5T54-493l119-279q8-20 26.5-32t40.5-12Zm384 72H240L120-465v57h352l-49 243 201-201v-378Zm0 378v-378 378Zm72 30v-72h144v-336H696v-72h216v480H696Z');
  dislikeSvg.appendChild(dislikePath);
  dislikeButton.appendChild(dislikeSvg);

  const dislikeNotification = document.createElement('span');
  dislikeNotification.classList.add('notification-icon');
  dislikeNotification.textContent = Post.dislikes;

  dislike.appendChild(dislikeButton);
  dislike.appendChild(dislikeNotification);

  // Comment button and notification
  const comment = document.createElement('div');
  comment.classList.add('comments');

  const commentButton = document.createElement('button');
  const commentSvg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
  commentSvg.setAttribute('xmlns', 'http://www.w3.org/2000/svg');
  commentSvg.setAttribute('height', '20px');
  commentSvg.setAttribute('viewBox', '0 -960 960 960');
  commentSvg.setAttribute('width', '20px');
  commentSvg.setAttribute('fill', '#707C97');
  const commentPath = document.createElementNS('http://www.w3.org/2000/svg', 'path');
  commentPath.setAttribute('d', 'M864-96 720-240H360q-29.7 0-50.85-21.15Q288-282.3 288-312v-48h384q29.7 0 50.85-21.15Q744-402.3 744-432v-240h48q29.7 0 50.85 21.15Q864-629.7 864-600v504ZM168-462l42-42h390v-288H168v330ZM96-288v-504q0-29.7 21.15-50.85Q138.3-864 168-864h432q29.7 0 50.85 21.15Q672-821.7 672-792v288q0 29.7-21.15 50.85Q629.7-432 600-432H240L96-288Zm72-216v-288 288Z');
  commentSvg.appendChild(commentPath);
  commentButton.appendChild(commentSvg);

  const commentNotification = document.createElement('span');
  commentNotification.classList.add('notification-icon');
  commentNotification.textContent = Post.commentsCount;

  comment.appendChild(commentButton);
  comment.appendChild(commentNotification);
  // Append all post actions to the post-actions container
  postActions.appendChild(like);
  postActions.appendChild(dislike);
  postActions.appendChild(comment);
  let commentscontainer = document.createElement("div")
  commentscontainer.style.display = "none"
  commentscontainer.classList.add("comments-section")
  let cmtnum = 1
  comment.addEventListener("click", async function(){
    if (commentscontainer.style.display == "none"){
    let cmtloading = false
    commentscontainer.style.display = "block"
    let cmnts = await loadPosComments(Post.id,cmtnum);
    if (cmnts !== "baraka elik"){
      cmnts.forEach(cmt => createcomment(cmt,commentscontainer));
      cmtnum++
    commentscontainer.addEventListener("scroll", async () => {
      if (commentscontainer.scrollTop + commentscontainer.clientHeight >= commentscontainer.scrollHeight * 0.95 && !cmtloading || cmtnum == 1){
    try {
      let cmnts = await loadPosComments(Post.id,cmtnum);
      if (cmnts !== "baraka elik"){
      cmnts.forEach(cmt => createcomment(cmt,commentscontainer));
      commentscontainer.scrollTo(0, commentscontainer.scrollHeight*0.80)
      cmtnum = cmtnum+1
      cmtloading = false;
      }
    } catch (error) {
      console.error("Error loading comments:", error);
    }
  }
  })}
}else{
  commentscontainer.style.display = "none"
}
})
  // Append all sections to the post
  post.appendChild(userInfo);
  post.appendChild(postContent);
  post.appendChild(postActions);
  post.appendChild(commentscontainer);
  post.id = Post.id
  let postsection = document.querySelector(".posts-section")
  if (postsection == null){
    postsection = document.createElement("div")
    postsection.classList.add("posts-section")
    let postinput = postin()
    let addpostbutton = postinput.querySelector(".addpost")
    addpostbutton.addEventListener("click",async function() {
      let content = postinput.querySelector(".input").value
      let cats = postinput.querySelectorAll(".categorie")
      // todo for gheda sbt
    })
    postsection.appendChild(postinput)
  }
  let container = document.querySelector(".container")
  // Append the post to the body
  postsection.appendChild(post);
  container.appendChild(postsection)
  postsection.addEventListener("scroll", async () => {
    if (postsection.scrollTop + postsection.clientHeight >= postsection.scrollHeight&& !loading) {
      loading = true;
      console.log("Loading more posts...");
      try {
        let posts = await loadPosts(num);
        if(posts!= "baraka elik"){
          posts.forEach(post => createPost(post));
          postsection.scrollTo(0, postsection.scrollHeight*0.80)
          num = num+1
          loading = false;
        }
        
      } catch (error) {
        console.error("Error loading posts:", error);
      }
    }
  })
}
 function validinfos(user,action){
  function validbs(fields){
    for (const field of fields) {
      if (!user[field]) {
        return false;
      }
    }
    return true;
   }
  if (action == "login"){
    if (!validateEmail(email)) return false;
    if (!validatePassword(password)) return false;
  }else if (action == "register"){
    const {username, age, gender, firstname, lastname, email, password} = user

    if (!validbs(["username", "age", "gender", "firstname", "lastname", "email", "password"]))return false;
    if (!validateEmail(email))return false;
    if (!validatePassword(password))return false;
    if (!validlen(username,3,15) || !validlen(firstname,3,15) || !validlen(lastname,3,15) || (age > 100 || age < 12) || (gender != "male" && gender != "female")) { console.log(user);return false};
  }
  return true
 }
 function validlen(str,x,y){
  if (str.length < x || str.length > y){
    return false
  }
  return true
 }
 function validateEmail(email){
  if (email.length < 5 || email.length > 50) {
    return false;
  }
  return true;
 }
 function validatePassword(email){
  if (email.length < 5 || email.length > 30) {
    return false;
  }
  return true;
 }
 async function sendRegisterinfo(user){
  try {
    const data = await fetch("/api/signup",{
      method: "post",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    })
    if (data.ok){
      let card = document.querySelector(".card")
      card.classList.remove('flipped')
    }
  } catch (error) {
  }
 }
async function sendlogininfo(user){
  try {
    const data = await fetch("/api/login", {
      method: "post",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });
    if (data.ok) {
      servehome(user)
    } else {
      console.log( await data.text());
    }
  } catch (error) {
    errorSpan.textContent = "An error occurred. Please try again.";
  } finally {
    isSubmitting = false;
  }
}
async function servehome(user){
  Removecard()
  createSidebar(user)
 loadPosts(num).then(posts => {
   num++
  for (let post in posts){
    createPost(posts[post])
   }
 })
}
async function logout(){
      document.querySelector(".container").innerHTML = "";
      createCard()
      document.cookie ="session_token=;Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;"  
}
async function fetchPosts(num){
  const res = await fetch(`/api/post/?page-number=${num}`);
  const data = await res.json();
  return data;
}
async function loadPosts(num) {
  let response = await fetchPosts(num);
  let posts = response.Posts;
  if (posts.length == 0){
    return "baraka elik"
  }
  return posts;
}
async function getInfoData(){
  const res = await fetch("/api/info");
  const data = await res.json()
  if (res.ok) {
    return data;
  }
}
async function fetchComments(postId, cnum){
  const res = await fetch(`/api/post/${postId}/comments/${cnum}`);
  return await res.json();
}
async function loadPosComments(postId,cnum) {
  let response = await fetchComments(postId,cnum);
  let comments = response.Comments;
  if(comments.length == 0){
     return "baraka elik"
  }
  console.log(comments)
  return comments;
}
async function addpost(contentInput, categories) {
  const res = await fetch("/api/post/", {
    method: "post",
    body: JSON.stringify({
      content: contentInput,
      categories: categories,
    }),
  });
  return res
}
(async function(){
  if (!info.authorize){
    createCard()
    document.querySelector('.container').addEventListener("submit", async(e)=>{
      e.preventDefault()
      if (isSubmitting) return;
       isSubmitting = true;
       try {
         // Handle Login Form
         if (e.target.id === 'login-form') {
           const email = e.target.querySelector('#login-id').value;
           const password = e.target.querySelector('#login-password').value;
           
           if (validinfos({ email, password }, "login")) {
             await sendlogininfo({ email, password });
           }
         }
         // Handle Register Form
         else if (e.target.classList.contains('register-form')) {
           const username = e.target.querySelector('#nickname').value;
           const age = e.target.querySelector('#age').value;
           const gender = e.target.querySelector('#gender').value;
           const firstname = e.target.querySelector('#first-name').value;
           const lastname = e.target.querySelector('#last-name').value;
           const email = e.target.querySelector('#email').value;
           const password = e.target.querySelector('#password').value;
     
           if (validinfos({ username, age, gender, firstname, lastname, email, password }, "register")) {
           await sendRegisterinfo({ username, age: +age, gender, firstname, lastname, email, password });
           }
         }
       } finally {
         isSubmitting = false;
       }
     })
  }else{
    servehome(info)
  }
}());

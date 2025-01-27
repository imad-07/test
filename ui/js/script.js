function createSidebar() {
    // Create the main sidebar container
    const sidebar = document.createElement('div');
    sidebar.classList.add('sidebar');
  
    // Create the profile section
    const profileSection = document.createElement('div');
    profileSection.classList.add('profile-section');
  
    const profilePic = document.createElement('img');
    profilePic.id = 'profile-pic';
    profilePic.src = 'css/default-profile.jpg';
    profilePic.alt = 'Profile Picture';
  
    const userName = document.createElement('h3');
    userName.id = 'user-name';
    userName.textContent = 'User Name';
  
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
  }
  
  // Call the function to create and append the sidebar
  //createSidebar();
  let postsection = document.createElement("div")
  document.querySelectorAll('.like, .dislike, .comment').forEach((element) => {
    element.addEventListener('click', () => {
      element.classList.toggle('active');
    });
  });
  function Removesidebar(){
    document.addEventListener("keydown",e=>{
        if (e.key == "K"){
            let sidebar = document.querySelector(".sidebar")
            sidebar.remove()
        }
      })
  }
  const card = document.querySelector('.card');
const switchToRegister = document.getElementById('switch-to-register');
const switchToLogin = document.getElementById('switch-to-login');

switchToRegister.addEventListener('click', () => {
  card.classList.add('flipped');
});

switchToLogin.addEventListener('click', () => {
  card.classList.remove('flipped');
});

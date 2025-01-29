import { getInfoData } from "./auth.js";
import {
  sendLike,
  setLike,
  createFragment,
  addEvent,
  getIdFromUrl,
  showToast,
} from "./helpers.js";
import {
  createPostForm,
  getPosts,
  getSinglePost,
  sendPostFunctionality,
  sendCommentFunctionality,
  getComments,
} from "./post.js";

// Templates
const createAuthButton = (username) => `
  <div class="profile_menu dropdown">
      <button class="holder dropbtn">
          <div class="profile_image">
              <img src="/client/assets/img/author.jpg" alt="profile_img" />
          </div>
          <span class="username">${username}</span>
          <i class="fa-solid fa-chevron-down"></i>
      </button>
      <div class="dropdown-content">
          <a id="logoutBtn"><i class="fa-solid fa-arrow-right-from-bracket"></i> Logout</a>
      </div>
  </div>`;

const createLoginSignup = () => `
  <div class="login_signup">
      <a href="/login" class="login">Login</a>
      <a href="/signup" class="signup">Sign Up</a>
  </div>`;

const createPostTemplate = (username) => `
  <section class="create_post">
      <div class="profile_img">
          <img src="/client/assets/img/author.jpg" alt="profile_img" />
      </div>
      <input class="openModal" type="text" id="post_area" placeholder="What's on your mind, ${username}?"/>
      <button class="create_post openModal" id="create_post">Create Post</button>
  </section>`;

const createAuthCategory = () => `
  <div class="each_filter">
    <i class="fa-regular fa-newspaper"></i>
    <div class="filter_content">
        <h3>Posts:</h3>
        <select name="by_post" id="by_post">
            <option value="">All Posts</option>
            <option value="posted">Posts I Posted</option>
            <option value="liked">Posts I Liked</option>
        </select>
    </div>
  </div>`;

// Initialize Authentication UI
const initAuthUI = (info) => {
  const header = document.querySelector("header .content");
  const content = info.authorize
    ? createFragment(createAuthButton(info.username))
    : createFragment(createLoginSignup());
  header.appendChild(content);
};
const main_filter = document.querySelector(".left_aside.other_devices");
const mobile_filter = document.querySelector(".left_aside.filter_mobile");

// Initialize categories UI
const initCategoriesUI = (info) => {
  const filters = document.querySelectorAll(".filter");
  if (info.authorize) {
    filters.forEach((filter) => {
      filter.prepend(createFragment(createAuthCategory()));
    });
  }
  const categoriesSelects = document.querySelectorAll("#by_category");
  categoriesSelects.forEach((select) => {
    select.innerHTML = '<option value="">All</option>';
    select.innerHTML += info.categories
      .map((category) => `<option value="${category}">${category}</option>`)
      .join("");
  });
};
if (main_filter && getComputedStyle(main_filter).display != "none") {
  filter(main_filter);
} else if (mobile_filter && getComputedStyle(mobile_filter).display != "none") {
  filter(mobile_filter);
}
export const FiltringParams = new Map();

function filter(element) {
  const Submitbtn = element.querySelector(".submit");
  const ResetBtn = element.querySelector(".reset");
  Submitbtn.addEventListener("click", async () => {
    const categorie = document.querySelector("#by_category");
    const post = document.querySelector("#by_post");
    if (post) {
      FiltringParams["posts"] = post.value;
    }
    FiltringParams["categorie"] = categorie.value;

    postsPageNumber = await getPosts(
      document.querySelector("main .posts-holder"),
      0
    );
  });

  ResetBtn.addEventListener("click", async () => {
    const categorie = document.querySelector("#by_category");
    const post = document.querySelector("#by_post");
    if (post) {
      post.options[0].selected = true;
    }
    categorie.options[0].selected = true;

    FiltringParams["posts"] = "";
    FiltringParams["categorie"] = "";

    postsPageNumber = await getPosts(
      document.querySelector("main .posts-holder"),
      0
    );
  });
}

// Initialize Post Features
let postsPageNumber = 0;
const initPostFeatures = async (mainHolder, info) => {
  if (mainHolder) {
    const postsHolder = mainHolder.querySelector(".posts-holder");
    if (info.authorize) {
      const createPostElement = createFragment(
        createPostTemplate(info.username)
      );
      mainHolder.prepend(createPostElement);
      createPostForm(mainHolder, info.categories);
      const formPostElement = document.querySelector("form#modal");
      sendPostFunctionality(formPostElement, postsHolder);
    }
    postsPageNumber = await getPosts(postsHolder, 0);

    window.addEventListener("scroll", async () => {
      // Check if we've scrolled to the bottom
      if (
        window.scrollY + window.innerHeight >=
        document.documentElement.scrollHeight
      ) {
        postsPageNumber = await getPosts(postsHolder, postsPageNumber);
      }
    });
  }
};

// Initialize Single Post Features
const initSinglePostFeatures = async (mainHolder) => {
  if (mainHolder) {
    const postId = getIdFromUrl(location.href);

    await getSinglePost(mainHolder, postId);

    const commentsContainer = document.querySelector(".old_comments");
    // Send Comment
    if (info.authorize) {
      const commentInput = mainHolder.querySelector(".comment_input");
      sendCommentFunctionality(commentInput, commentsContainer);
    }

    // Get Comments
    let goToNextPage = await getComments(commentsContainer, postId, 1);
    let num = goToNextPage ? 2 : 1;

    window.addEventListener("scroll", async () => {
      // Check if we've scrolled to the bottom
      if (
        window.scrollY + window.innerHeight >=
        document.documentElement.scrollHeight
      ) {
        goToNextPage = await getComments(commentsContainer, postId, num);
        if (goToNextPage) num++;
      }
    });
  }
};

const liked = `../client/assets/img/like/thumbs-up-solid.svg`;
const notLiked = `../client/assets/img/like/thumbs-up-regular.svg`;
const disliked = `../client/assets/img/dislike/thumbs-down-solid.svg`;
const notDisliked = `../client/assets/img/dislike/thumbs-down-regular.svg`;

// Handle Likes and Dislikes
const handleLikeDislike = (selector, action, reactionType, info) => {
  document.querySelector(".main").addEventListener("click", async (e) => {
    if (e.target.matches(selector) || e.target.closest(selector)) {
      if (!info.authorize) {
        showToast("You must login first!!");
        return;
      }
      try {
        const post = e.target.closest(".post_card");
        const comment = e.target.closest(".comment");
        const id = comment ? comment.dataset.id : post.dataset.id;
        const type = comment ? "comment" : "post";
        const res = await action(type, reactionType, id);

        const dislikeImg = e.target
          .closest(".reactions")
          .querySelector(".dislike_btn");
        const likeImg = e.target
          .closest(".reactions")
          .querySelector(".like_btn");

        if (!res) {
          return;
        }

        likeImg.src = res.isliked ? liked : notLiked;
        dislikeImg.src = res.isdisliked ? disliked : notDisliked;
        setLike(
          e.target.closest(".reactions").querySelector(".total_likes"),
          res.Like
        );
        setLike(
          e.target.closest(".reactions").querySelector(".total_dislikes"),
          res.Dislike
        );
      } catch (err) {
        console.error("Error in like/dislike:", err);
      }
    }
  });
};

// Initialize App
export var info = await getInfoData();
var mainHolder = document.querySelector(".main-page");
(async () => {
  const singlePostHolder = document.querySelector(".single-post");

  // Initialize UI
  initAuthUI(info);
  initCategoriesUI(info);
  initPostFeatures(mainHolder, info);
  initSinglePostFeatures(singlePostHolder);

  // Like and Dislike
  handleLikeDislike(".like_btn", sendLike, 1, info);
  handleLikeDislike(".dislike_btn", sendLike, 2, info);

  // Logout
  addEvent("#logoutBtn", "click", () => {
    document.cookie =
      "session_token=;Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
    location.reload();
  });

  // Modal
  const modal = document.getElementById("modal");
  addEvent("#post_area", "focus", () => {
    modal.style.display = "flex";
    modal.querySelector("#title").focus();
  });
  addEvent("#create_post", "click", () => (modal.style.display = "flex"));
  window.addEventListener("click", (e) => {
    if (e.target === modal) modal.style.display = "none";
  });
})();

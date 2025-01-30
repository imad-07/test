import { createFragment, showToast } from "./helpers.js";
import { FiltringParams, info } from "./main.js";

const liked = `<img class="like_btn" src="../client/assets/img/like/thumbs-up-solid.svg" alt="like">`;
const notLiked = `<img class="like_btn" src="../client/assets/img/like/thumbs-up-regular.svg" alt="like">`;
const disliked = `<img class="dislike_btn" src="../client/assets/img/dislike/thumbs-down-solid.svg" alt="dislike">`;
const notDisliked = `<img class="dislike_btn" src="../client/assets/img/dislike/thumbs-down-regular.svg" alt="dislike">`;

// Categories Generating
const generateCategoryElement = (category) => `
    <li>
      <label class="category" for="${category}"
        ><input
          value="${category}"
          type="checkbox"
          id="${category}"
        />${category}</label
      >
    </li>
`;

const generateCategories = (categories) =>
  categories ? categories.map(generateCategoryElement).join("") : "";

// Templates
const createFormHTML = (categories) =>
  `<form id="modal" class="modal">
      <div class="modal_content">
        <div class="modal_header">
          <input name="title" id="title" type="text" placeholder="Title..." maxlength="500" required />
        </div>
        <div class="modal_body">
          <textarea name="content" id="content" placeholder="Tell your content..." maxlength="5000" required></textarea>
          <h4>
            Select Tags (up to 3) so readers know what your post is about:
          </h4>
          <ul class="categories">
          ${generateCategories(categories)}
          </ul>
        </div>
        <div class="modal_footer">
          <div class="acitons">
            <button class="publish_btn">Publish</button>
            <button class="cancel_btn">Cancel</button>
          </div>
        </div>
      </div>
    </form>`;

const createMainPostHTML = (postObj) => `
                            <section class="post_card" data-id="${postObj.id}">
                    <div class="about_post">
                        <div class="title">
                            <h1>
                                <a href="/post/${postObj.id}">${
  postObj.title
}</a>
                            </h1>
                        </div>
                        ${
                          postObj.categories != null
                            ? `
                          <div class="categories">
                            <ul>
                                ${
                                  postObj.categories
                                    ? postObj.categories
                                        .map(
                                          (category) =>
                                            "<li>" + category + "</li>"
                                        )
                                        .join("")
                                    : ""
                                }
                            </ul>
                        </div>`
                            : `<br>`
                        }
                        
                        <div class="info">
                            <div class="author">
                                <div class="profile_image">
                                    <img src="/client/assets/img/author.jpg" alt="profile_img" />
                                </div>
                                <div class="about_author">
                                    <p class="username">${postObj.author}</p>
                                    <span class="created_at">${new Date(
                                      postObj.date
                                    ).toLocaleDateString()}</span>
                                </div>
                            </div>
                            <div class="reactions">
                                <div class="likes">
                                        <span class="total_likes">${
                                          postObj.likes
                                        }</span>
                                        ${postObj.isliked ? liked : notLiked}

                                </div>
                                <div class="dislikes">
                                        <span class="total_dislikes">${
                                          postObj.dislikes
                                        }</span>

                                        ${
                                          postObj.isdisliked
                                            ? disliked
                                            : notDisliked
                                        }
                                      
                                </div>
                                <div class="comments">
                                    <a href="/post/${postObj.id}">
                                        <span class="total_comments">${
                                          postObj.commentsCount
                                        }</span>
                                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                                            stroke-width="1.5" stroke="currentColor" class="size-6">
                                            <path stroke-linecap="round" stroke-linejoin="round"
                                                d="M7.5 8.25h9m-9 3H12m-9.75 1.51c0 1.6 1.123 2.994 2.707 3.227 1.129.166 2.27.293 3.423.379.35.026.67.21.865.501L12 21l2.755-4.133a1.14 1.14 0 0 1 .865-.501 48.172 48.172 0 0 0 3.423-.379c1.584-.233 2.707-1.626 2.707-3.228V6.741c0-1.602-1.123-2.995-2.707-3.228A48.394 48.394 0 0 0 12 3c-2.392 0-4.744.175-7.043.513C3.373 3.746 2.25 5.14 2.25 6.741v6.018Z" />
                                        </svg>
                                    </a>
                                </div>
                            </div>
                        </div>
                    </div>
                </section>
        `;

const createSinglePostHTML = (postData) => `
        <section class="post_card" data-id="${postData.id}">
                    <div class="about_post">
                        <h1 class="title">${postData.title}</h1>
                        ${
                          postData.categories != null
                            ? `
                          <div class="categories">
                            <ul>
                                ${
                                  postData.categories
                                    ? postData.categories
                                        .map(
                                          (category) =>
                                            "<li>#" + category + "</li>"
                                        )
                                        .join("")
                                    : ""
                                }
                            </ul>
                        </div>`
                            : `<br>`
                        }
                        
                        <div class="content">
                            <pre>${postData.content}</pre>
                        </div>
                        <div class="reactions">
                <div class="likes">
                  <span class="total_likes">${postData.likes}</span>
                       ${postData.isliked ? liked : notLiked}
                </div>
                <div class="dislikes">
                  <span class="total_dislikes">${postData.dislikes}</span>
                       ${postData.isdisliked ? disliked : notDisliked}
                </div>
                <div class="comments">
                  <a>
                    <span class="total_comments">${
                      postData.commentsCount
                    }</span>
                    <svg class="comment_btn" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M7.5 8.25h9m-9 3H12m-9.75 1.51c0 1.6 1.123 2.994 2.707 3.227 1.129.166 2.27.293 3.423.379.35.026.67.21.865.501L12 21l2.755-4.133a1.14 1.14 0 0 1 .865-.501 48.172 48.172 0 0 0 3.423-.379c1.584-.233 2.707-1.626 2.707-3.228V6.741c0-1.602-1.123-2.995-2.707-3.228A48.394 48.394 0 0 0 12 3c-2.392 0-4.744.175-7.043.513C3.373 3.746 2.25 5.14 2.25 6.741v6.018Z"></path>
                    </svg>
                  </a>
                </div>
              </div>
                        <div class="comments">
                        ${
                          info.authorize
                            ? `
                          <div class="add_comment">
                                <div class="profile_image">
                                    <img src="/client/assets/img/author.jpg" alt="profile_image">
                                </div>
                                <form class="comment_input">
                                    <textarea name="comment" id="comment" placeholder="Add a comment"></textarea>
                                    <button class="btn_comment">Comment</button>
                                </form>
                            </div>
                          `
                            : ""
                        }
                            <div class="old_comments">
                            </div>
                        </div>
                    </div>
                </section>
    `;

const createCommentHTML = ({
  author,
  date,
  content,
  likes,
  dislikes,
  id,
  isliked,
  isdisliked,
}) =>
  `<div class="comment" data-id="${id}">
        <div class="profile_image">
            <img src="/client/assets/img/author.jpg" alt="profile_image">
        </div>
        <div class="comment_info">
            <div class="border">
                <h3 class="username">${author}<span class="comments_at">${new Date(
    date
  ).toLocaleString()}</span></h3>
                <p class="content">${content}</p>
            </div>
            <div class="reactions">
                <div class="likes">
                        <span class="total_likes">${likes}</span>
                        ${isliked ? liked : notLiked}
                </div>
                <div class="dislikes">
                        <span class="total_dislikes">${dislikes}</span>
                        ${isdisliked ? disliked : notDisliked}
                </div>
            </div>
        </div>
    </div>`;

// Form Post Creation
export const createPostForm = (container, categories) => {
  container.append(createFragment(createFormHTML(categories)));
};

// Send Posts Functionality
export const sendPostFunctionality = (postSender, postsHolder) => {
  const cancelBtn = postSender.querySelector(".cancel_btn");
  const titleInput = postSender.querySelector("#title");
  const contentInput = postSender.querySelector("#content");
  const resetForm = () => {
    titleInput.value = "";
    contentInput.value = "";
    postSender.style.display = "none";
    postSender
      .querySelectorAll(".categories input:checked")
      .forEach((el) => (el.checked = false));
  };

  postSender.addEventListener("submit", async (e) => {
    e.preventDefault();

    const categories = [
      ...postSender.querySelectorAll(".categories input:checked"),
    ].map((element) => element.value);

    if (!titleInput.value.trim() || !contentInput.value.trim()) {
      showToast("Invalid Content!!");
      return;
    }

    const res = await fetch("/api/post/", {
      method: "post",
      body: JSON.stringify({
        title: titleInput.value,
        content: contentInput.value,
        categories: categories,
      }),
    });

    if (res.ok) {
      resetForm();
      postsHolder.innerHTML = "";
      getPosts(postsHolder, 0);
    } else if (res.status === 401) {
      showToast("You must login first!!");
      return;
    } else if (res.status === 400) {
      showToast("Error creating post");
      return;
    } else if (res.status === 429) {
      showToast("Too many requests");
      return;
    }
  });

  cancelBtn?.addEventListener("click", resetForm);
};

// Get And Show Posts in the main page
let postIds = {};
const passToOldPosts = [false];
export const getPosts = async (container, num) => {
  const res = await fetchPosts(num, FiltringParams);
  const posts = res.Posts;

  if (!posts || posts.length === 0) {
    container.innerHTML =
      "<p class='posts-error'>There's no posts here, try to create your own!</p>";
    return;
  }

  if (num === 0) {
    postIds = {};
    container.innerHTML = "";
  }

  if (num % 10 === 0 && num != 0) {
    showOldPostsFunctionality(container, num, res.MetaData, passToOldPosts);
    if (!passToOldPosts[0]) {
      return;
    }
  }

  posts?.forEach((postObj) => {
    if (!postIds[postObj.id]) {
      postIds[postObj.id] = true;
      container.append(createFragment(createMainPostHTML(postObj)));
    }
  });

  return posts?.length === res.MetaData.StandardCount &&
    num < res.MetaData.PostsPages
    ? num + 1
    : num;
};

// show old posts functionality
const showOldPostsFunctionality = (container, num, metadata, pass) => {
  // check if the upcoming pages has posts
  if (
    num <= metadata.PostsPages &&
    !document.querySelector(".old-posts-btn") &&
    container.innerHTML
  ) {
    const showOldPostsBtn = document.createElement("button");
    showOldPostsBtn.classList.add("old-posts-btn");
    showOldPostsBtn.innerHTML = "Show old posts";

    container.append(showOldPostsBtn);
    pass[0] = false;
    showOldPostsBtn.addEventListener("click", () => {
      window.scrollTo(0, 0);
      container.innerHTML = "";
      pass[0] = true;
    });
  }
};

// Get Posts from the main endpoint
const fetchPosts = async (num, params) => {
  const urlParams = Object.entries(params)
    .map((param) => param.join("="))
    .join("&");

  const res = await fetch(`/api/post/?page-number=${num}&${urlParams}`);
  const data = await res.json();

  return data;
};

// Show Single Post
export const getSinglePost = async (container, id) => {
  const postData = await fetchSinglePost(id);
  if (id != postData.id) {
    location.href = "/";
    return;
  }
  container.innerHTML = createSinglePostHTML(postData);
  const usersince = document.querySelector(".usersince");
  const username = document.querySelector(".about_author .username");
  username.innerHTML = postData.author;
  usersince.innerHTML =
    "Joined at " +
    new Date(postData.joined_at).toLocaleDateString("en-US", {
      month: "short",
      year: "numeric",
    });
};

const fetchSinglePost = async (id) => {
  const res = await fetch(`/api/post/${id}`);
  const data = await res.json();

  return data;
};

// Send Comment functionality
export const sendCommentFunctionality = (commentForm, commentsContainer) => {
  commentForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const commentInput = commentForm.querySelector("textarea");
    const commentContent = commentInput.value;
    const postId = +commentForm.closest(".post_card").dataset.id;

    if (commentContent.trim().length == 0 || commentContent.length > 1000) {
      showToast("Invalid content!!");
      return;
    }
    commentInput.value = "";

    const res = await fetch("/api/comment", {
      method: "post",
      body: JSON.stringify({
        content: commentContent,
        postId: postId,
      }),
    });

    if (res.status === 201) {
      commentsContainer.innerHTML = "";
      getComments(commentsContainer, postId, 1);
    } else if (res.status === 401) {
      showToast("You must login first!!");
      return;
    }
  });
};

// Get Comments Funcitonality
let commentIds = {};
const passToOldComments = [false];
export const getComments = async (container, postId, pageNumber) => {
  const commentsRes = await fetchComments(postId, pageNumber);
  const comments = commentsRes.Comments;
  if (pageNumber === 1) {
    commentIds = {};
    container.innerHTML = "";
  }

  if (pageNumber % 10 === 0) {
    showOldCommentsFunctionality(
      container,
      pageNumber,
      commentsRes.MetaData,
      passToOldComments
    );
    if (!passToOldComments[0]) {
      return;
    }
  }

  comments.forEach((comment) => {
    if (!commentIds[comment.id]) {
      commentIds[comment.id] = true;
      container.appendChild(createFragment(createCommentHTML(comment)));
    }
  });

  // update comments count
  container.closest(".post_card").querySelector(".total_comments").innerHTML =
    commentsRes.MetaData.CommentsCount;

  return comments.length === commentsRes.MetaData.StandardCount;
};

// Fetch Comments
const fetchComments = async (postId, pageNmber) => {
  const res = await fetch(`/api/post/${postId}/comments/${pageNmber}`);
  return await res.json();
};

// show old comments functionality
const showOldCommentsFunctionality = (container, num, metadata, pass) => {
  // check if the upcoming pages has comments
  if (
    num <= metadata.CommentsPages &&
    !document.querySelector(".old-comments-btn") &&
    container.innerHTML
  ) {
    const showOldCommentsBtn = document.createElement("button");
    showOldCommentsBtn.classList.add("old-comments-btn", "old-posts-btn");
    showOldCommentsBtn.innerHTML = "Show old Comments";

    container.append(showOldCommentsBtn);
    pass[0] = false;
    showOldCommentsBtn.addEventListener("click", () => {
      window.scrollTo(0, 0);
      container.innerHTML = "";
      pass[0] = true;
    });
  }
};

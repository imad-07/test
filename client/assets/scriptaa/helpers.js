// Send Like to server
// React types 1 = like, 2 dislike
export const sendLike = async (type, reactType, id) => {
  let res = await fetch(`/api/reaction`, {
    method: "POST",
    body: JSON.stringify({
      thread_type: type,
      thread_id: Number(id),
      react: reactType,
    }),
  });
  if (res.ok) {
    return await res.json();
  } else if (res.status === 429) {
    showToast("Too many requests...");
  }
};

// Show likes and dislikes
export const setLike = (likeElement, totalLikes) => {
  const totalLikesElement = likeElement;
  totalLikesElement.textContent = totalLikes;
};

export const createFragment = (htmlString) =>
  document.createRange().createContextualFragment(htmlString);

export const addEvent = (selector, event, callback) => {
  const element = document.querySelector(selector);
  if (element) element.addEventListener(event, callback);
};

export const getIdFromUrl = (url) => {
  return url.split("/")[4];
};

// Show Toast
export const showToast = (() => {
  const toast = document.querySelector(".toast");
  const toastProgress = document.querySelector(".toast .progress");
  const toastCancel = document.querySelector(".toast .close");
  const toastText = document.querySelector(".toast .text-2");
  const resetToast = (IntervalId) => {
    if (IntervalId) {
      clearInterval(IntervalId);
    }
    toast.classList.remove("active");
    toastProgress.classList.remove("active");
  };
  let IntervalId = 0;
  toastCancel.addEventListener("click", () => resetToast(IntervalId));

  return (message) => {
    if (toast.classList.contains("active")) {
      return;
    }

    toastText.textContent = message;
    toast.classList.add("active");
    toastProgress.classList.add("active");

    IntervalId = setTimeout(resetToast, 5000);
  };
})();

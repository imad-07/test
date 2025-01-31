commentscontainer.style.display = "none"
commentscontainer.classList.add("comments-section")
let cmtnum = 1
let cclik = 1
comment.addEventListener("click", async function(){
  if (cclik % 2 == 0){
  let cmtloading = false
  commentscontainer.style.display = "block"
  let cmnts = await loadPosComments(Post.id,cmtnum);
    cmnts.forEach(cmt => createcomment(cmt));
  commentscontainer.addEventListener("scroll", async () => {
    if (commentscontainer.scrollTop + commentscontainer.clientHeight >= commentscontainer.scrollHeight&& !cmtloading){
  try {
    let cmnts = await loadPosComments(Post.id,cmtnum);
    cmnts.forEach(cmt => createcomment(cmt));
  } catch (error) {
    console.error("Error loading comments:", error);
  }
  commentscontainer.scrollTo(0, commentscontainer.scrollHeight*0.80)
  cmtnum = cmtnum+1
  cmtloading = false;
}
})
}else{c
cclik++
commentscontainer.style.display = "none"
}
})
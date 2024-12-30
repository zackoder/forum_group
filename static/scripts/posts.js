export async function addEventOnPosts(path) {
  document.addEventListener("DOMContentLoaded", function () {
    const postsContainer = document.getElementById("posts-container");
    PostCategory();

    // Event delegation for click events
    postsContainer.addEventListener("click", function (event) {
      const postElement = event.target.closest(".post-container");
      const commentEl = event.target.closest(".commentC");
      if (commentEl) {
        const commentId = commentEl.getAttribute("data-comment-id");
        console.log(commentId);
        if (event.target.classList.contains("like-btn-comment")) {
          handleLike("comment", commentId, "like");
          console.log(commentId);
        } else if (event.target.classList.contains("dislike-btn-comment")) {
          console.log(commentId);
          handleLike("comment", commentId, "dislike");
        }
      }
      const postId = postElement.getAttribute("data-post-id");

      if (
        event.target.classList.contains("like-btn") ||
        event.target.classList.contains("likeIcon")
      ) {
        handleLike("posts", postId, "like");
      } else if (
        event.target.classList.contains("dislike-btn") ||
        event.target.classList.contains("dislikeIcon")
      ) {
        handleLike("posts", postId, "dislike");
      }
      if (event.target.classList.contains("like-btn-comment")) {
        handleLike("comment", commentId, "like");
        console.log(commentId);
      } else if (event.target.classList.contains("dislike-btn-comment")) {
        console.log(commentId);
        handleLike("comment", commentId, "dislike");
      }
    });

    postsContainer.addEventListener("submit", function (event) {
      const postElement = event.target.closest(".post-container");
      if (event.target.classList.contains("comment_form")) {
        event.preventDefault();
        const form = event.target;

        const postId = postElement.getAttribute("data-post-id");
        const commentText = form.querySelector(".comment").value.trim();

        if (commentText === "") {
          alert("Comment cannot be empty.");
          return;
        }

        handleComment(postId, commentText);
        form.reset();
      }
    });

    postsContainer.addEventListener("click", function (event) {
      const postElement = event.target.closest(".post-container");

      const postId = postElement.getAttribute("data-post-id");

      const CommentClass = event.target.classList.contains("see_comments");

      if (event.target.classList.contains("see_comments")) {
        event.target.disabled = true;
      }
      const divcomments = document.querySelector(".divcomments" + postId);
      console.log(divcomments);

      if (CommentClass && postId) {
        GetComments(postId, divcomments);
      }
    });

    loadMorePosts(path);
    window.addEventListener("scrollend", () => {
      setInterval(handleScroll, 5000);
    });
  });
}

function handleLike(path, id, like) {
  console.log(like);

  fetch(`/api/${path}/reaction/${id}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: `action=${like}`,
  })
    .then((response) => {
      console.log(response);

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json();
    })
    .then((data) => {
      console.log("Like/Dislike updated:", data);
    })
    .catch((error) => console.error("Error updating like/dislike:", error));
}

function handleComment(postId, comment) {
  fetch(`/api/${postId}/comment/new`, {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: `comment=${comment}`,
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json();
    })
    .then((data) => {
      console.log(data);
      if (data.message != 200) {
        alert(" faild to add Comment");
      }
      let comment_form = ".divcomments" + postId;
      let commentElement = document.querySelector(comment_form);
      GetComments(postId, commentElement);
    })
    .catch((error) => alert("Error submitting comment:", error));
}

let profile = document.getElementById("profile");
if (!profile) {
  document.getElementById("posts-container").style.paddingTop = "120px";
}

let offset = 0;
const limit = 20;
let loading = false;
const onehour = 1000 * 60 * 60;
const oneday = 1000 * 60 * 60 * 24;
const onemin = 1000 * 60;

export async function loadMorePosts(path) {
  if (loading) return;
  loading = true;

  try {
    const response = await fetch(`${path}?offset=${offset}`);

    const posts = await response.json();

    if (!posts || posts.length === 0) return;
    createPosts(posts);

    offset += limit;
  } catch (error) {
    console.error("Error loading posts:", error);
  } finally {
    loading = false;
  }
}

function createPosts(posts) {
  const postsContainer = document.getElementById("posts-container");
  posts.forEach((post) => {
    const postElement = document.createElement("div");
    postElement.className = "post-container";
    postElement.dataset.postId = post.Id;

    /* h2 will contain the image and name of the persen who posted */
    const posterName = createEle("h2");
    posterName.className = "poster";
    const posterImg = createEle("img");
    posterImg.src =
      "/static/images/466006304_871124095226532_8631138819273739648_n.jpg";
    const nameContainer = createEle("span");
    nameContainer.className = "usrname";
    nameContainer.innerText = post.UserName;
    const createdat = createEle("span");

    createdat.innerText = formDate(post.Date);
    createdat.className = "creationdate";
    posterName.appendChild(posterImg);
    posterName.appendChild(nameContainer);
    postElement.appendChild(posterName);
    postElement.appendChild(createdat);

    /* creating a div that will contain all the elements bellow */
    const pc = createEle("div");
    pc.className = "pc";

    /* creating an h3 element to contain the post title */
    const title = createEle("h3");
    title.className = "title";
    title.innerText = post.Title;

    /* creating a p element that will contain the content of the post */
    const content = createEle("p");
    content.className = "content";
    content.innerText = post.Content;
    pc.append(title, content);

    const categories_container = createEle("div");
    categories_container.className = "categories";

    for (let cate of post.Categories) {
      const span = createEle("span");
      span.className = "category";
      span.innerText = cate;
      categories_container.appendChild(span);
    }
    pc.appendChild(categories_container);
    /* creating like and dislike button */
    const like_dislike_container = createEle("div");
    like_dislike_container.className = "like-dislike-container";

    /* creating of the like button */
    const likebnt = createEle("button");
    likebnt.className = "like-btn";

    /* create an img element to contain like icon */
    const likeIcon = createEle("img");
    likeIcon.className = "likeIcon";
    likeIcon.src = "/static/images/like.png";

    const likeNbr = createEle("span");
    likeNbr.className = "likeNbr";
    likeNbr.innerText = post.Reactions.Likes;

    likebnt.appendChild(likeIcon);
    // likebnt.appendChild(likeNbm);

    /* creationg of the dislike button */
    const dislikebnt = createEle("button");
    dislikebnt.className = "dislike-btn";

    /* creating an img tag to containg dislike icon */
    const dislikeIcon = createEle("img");
    dislikeIcon.className = "dislikeIcon";
    dislikeIcon.src = "/static/images/dislike.png";

    const dislikeNbr = createEle("span");
    dislikeNbr.className = "dislikeNbr";
    dislikeNbr.innerText = post.Reactions.Dislikes;

    dislikebnt.appendChild(dislikeIcon);
    // dislikebnt.appendChild(dislikeNbr);

    if (post.Reactions.Action === "like") {
      likebnt.style.classList.add("liked");
    } else if (post.Reactions.Action === "dislike") {
      dislikebnt.style.classList.add("disliked");
    }
    /* appending like and dislike buttons to like container */
    like_dislike_container.append(likebnt, likeNbr, dislikebnt, dislikeNbr);

    /* appending like container to the post contaner */
    pc.appendChild(like_dislike_container);

    ////add div comments
    const divcomments = createEle("div");
    divcomments.className = `divcomments${post.Id} divcomments`;

    pc.appendChild(divcomments);

    /* adding a button to see comments */
    const seecomments = createEle("button");
    seecomments.className = "see_comments";
    seecomments.innerText = "see comments";
    pc.appendChild(seecomments);
    /* creating the form that sends comments */
    const comment_form = createEle("form");
    comment_form.method = "POST";
    comment_form.className = "comment_form";

    const title_impt = createEle("input");
    title_impt.className = "comment";
    title_impt.name = "comment";
    title_impt.type = "text";
    title_impt.placeholder = "Add your comment";
    title_impt.required = true;

    const submit_comment = createEle("button");
    submit_comment.className = "send_comment";
    submit_comment.type = "submit";

    const send_icon = createEle("img");
    send_icon.className = "sendimg";
    send_icon.src = "/static/images/send-message.png";
    submit_comment.appendChild(send_icon);
    comment_form.appendChild(title_impt);
    comment_form.appendChild(submit_comment);

    pc.appendChild(comment_form);
    postElement.appendChild(pc);

    postsContainer.appendChild(postElement);
  });
}

function formDate(date) {
  let CreattionDate = new Date(date).getTime();
  const currentTime = Date.now();
  const elapsed = currentTime - CreattionDate;

  const days = Math.floor(elapsed / oneday);
  const hours = Math.floor((elapsed % oneday) / onehour);
  const minutes = Math.floor((elapsed % onehour) / onemin);

  let timeText = "";

  if (days > 0) {
    timeText += `${days}d `;
  }
  if (hours > 0) {
    timeText += `${hours}h `;
  }
  if (minutes > 0 && (days == 0 || hours == 0)) {
    timeText += `${minutes}min`;
  }
  return timeText;
}

function createEle(elename) {
  return document.createElement(elename);
}

const showPostFormButton = document.querySelector(".show-postForm");
const postForm = document.querySelector(".postForm");
const layout = document.querySelector(".lay-out"); // Optional dimmed background

// Show the form
if (showPostFormButton) {
  showPostFormButton.addEventListener("click", () => {
    postForm.style.display = "flex";
    layout.style.display = "block";
    document.body.style.overflow = "hidden";

    // Hide the form when clicking outside or on a cancel button
    layout.addEventListener("click", () => {
      postForm.style.display = "none";
      layout.style.display = "none";
      document.body.style.overflow = "";
    });
  });
}

export function handleScroll(path) {
  const scrollPosition = window.scrollY + window.innerHeight;
  const threshold = document.body.scrollHeight - 1000;

  if (scrollPosition > threshold) {
    loadMorePosts(path);
  }
}

const form = document.getElementById("postForm");
if (form) {
  form.addEventListener("submit", async function (event) {
    event.preventDefault();
    // Declare validation flags
    let isValidTitle = true;
    let isValidContent = false;
    let isValidCheckboxes = false;

    // Get form values
    let Title = document.getElementById("post").value;
    let Content = document.getElementById("content").value;
    let categoryName = [];

    // Get selected checkboxes
    let checkboxes = document.querySelectorAll('input[name="options"]:checked');
    checkboxes.forEach((checkbox) => {
      categoryName.push(checkbox.getAttribute("data-name"));
    });

    if (Content === "") {
      document.getElementById("errorContent").innerHTML = "Content is required";
      document.getElementById("errorContent").style.color = "red";
      isValidContent = false;
    } else {
      document.getElementById("errorContent").innerHTML = "";
      isValidContent = true;
    }

    if (categoryName.length === 0) {
      document.getElementById("errorcategory").innerHTML =
        "Please select at least one category";
      document.getElementById("errorcategory").style.color = "red";
      isValidCheckboxes = false;
    } else {
      document.getElementById("errorcategory").innerHTML = "";
      isValidCheckboxes = true;
    }

    if (isValidTitle && isValidContent && isValidCheckboxes) {
      try {
        const res = await fetch("/add-post", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          body: new URLSearchParams({
            Title: Title,
            Content: Content,
            options: categoryName,
          }),
        });

        if (res.ok) {
          window.location.href = res.url;
        } else {
          alert("Failed to submit post");
          // console.log(res);
        }
      } catch (error) {
        alert("Error: " + error.message);
      }
    }
  });
}

async function PostCategory() {
  let category = document.getElementById("category");
  try {
    const res = await fetch("http://localhost:8001/api/category/list");
    const categories = document.querySelector("#categories");
    const data = await res.json();
    data.forEach((catg) => {
      let li = createEle("li");
      let a = createEle("a");
      console.log(li);
      console.log(a);
      console.log(catg.Name);

      a.href = `/categoreis/${catg.Id}`;
      a.innerText = catg.Name;
      category.innerHTML += `
      <label class="catLabel" for="${catg.Name}">
        <input type="checkbox" name="options" id="${catg.Name}" value="${catg.Name}" data-name="${catg.Name}"> <splan>${catg.Name}</span>
      </label>

      `;
      li.append(a);
      categories.append(li);
    });
  } catch {
    console.log("erroure");
  }
}

async function GetComments(idPost, str) {
  str.innerText = "";
  str.style.display = "block";
  try {
    const response = await fetch(
      `http://localhost:8001/api/${idPost}/comments`
    );

    if (response.ok) {
      const data = await response.json();

      if (data == null) {
        str.innerHTML = "there is no comments";
      } else {
        const comments = createEle("div");
        comments.className = "commentsDiv";

        data.forEach((e) => {
          const commentC = createEle("div");
          commentC.className = "commentC";
          commentC.dataset.commentId = e.Id;

          const commentHe = createEle("div");
          commentHe.className = "commentHe";

          const commentH = createEle("h3");
          commentH.className = "commentH";
          commentH.innerText = e.Username;

          const commentTime = createEle("p");
          commentTime.className = "commentTime";
          commentTime.innerText = formDate(e.Date);

          const commentP = createEle("p");
          commentP.className = "commentp";
          commentP.innerText = e.Comment;

          const like_dislike_container = createEle("div");
          like_dislike_container.className = "like-dislike-container-comment";

          /* creating of the like button */
          const likebnt = createEle("button");
          likebnt.className = "like-btn-comment";

          /* create an img element to contain like icon */
          const likeIcon = createEle("img");
          likeIcon.className = "likeicon-comment";
          likeIcon.src = "/static/images/like.png";

          likebnt.appendChild(likeIcon);

          const likeNmb = createEle("span");
          likeNmb.className = "likeNbr";
          likeNmb.innerText = e.Reactions.Likes;
          // likebnt.appendChild(likeNmb);

          /* creationg of the dislike button */
          const dislikebnt = createEle("button");
          dislikebnt.className = "dislike-btn-comment";

          /* creating an img tag to containg dislike icon */
          const dislikeIcon = createEle("img");
          dislikeIcon.className = "dislikeicon-comment";
          dislikeIcon.src = "/static/images/dislike.png";

          dislikebnt.appendChild(dislikeIcon);

          const dislikeNmb = createEle("span");
          dislikeNmb.className = "dislikeNbr";
          dislikeNmb.innerText = e.Reactions.Dislikes;
          // dislikebnt.appendChild(dislikeNmb);

          /* appending like and dislike buttons to like container */
          like_dislike_container.append(
            likebnt,
            likeNmb,
            dislikebnt,
            dislikeNmb
          );

          commentHe.append(commentH, commentTime);
          commentC.append(commentHe, commentP, like_dislike_container);
          comments.appendChild(commentC);
        });
        str.appendChild(comments);
      }
    } else {
      console.error("Request failed with status:", response.status);
      // document.getElementById("responseMessage").innerText = "Error fetching comments.";
    }
  } catch (error) {
    console.error(error);
  }
}

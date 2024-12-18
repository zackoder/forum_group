document.addEventListener("DOMContentLoaded", function () {
  const postsContainer = document.getElementById("posts-container");

  // Event delegation for click events
  postsContainer.addEventListener("click", function (event) {
    const postElement = event.target.closest(".post-container");
    if (!postElement) return;

    const postId = postElement.getAttribute("data-post-id");

    if (event.target.classList.contains("like-btn")) {
      console.log("Like Post ID:", postId);
      handleLike(postId, true);
    } else if (event.target.classList.contains("dislike-btn")) {
      console.log("Dislike Post ID:", postId);
      handleLike(postId, false);
    }
  });

  postsContainer.addEventListener("submit", function (event) {
    if (event.target.classList.contains("comment_form")) {
      const postElement = event.target.closest(".post-container");
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

  loadMorePosts();
  window.addEventListener("scroll", _.throttle(handleScroll, 500));
});

function handleLike(postId, like) {
  fetch("/like-post", {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: `post_id=${postId}&like=${like}`,
  })
    .then((response) => {
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
  fetch("/comments", {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: `post_id=${postId}&comment=${comment}`,
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json();
    })
    .then((data) => {
      if (data.success) {
        alert("Comment added successfully!");
      } else {
        alert("Failed to add comment.");
      }
    })
    .catch((error) => console.error("Error submitting comment:", error));
}

let home = "home";
let profile = document.getElementById("profile");
if (profile) {
  profile.addEventListener("click", () => {
    loadMorePosts(profile);
  });
}

let offset = 0;
const limit = 20;
let loading = false;

async function loadMorePosts(name) {
  // console.log(name);
  // if (name !== "home" || name !== "profile") name = "home";
  // if (loading) return;
  // loading = true;

  try {
    const response = await fetch(`http://localhost:8001/api/posts`)
      .then((res) => {
        console.log(res.json());
      })
      .catch((err) => {
        alert(err);
        console.log(err);
      });

    if (!response.ok) alert("not ok");
    const posts = await response.json();
    // if (!posts || posts.length === 0) return;
    console.log(posts);

    const postsContainer = document.getElementById("posts-container");
    posts.forEach((post) => {
      const postElement = document.createElement("div");
      postElement.className = "post-container";
      postElement.dataset.postId = post.ID;

      /* h2 will contain the image and name of the persen who posted */
      const posterName = createEle("h2");
      posterName.className = "poster";
      const posterImg = createEle("img");
      posterImg.src =
        "/css/466006304_871124095226532_8631138819273739648_n.jpg";
      const nameContainer = createEle("span");
      nameContainer.innerText = post.UserName;
      posterName.append(posterImg, nameContainer);
      postElement.appendChild(posterName);

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

      /* creating like and dislike button */
      const like_dislike_container = createEle("div");
      like_dislike_container.className = "like-dislike-container";

      /* creating of the like button */
      const likebnt = createEle("button");
      likebnt.className = "like-btn";

      /* create an img element to contain like icon */
      const likeIcon = createEle("img");
      likeIcon.src = "/css/like.png";

      likebnt.appendChild(likeIcon);

      /* creationg of the dislike button */
      const dislikebnt = createEle("button");
      dislikebnt.className = "dislike-btn";

      /* creating an img tag to containg dislike icon */
      const dislikeIcone = createEle("img");
      dislikeIcone.src = "/css/dislike.png";

      dislikebnt.appendChild(dislikeIcone);

      /* appending like and dislike buttons to like container */
      like_dislike_container.append(likebnt, dislikebnt);

      /* appending like container to the post contaner */
      pc.appendChild(like_dislike_container);

      /* creating the form that sends comments */
      const comment_form = createEle("form");
      comment_form.method = "post";
      comment_form.className = "comment_form";

      const title_impt = createEle("input");
      title_impt.className = "comment";
      title_impt.name = "comment";
      title_impt.type = "text";
      title_impt.placeholder = "Add your comment";
      title_impt.required = true;

      const submit_comment = createEle("button");
      submit_comment.className = "send_comment";

      const send_icon = createEle("img");
      send_icon.className = "sendimg";
      send_icon.src = "/css/send-message.png";
      comment_form.appendChild(title_impt);
      comment_form.appendChild(send_icon);

      pc.appendChild(comment_form);
      postElement.appendChild(pc);

      postsContainer.appendChild(postElement);
    });

    offset += limit;
  } catch (error) {
    console.error("Error loading posts:", error);
  } finally {
    loading = false;
  }
}

function createEle(elename) {
  return document.createElement(elename);
}
let lay_outbtn = document.querySelector(".show-postForm");

lay_outbtn.addEventListener("click", () => {
  let layOutDiv = document.querySelector(".lay-out");
  let postForm = document.querySelector(".postForm");

  if (layOutDiv.classList.contains("active")) {
    layOutDiv.classList.remove("active");
    postForm.classList.remove("active");
    layOutDiv.style.display = "none";
    postForm.style.display = "none";
  } else if (!layOutDiv.classList.contains("active")) {
    layOutDiv.classList.add("active");
    postForm.classList.add("active");
    layOutDiv.style.display = "block";
    postForm.style.display = "flex";
    document.body.style.overflow = "hidden";
  }
});

document.body.addEventListener("keyup", (e) => {
  if (e.key === "Escape") {
    document.body.style.overflow = "";
    document.querySelectorAll(".active").forEach((btn) => {
      btn.classList.remove("active");
      btn.style.display = "none";
    });
  }
});

function handleScroll() {
  const scrollPosition = window.scrollY + window.innerHeight;
  const threshold = document.body.scrollHeight - 1000;

  if (scrollPosition > threshold) {
    loadMorePosts();
  }
}
setInterval();

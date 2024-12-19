document.addEventListener("DOMContentLoaded", function () {
  const postsContainer = document.getElementById("posts-container");
  PostCategory();

  // Event delegation for click events
  postsContainer.addEventListener("click", function (event) {

    const postElement = event.target.closest(".post-container");
    if (!postElement) return;


    const postId = postElement.getAttribute("data-post-id");

    if (event.target.classList.contains("like-btn")) {
      handleLike(postId, true);
    } else if (event.target.classList.contains("dislike-btn")) {
      handleLike(postId, false);
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
    if (!postElement) return;


    const postId = postElement.getAttribute("data-post-id");


    console.log(postId);
    let CommentClass = event.target.classList.contains("see_comments")
    if (CommentClass) {
      alert(CommentClass)
      GetComments(postId, "see_comments")
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
    body: `{
          "post_id": ${postId},
          "like":${like}
          }`,
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
  fetch(`api/${postId}/comment/new`, {
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
      console.log(data)
      if (data.message != 200) {
        
        alert(" faild to add Comment");
      } 
    })
    .catch((error) => alert("Error submitting comment:", error));
    
}

let home = "home";
let profile = document.getElementById("profile");
if (profile) {
  profile.addEventListener("click", () => {
    loadMorePosts(profile);
  });
} else {
  document.getElementById("posts-container").style.paddingTop = "120px";
}

let offset = 0;
const limit = 20;
let loading = false;

async function loadMorePosts(name = "home") {
  console.log(name);

  if (loading) return;
  loading = true;

  try {
    console.log("hello");
    const response = await fetch(`/api/posts?offset=${offset}`);

    const posts = await response.json();
    if (!posts || posts.length === 0) return;

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
      nameContainer.innerText = post.Username;
      const createdat = createEle("span");
      const date = +new Date(post.Date);
      const curontTime = +new Date();
      // alert(curontTime - data);
      console.log("minets", Math.floor((curontTime - date) / 1000 / 60));
      console.log("hours", Math.floor((curontTime - date) / 1000 / 60 / 60));
      console.log(
        "hours",
        Math.floor((curontTime - date) / 1000 / 60 / 60 / 24)
      );

      if (Math.floor((curontTime - date) / 1000 / 60 / 60) > 24) {
        createdat.innerText += `${Math.floor(
          (curontTime - date) / 1000 / 60 / 60 / 24
        )}d`;
      }
      if (Math.floor((curontTime - date) / 1000 / 60 / 60) !== 0) {
        createdat.innerText += `${Math.floor(
          (curontTime - date) / 1000 / 60 / 60
        )}h`;
      }
      if ((curontTime - date) / 1000 / 60 !== 0) {
        createdat.innerText += `${Math.floor(
          (curontTime - date) / 1000 / 60
        )}min`;
      }

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

      for (let cate of post.Categories.split(",")) {
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
      likeIcon.src = "/static/images/like.png";

      likebnt.appendChild(likeIcon);

      /* creationg of the dislike button */
      const dislikebnt = createEle("button");
      dislikebnt.className = "dislike-btn";

      /* creating an img tag to containg dislike icon */
      const dislikeIcone = createEle("img");
      dislikeIcone.src = "/static/images/dislike.png";

      dislikebnt.appendChild(dislikeIcone);

      /* appending like and dislike buttons to like container */
      like_dislike_container.append(likebnt, dislikebnt);

      /* appending like container to the post contaner */
      pc.appendChild(like_dislike_container);

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

const showPostFormButton = document.querySelector(".show-postForm");
const postForm = document.querySelector(".postForm");
const layout = document.querySelector(".lay-out"); // Optional dimmed background

// Show the form
showPostFormButton.addEventListener("click", () => {
  postForm.style.display = "flex";
  layout.style.display = "block";
  document.body.style.overflow = "hidden";
});

// Hide the form when clicking outside or on a cancel button
layout.addEventListener("click", () => {
  postForm.style.display = "none";
  layout.style.display = "none";
  document.body.style.overflow = "";
});

document.addEventListener("DOMContentLoaded", () => {
  const postsContainer = document.getElementById("posts-container");
  const profileLink = document.getElementById("profile");

  if (profileLink) {
    profileLink.addEventListener("click", async (event) => {
      event.preventDefault(); // Prevent default navigation
      let offset = 0;
      postsContainer.innerHTML = "";

      // Fetch posts for the profile
      try {
        const response = await fetch(`/api/posts?offset=${offset}`);
        if (!response.ok) {
          throw new Error(
            `Error fetching profile posts: ${response.statusText}`
          );
        }

        const posts = await response.json();
        renderPosts(posts, postsContainer);
      } catch (error) {
        console.error("Failed to fetch profile posts:", error);
        postsContainer.innerHTML =
          "<p>Error loading profile posts. Please try again later.</p>";
      }
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

const form = document.getElementById("postForm");

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
      const res = await fetch("http://localhost:8001/add-post", {
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
        console.log(res);
      }
    } catch (error) {
      alert("Error: " + error.message);
    }
  }
});

async function PostCategory() {
  let category = document.getElementById("category");
  try {
    const res = await fetch("http://localhost:8001/api/category/list");
    const data = await res.json();
    data.forEach((catg) => {
      category.innerHTML += `
      <label class="catLabel" for="${catg.Name}">
        <input type="checkbox" name="options" id="${catg.Name}" value="${catg.Name}" data-name="${catg.Name}"> <splan>${catg.Name}</span>
      </label>

      `;
    });
  } catch {
    console.log("erroure");
  }
}




async function GetComments(idPost, str) {
  alert(idPost)
  alert(str)
  try {
    const response = await fetch(`http://localhost:8001/api/${idPost}/comments`)

    if (response.ok) {
      const data = await response.json();

      console.log(data);
    } else {
      console.error("Request failed with status:", response.status);
      // document.getElementById("responseMessage").innerText = "Error fetching comments.";
    }
  } catch (error) {

  }
}
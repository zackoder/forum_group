export function HandulLike(type, numlike, numdislike, path, id) {
  console.log(type);
  const likebtn = document.createElement("button");
  const dislikebtn = document.createElement("button");
  likebtn.className = "like";
  dislikebtn.className = "dislike";

  if (type === "like") {
    likebtn.classList.add("liked");
  } else if (type === "dislike") {
    dislikebtn.classList.add("disliked");
  }
  const likespan = document.createElement("span");
  const dislikespan = document.createElement("span");
  likespan.innerHTML = numlike;
  dislikespan.innerHTML = numdislike;
  likebtn.addEventListener("click", () => {
    if (likebtn.classList.length == 2) {
      likebtn.classList.remove("like");
      numlike--;
    } else {
      if (dislikebtn.classList.length == 2) {
        dislikebtn.classList.remove("dislike");
        numdislike--;
      }
      numlike++;
      likebtn.classList.add("like");
    }
    likespan.innerHTML = numlike;
    dislikespan.innerHTML = numdislike;
    fetch(`/api/${path}/reaction/${id}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: `action=like`,
    });
  });
  dislikebtn.addEventListener("click", () => {
    if (dislikebtn.classList.length == 2) {
      dislikebtn.classList.remove("dislike");
      numdislike--;
    } else {
      if (likebtn.classList.length == 2) {
        likebtn.classList.remove("like");
        numlike--;
      }
      numdislike++;
      dislikebtn.classList.add("dislike");
    }
    likespan.innerHTML = numlike;
    dislikespan.innerHTML = numdislike;
    fetch(`/api/${path}/reaction/${id}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: `action=dislike`,
    }).then((res) => {
      // console.log(res);
    });
  });
  return [likebtn, dislikebtn, likespan, dislikespan];
}

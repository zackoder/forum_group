export function HandulLike(type, numlike, numdislike, path, id) {
  const likebtn = document.createElement("button");
  const dislikebtn = document.createElement("button");
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
      likebtn.classList.remove("liked");
      numlike--;
    } else {
      if (dislikebtn.classList.length == 2) {
        dislikebtn.classList.remove("disliked");
        numdislike--;
      }
      numlike++;
      likebtn.classList.add("liked");
    }
    likespan.innerHTML = numlike;
    dislikespan.innerHTML = numdislike;
    fetch(`/api/${path}/reaction/${id}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: `action=like`,
    }).then(res=>{    
      if (res.redirected){
        window.location.href="login"
      }
    });
  });
  dislikebtn.addEventListener("click", () => {
    if (dislikebtn.classList.length == 2) {
      dislikebtn.classList.remove("disliked");
      numdislike--;
    } else {
      if (likebtn.classList.length == 2) {
        likebtn.classList.remove("liked");
        numlike--;
      }
      numdislike++;
      dislikebtn.classList.add("disliked");
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
      if (res.redirected) {
        window.location.href = "login";
      }
    });
  });
  return [likebtn, dislikebtn, likespan, dislikespan];
}

import { addEventOnPosts, handleScroll } from "./posts.js";
path = "/api/created/posts";
console.log(path);

addEventOnPosts(path);

window.addEventListener("scrollend", () => {
  setInterval(handleScroll(path));
});

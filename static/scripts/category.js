import { addEventOnPosts, handleScroll } from "./posts.js";
let path = "api/filter"+window.location.pathname;
console.log(path);

addEventOnPosts(path);

window.addEventListener("scrollend", () => {
  setInterval(handleScroll(path));
});

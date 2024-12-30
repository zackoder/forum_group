import { addEventOnPosts, handleScroll } from "./posts.js";
let path = "api/"+window.location.pathname;
path = "/api/categories/filter/1";
console.log(path);

addEventOnPosts(path);

window.addEventListener("scrollend", () => {
  setInterval(handleScroll(path));
});

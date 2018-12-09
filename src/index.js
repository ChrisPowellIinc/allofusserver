import m from "mithril";

// Bootstrap
import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import "izitoast/dist/css/iziToast.min.css";

// Pages
import Landing from "pages/landing";
import Register from "pages/auth/register";
import Login from "pages/auth/login";

const root = document.getElementById("app");

m.route.prefix("");

m.route(root, "/", {
  "/": Landing,
  "/register": Register,
  "/login": Login,
  // Lazy load 404 page, use this method to lazy load other pages
  "/:404": {
    onmatch: () =>
      new Promise(resolve => {
        console.log("what was you thinking...");
        return resolve;
      })
  }
});

import m from "mithril";

// Bootstrap
import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import "izitoast/dist/css/iziToast.min.css";

// Pages
import _404 from "pages/404";
import Landing from "pages/landing";
import Register from "pages/auth/register";
import Login from "pages/auth/login";

import Profile from "pages/profile";

const root = document.getElementById("app");

m.route.prefix("");

m.route(root, "/", {
  "/": Landing,
  "/register": Register,
  "/login": Login,
  "/profile": {
    onmatch: (args, requestedPath) => {
      console.log("I am going to check out login auth here..");
      return new Promise(resolve => resolve(Profile));
    }
  },
  // Lazy load 404 page, use this method to lazy load other pages
  "/:404": {
    onmatch: (args, requestedPath) => new Promise(resolve => resolve(_404))
  }
});

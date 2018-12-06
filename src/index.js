import m from "mithril";

// Pages
import Landing from "pages/landing";

const root = document.getElementById("app");

m.route.prefix("");

m.route(root, "/", {
  "/": Landing,

  // Lazy load 404 page, use this method to lazy load other pages
  "/:404": {
    onmatch: () =>
      new Promise(resolve => {
        console.log("what was you thinking...");
        return resolve;
      })
  }
});

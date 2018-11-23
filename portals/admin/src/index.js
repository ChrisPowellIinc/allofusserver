import m from "mithril";

// Global stylesheet
import "./styles/custom.scss";

// Pages
import Landing from "./pages/landing";
import _404 from "./pages/404";

m.route.prefix("");

m.route(document.body, "/404", {
  "/": Landing,
  "/*": _404
});
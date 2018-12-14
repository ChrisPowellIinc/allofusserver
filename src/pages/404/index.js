import m from "mithril";
import "css/404.css";

const _404 = {
  view: () => (
    <section class="error_section">
      <p class="error_section_subtitle">Opps Page is not available !</p>
      <h1 class="error_title">
        <p>404</p>
        404
      </h1>
      <a href="/" class="btn btn-primary" oncreate={m.route.link}>
        Back to home
      </a>
    </section>
  )
};

export default _404;

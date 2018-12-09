import m from "mithril";
import "css/cover.css";

const Landing = {
  view: () => (
    <div class="landing text-center">
      <div class="cover-container d-flex w-100 h-100 p-3 mx-auto flex-column">
        <header class="masthead mb-auto">
          <div class="inner">
            <h3 class="masthead-brand">All of Us</h3>
            <nav class="nav nav-masthead justify-content-center">
              <a class="nav-link active" href="#">
                Home
              </a>
              <a class="nav-link" href="#">
                Features
              </a>
              <a class="nav-link" href="#">
                Contact
              </a>
            </nav>
          </div>
        </header>

        <main role="main" class="inner cover">
          <h1 class="cover-heading">All of Us.</h1>
          <p class="lead">All of us is a social media network for ...</p>
          <p class="lead">
            <a
              href="/login"
              oncreate={m.route.link}
              class="btn btn-lg btn-secondary mr-1"
            >
              Login
            </a>
            <a
              href="/register"
              oncreate={m.route.link}
              class="btn btn-lg btn-secondary ml-1"
            >
              Signup
            </a>
          </p>
        </main>

        <footer class="mastfoot mt-auto">
          <div class="inner">
            <p>
              &copy; 2018{" "}
              <a class="text-white" href="/">
                allofus
              </a>
              , by{" "}
              <a class="text-white" href="/">
                allofus
              </a>
              .
            </p>
          </div>
        </footer>
      </div>
    </div>
  )
};

export default Landing;

import m from "mithril";
import "css/auth.css";
import Auth from "services/auth";

const Login = {
  User: {},
  login: e => {
    e.preventDefault();
    Auth.login(Login.User).catch(err => {});
  },
  oncreate: vnode => {
    vnode.state.errors = [];
  },
  view: vnode => (
    <section class="register">
      <form
        class="form-signin"
        onsubmit={Login.login}
        autocomplete="off"
        novalidate
      >
        <div class="text-center mb-4">
          <img
            class="mb-4"
            src="assets/img/logo.png"
            alt=""
            width="72"
            height="72"
          />
          <h1 class="h3 mb-3 font-weight-normal">Login</h1>
          <p>
            Login here and have a great experience. Or you don't have an
            account?{" "}
            <a href="/register" oncreate={m.route.link} class="">
              register here
            </a>
          </p>
        </div>

        <div class="form-label-group">
          <input
            type="text"
            id="username"
            name="username"
            class="form-control"
            placeholder="Username"
            required
            oninput={m.withAttr("value", value => {
              Login.User.username = value;
            })}
          />
          <label for="username">Username</label>
          {Auth.errors.username && (
            <small class="form-text text-danger">{Auth.errors.username}</small>
          )}
        </div>

        <div class="form-label-group">
          <input
            type="password"
            id="password"
            name="password"
            class="form-control"
            placeholder="Password"
            required
            oninput={m.withAttr("value", value => {
              Login.User.password = value;
            })}
          />
          <label for="password">Password</label>
          {Auth.errors.password && (
            <small class="form-text text-danger">{Auth.errors.password}</small>
          )}
        </div>

        <div class="checkbox mb-3">
          <label>
            <input type="checkbox" value="remember-me" /> Remember me
          </label>
        </div>
        <button class="btn btn-lg btn-primary btn-block" type="submit">
          Sign in
        </button>
        <p class="mt-5 mb-3 text-muted text-center">&copy; 2017-2018</p>
      </form>
    </section>
  )
};

export default Login;

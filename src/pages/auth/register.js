import m from "mithril";
import "css/auth.css";
import Auth from "services/auth";

const Register = {
  User: {},
  register: e => {
    e.preventDefault();
    Auth.register(Register.User).catch(err => {});
  },
  oncreate: vnode => {
    vnode.state.errors = [];
  },
  view: vnode => (
    <section class="register">
      <form
        class="form-signin"
        onsubmit={Register.register}
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
          <h1 class="h3 mb-3 font-weight-normal">Register</h1>
          <p>
            You don't have an account with us? register here and have a great
            experience. Or you have an account already?{" "}
            <a href="/login" oncreate={m.route.link} class="">
              sign in here
            </a>
          </p>
        </div>

        <div class="form-label-group">
          <input
            type="text"
            name="first_name"
            id="first_name"
            class="form-control"
            placeholder="First name"
            required
            oninput={m.withAttr("value", value => {
              Register.User.first_name = value;
            })}
          />
          <label for="first_name">First name</label>
          {Auth.errors.first_name && (
            <small class="form-text text-danger">
              {Auth.errors.first_name}
            </small>
          )}
        </div>

        <div class="form-label-group">
          <input
            type="text"
            name="last_name"
            id="last_name"
            class="form-control"
            placeholder="Last name"
            required
            oninput={m.withAttr("value", value => {
              Register.User.last_name = value;
            })}
          />
          <label for="last_name">Last name</label>
          {Auth.errors.last_name && (
            <small class="form-text text-danger">{Auth.errors.last_name}</small>
          )}
        </div>

        <div class="form-label-group">
          <input
            type="text"
            name="phone"
            id="phone"
            class="form-control"
            placeholder="Phone number"
            required
            oninput={m.withAttr("value", value => {
              Register.User.phone = value;
            })}
          />
          <label for="phone">Phone number</label>
          {Auth.errors.phone && (
            <small class="form-text text-danger">{Auth.errors.phone}</small>
          )}
        </div>

        <div class="form-label-group">
          <input
            type="email"
            name="email"
            id="email"
            class="form-control"
            placeholder="Email address"
            required
            oninput={m.withAttr("value", value => {
              Register.User.email = value;
            })}
          />
          <label for="email">Email address</label>
          {Auth.errors.email && (
            <small class="form-text text-danger">{Auth.errors.email}</small>
          )}
        </div>

        <div class="form-label-group">
          <input
            type="text"
            name="username"
            id="username"
            class="form-control"
            placeholder="Username"
            required
            oninput={m.withAttr("value", value => {
              Register.User.username = value;
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
            name="password"
            id="password"
            class="form-control"
            placeholder="Password"
            required
            oninput={m.withAttr("value", value => {
              Register.User.password = value;
            })}
          />
          <label for="password">Password</label>
          {Auth.errors.password && (
            <small class="form-text text-danger">{Auth.errors.password}</small>
          )}
        </div>

        <div class="form-label-group">
          <input
            type="password"
            name="confirm_password"
            id="confirm_password"
            class="form-control"
            placeholder="Confirm Password"
            required
            oninput={m.withAttr("value", value => {
              Register.User.confirm_password = value;
            })}
          />
          <label for="confirm_password">Confirm Password</label>
          {Auth.errors.confirm_password && (
            <small class="form-text text-danger">
              {Auth.errors.confirm_password}
            </small>
          )}
        </div>

        <button class="btn btn-lg btn-primary btn-block" type="submit">
          Register
        </button>
        <p class="mt-5 mb-3 text-muted text-center">&copy; 2017-2018</p>
      </form>
    </section>
  )
};

export default Register;

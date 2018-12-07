import m from "mithril";
import "css/auth.css";

const Register = {
  view: vnode => (
    <section class="register">
      <form class="form-signin">
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
            id="first_name"
            class="form-control"
            placeholder="First name"
            required
            autofocus
          />
          <label for="name">First name</label>
        </div>

        <div class="form-label-group">
          <input
            type="text"
            id="last_name"
            class="form-control"
            placeholder="Last name"
            required
            autofocus
          />
          <label for="name">Last name</label>
        </div>

        <div class="form-label-group">
          <input
            type="text"
            id="phone"
            class="form-control"
            placeholder="Phone number"
            required
            autofocus
          />
          <label for="phone">Phone number</label>
        </div>

        <div class="form-label-group">
          <input
            type="email"
            id="inputEmail"
            class="form-control"
            placeholder="Email address"
            required
            autofocus
          />
          <label for="inputEmail">Email address</label>
        </div>

        <div class="form-label-group">
          <input
            type="text"
            id="username"
            class="form-control"
            placeholder="Username"
            required
            autofocus
          />
          <label for="username">Username</label>
        </div>

        <div class="form-label-group">
          <input
            type="password"
            id="inputPassword"
            class="form-control"
            placeholder="Password"
            required
          />
          <label for="inputPassword">Password</label>
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

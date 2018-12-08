import m from "mithril";
import Auth from "services/auth";

export var request = options =>
  Auth.GetUserFromStorage().then(() => {
    options.headers = {
      Authorization: `Bearer ${Auth.user.token}`
    };
    options.extract = (xhr, opt) => {
      if (xhr.status === 401) {
        Auth.logout();
        return;
      }
      var response;
      if (options.old_extract) {
        response = options.old_extract(xhr, opt);
      } else {
        response = JSON.parse(xhr.responseText);
      }
      if (typeof response === "object") {
        response.status_code = xhr.status;
      }
      return response;
    };
    return m
      .request(options)
      .then(response => response)
      .catch(err => err);
  });

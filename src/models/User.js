import * as Joi from "joi-browser";

export const LoginSchema = Joi.object().keys({
  username: Joi.string()
    .trim()
    .required()
    .label("Username")
    .options({
      language: {
        key: "{{!label}} "
      }
    }),
  password: Joi.string()
    .required()
    .label("Password")
    .options({
      language: {
        key: "{{!label}} "
      }
    })
});

export const UserSchema = Joi.object().keys({
  first_name: Joi.string()
    .trim()
    .required()
    .label("First Name")
    .options({
      language: {
        key: "{{!label}} "
      }
    }),
  last_name: Joi.string()
    .trim()
    .required()
    .label("Last Name")
    .options({
      language: {
        key: "{{!label}} "
      }
    }),
  phone: Joi.string()
    .trim()
    .required()
    .label("Phone Number")
    .options({
      language: {
        key: "{{!label}} "
      }
    }),
  email: Joi.string()
    .trim()
    .email()
    .required()
    .label("Email")
    .options({
      language: {
        key: "{{!label}} "
      }
    }),
  username: Joi.string()
    .trim()
    .alphanum()
    .min(3)
    .max(30)
    .required()
    .label("Username")
    .options({
      language: {
        key: "{{!label}} "
      }
    }),
  password: Joi.string()
    .label("Password")
    .min(6)
    .required()
    .options({
      language: {
        key: "{{!label}} "
      }
    }),
  confirm_password: Joi.string()
    .label("Password Confirmation")
    .min(6)
    .required()
    .valid(Joi.ref("password"))
    .options({
      language: {
        any: {
          allowOnly: "!!Passwords do not match"
        },
        key: "{{!label}} "
      }
    })
});

export const User = {
  first_name: "",
  last_name: "",
  phone: "",
  email: "",
  username: "",
  password: ""
};

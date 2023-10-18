import { fail, type Actions, redirect } from "@sveltejs/kit";
import IsValidEmail from "../../../utils/emailPattern";
import * as Exec from "../../../utils/exec";

export const actions: Actions = {
  login: async ({ request, cookies }) => {
    const formData = await request.formData();

    const { email, password } = Object.fromEntries(formData) as {
      email: string;
      password: string;
    };

    const errors: any = {};

    if (!email || email.trim().length === 0) {
      errors.email = "L'adresse mail est requise";
    } else if (!IsValidEmail(email)) {
      errors.email = "Adresse mail invalide";
    }

    if (!password || password.trim().length === 0) {
      errors.password = "Le mot de passe est requis";
    }

    if (Object.keys(errors).length > 0) {
      const data = {
        errors,
      };
      return fail(400, data);
    }

    const url: string = "http://localhost:8090/api/auth/login";
    const data = Object.assign({ email, password }),
      method: Exec.RequestWithMethod<"POST"> = { Method: "POST" };

    const response = await Exec.default(url, method, data);
    const jsonFromResponse = await response.json();

    if(response.ok) {
      console.log(jsonFromResponse);
      cookies.set("session", jsonFromResponse.session, {
        httpOnly: true,
        secure: false,
        maxAge: 60 * 60 * 24 * 7,
        path: "/",
      })
      cookies.set("userid", jsonFromResponse.userid, {
        httpOnly: true,
        secure: false,
        maxAge: 60 * 60 * 24 * 7,
        path: "/",
      })

      throw redirect(303, "/")
    } else {
      errors.response = jsonFromResponse.message;
      return fail(response.status, Object.assign({errors}));
    }
  },
};

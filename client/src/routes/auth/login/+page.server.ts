import { fail, type Actions } from "@sveltejs/kit";
import IsValidEmail from "../../../utils/emailPattern";
import * as Exec from "../../../utils/exec";

export const actions: Actions = {
  login: async ({ request }) => {
    const formData= await request.formData();

    const { email, password } = Object.fromEntries(formData) as {
      email: string;
      password: string;
    };

    const errors: Record<string, unknown> = {};

    if (!email || email.trim().length === 0) {
      errors.email = "L'adresse mail est requise";
    }

    if (!IsValidEmail(email)) {
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
    const   data = Object.assign({email, password}),
            method: Exec.RequestWithMethod<"POST"> = {Method: "POST"}

    const response = await Exec.default(url, method, data);

    console.log(response);

    // if(Object.keys(errors).length ! == 0) {
    //     return {
    //         status: 400,
    //         body: {
    //             errors
    //         }
    //     }
    // }
  },
};

import type { Actions } from "@sveltejs/kit";
import { fail } from "@sveltejs/kit";
import EmailPattern from "../../../utils/emailPattern";
import IsNull, { type IsNullError } from "../../../utils/isNull";
import IsPasswordValid, {
  ArePasswordMatching,
} from "../../../utils/passwordValidation";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ url }) => {};

export const actions: Actions = {
  register: async ({ request }: { request: any }) => {
    const formData = await request.formData();
    const user = Object.fromEntries(formData);

    let { step = user.step, ...userData } = user;
    step = Number(step);

    const errors: Record<string, unknown> = {};
    console.log(step);
    let requiredFields: string[] = [];
    if (step === 1) {
      requiredFields = ["email", "firstname", "lastname"];

      for (const field of requiredFields) {
        const err: IsNullError | null = IsNull(user[field]);
        if (err) {
          errors[field] = err.Message;
        }

        if (field == "email") {
          const isValidEmail: boolean = EmailPattern(user[field]);
          if (!isValidEmail) {
            errors[field] = "L'email n'est pas valide";
          }
        }
      }
      if (Object.keys(errors).length > 0) {
        return fail(405, errors);
      }
      return {
        step: step + 1,
      };
    } else {
      requiredFields = ["password", "password_confirm"];
      for (let field of requiredFields) {
        let err: IsNullError | null;
        err = IsNull(user[field]);
        if (err) {
          errors[field] = err.Message;
        }

        if ((field = "password")) {
          err = IsPasswordValid(user[field]);
          if (err) {
            errors[field] = err.Message;
          }
        }
        err = ArePasswordMatching(user.password, user.password_confirm);

        if (err) {
          errors[field] = err.Message;
        }
      }
      if (Object.keys(errors).length != 0) {
        console.log("J'ai une erreur");
        return {
          errors,
          step: step,
        };
      }

      return {
        step: 3,
      };
    }
  },
  step: async ({ request }: { request: any }) => {
    const formData = await request.formData();
    const user = Object.fromEntries(formData);

    let step = Number(user.step);

    return {
      step: step - 1,
    };
  },
};

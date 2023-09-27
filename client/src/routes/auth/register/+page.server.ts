import { fail, type Actions, redirect } from "@sveltejs/kit";
import IsValidEmail from "../../../utils/emailPattern";
import * as PasswordValidation from "../../../utils/passwordValidation";
import * as Exec from "../../../utils/exec";
import type { PageLoad } from "../../$types";
import PhoneNumberByCountry from "../../../utils/PhoneNumberByCountry.json";
import { SortCountries } from "../../../utils/sortCountries";

let store: object = {};

export const load: PageLoad = async () => {
  const sortedCountries: Record<string, string>[] = SortCountries(PhoneNumberByCountry);
  return Object.assign({ countries: sortedCountries });
};

export const actions: Actions = {
  register: async ({ request }) => {
    const formData: FormData = await request.formData();
    const { step, ...user } = Object.fromEntries(formData) as {
      step: FormDataEntryValue;
      firstname: string;
      lastname: string;
      email: string;
      password: string;
      passwordconfirm: string;
    };

    console.log(formData);

    const stepAsNumber: number = Number(step);

    const errors: Record<string, unknown> = {};
    const allowedInputs: string[] =
      stepAsNumber === 1
        ? ["firstname", "lastname", "email"]
        : ["password", "passwordconfirm"];
    if (stepAsNumber <= 2) {
      for (const [key, value] of Object.entries(user)) {
        if (!allowedInputs.includes(key)) {
          errors[key] = "Champs non autorisé" + key;
        }
        if (!value || value.trim().length === 0) {
          errors[key] = "Ce champs est requis";
        }

        if (stepAsNumber === 1) {
          if (key === "email" && value.trim().length !== 0) {
            if (!IsValidEmail(value)) {
              errors[key] = "Format invalide";
            }
          }
          store = {
            ...store,
            ...user,
          };
        } else if (stepAsNumber === 2) {
          let err: PasswordValidation.PasswordValidationError | null = null;

          if (key === "password" && value.trim().length !== 0) {
            err = PasswordValidation.IsPasswordValid(value);
            if (err) {
              errors[key] = err.Message;
            }
          } else if (key === "passwordconfirm") {
            const password: FormDataEntryValue = user["password"];
            err = PasswordValidation.ArePasswordMatching(
              password.toString(),
              value
            );
            if (err) {
              errors[key] = err.Message;
            }
          }
          store = {
            ...store,
            ...user,
          };
        }
      }

      if (Object.keys(errors).length !== 0) {
        const data = {
          step: stepAsNumber,
          errors,
        };
        return fail(400, data);
      }

      console.log("Passed");
      return Object(Object.assign({ step: stepAsNumber + 1 }));
    } else {
      if (
        Object.keys(store).length < 5 ||
        Object.values(store).some((v): boolean => String(v).trim().length === 0)
      )
        throw new Error("Champs manquants");

      const url: string = "http://localhost:8090/api/auth/register",
        method: Exec.RequestWithMethod<"POST"> = {
          Method: "POST",
        };

      const response = await Exec.default(url, method, store);
      const data = await response.json();

      if (!response.ok) {
        return fail(response.status, data);
      }

      throw redirect(303, "/login");
    }
  },
};

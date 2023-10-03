import { fail, type Actions, redirect, error } from "@sveltejs/kit";
import IsValidEmail from "../../../utils/emailPattern";
import * as PasswordValidation from "../../../utils/passwordValidation";
import * as Exec from "../../../utils/exec";
import type { PageLoad } from "../../$types";
import PhoneNumberByCountry from "../../../utils/PhoneNumberByCountry.json";
import { SortCountries } from "../../../utils/sortCountries";
import ValidatePhoneNumber from "../../../utils/isValidPhoneNumber";

interface Country {
  id: number;
  name: string;
  country_code: string;
  phone_number: string;
  logo: string;
}

let store: object = {};

export const load: PageLoad = async () => {
  const sortedCountries: Record<string, string | any>[] =
    SortCountries(PhoneNumberByCountry);
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
      phone_number: string;
      phone_number_country_prefix: string;
      password: string;
      passwordconfirm: string;
      terms: string;
    };

    const stepAsNumber: number = Number(step);

    const errors: Record<string, unknown> = {};
    const allowedInputs: string[] =
      stepAsNumber === 1
        ? [
            "firstname",
            "lastname",
            "email",
            "phone_number",
            "phone_number_country_prefix",
          ]
        : stepAsNumber === 2
        ? ["password", "passwordconfirm"]
        : [
            "firstname",
            "lastname",
            "email",
            "password",
            "phone_number",
            "terms",
          ];

    for (const [key, value] of Object.entries(user)) {
      if (!allowedInputs.includes(key)) {
        errors[key] = "Champs non autorisé: " + key;
      }
      if (!value || value.trim().length === 0) {
        errors[key] = "Ce champs est requis";
      }

      if (stepAsNumber === 1) {
        if (key === "email" && value.trim().length !== 0) {
          if (!IsValidEmail(value)) {
            errors[key] = "Format invalide";
          }
        } else if (key === "phone_number") {
          const phone_number_country_prefix: string =
            user["phone_number_country_prefix"];
          const phone_number: string = user["phone_number"];
          if (
            !phone_number_country_prefix ||
            phone_number.trim().length === 0
          ) {
            errors[key] = "Ce champs requis";
          }
          const isValidPhoneNumber: boolean = ValidatePhoneNumber(phone_number);
          if (!isValidPhoneNumber) {
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
      } else if (stepAsNumber === 3) {
        if (user.terms !== "true") {
          errors["terms"] = "Vous devez accepter les conditions d'utilisation";
        } else if (Object.keys(store).length !== 8) {
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
    } else {
      if (stepAsNumber === 3) {
        const { terms, phone_number_country_prefix,  ...newStore } = store as { terms: boolean, phone_number_country_prefix: string };
        console.log(newStore);
        const response = await Exec.default(
          "http://localhost:8090/api/auth/register",
          { Method: "POST" },
          {
            ...newStore,
          }
        );

        const data = await response.json();
        console.log(data);


        if (response.status !== 200) {
          throw redirect(303, "/auth/register");
        }
      }
    }
    if (stepAsNumber !== 3) {
      return Object(Object.assign({ step: stepAsNumber + 1 }));
    }
  },
};

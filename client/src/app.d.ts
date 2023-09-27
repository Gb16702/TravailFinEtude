// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
  namespace App {
    // interface Error {}
    // interface Locals {}
    // interface PageData {}
    // interface Platform {}
  }

  type formValueType = "text" | "password" | "email" | "number" | "tel" | null;

  interface formValue {
    type: formValueType;
    placeholder: string;
    name: string;
    label: string;
    value: string;
  }

  type FormDataObject = {
    [key in
      | "firstName"
      | "lastName"
      | "email"
      | "phone_number"
      | "password"
      | "password_confirm"]: formValue;
  };
}

export {};

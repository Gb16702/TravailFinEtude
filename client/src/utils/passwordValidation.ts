const minLength = 4;
const maxLength = 32;

export type PasswordValidationError = {
  Message: string;
};

/**
 *  This function is used to validate a password by using a closure with two sub-functions.
 *
 * @param password string
 * @returns err | null
 */
const IsPasswordValid: (password: string) => PasswordValidationError | null = (
  password: string
) => {
  let err: PasswordValidationError | null = null;
  if (isPasswordTooShort(password)) {
    err = { Message: "Le mot de passe est trop court" };
  }

  if (isPasswordTooLong(password)) {
    err = { Message: "Le mot de passe est trop long" };
  }

  return err ? err : null;
};

const isPasswordTooShort: (password: string) => boolean = (
  password: string
) => {
  if (password.length < minLength) {
    return true;
  }

  return false;
};

const isPasswordTooLong: (password: string) => boolean = (password: string) => {
  if (password.length > maxLength) {
    return true;
  }

  return false;
};

/**
 *
 * This function is used to check if the password and the password_confirm are matching
 *
 * @param password
 * @param password_confirm
 * @returns err | null
 */
export const ArePasswordMatching: (password: string, password_confirm: string) => PasswordValidationError | null = (password: string, password_confirm: string) => {
    let err: PasswordValidationError | null = null;
    if(String(password).trim() != String(password_confirm).trim()) {
        err = {Message: "Les mots de passe ne correspondent pas"}
    }

    return err ? err : null;
}

export default IsPasswordValid;

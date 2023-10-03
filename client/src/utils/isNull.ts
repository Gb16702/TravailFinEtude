export type IsNullError = {
  Message: string;
};

/**
 * Checks if a string is null or empty
 *
 * @param params string
 * @returns {boolean}
 */
const IsNull: (params: string) => IsNullError | null = (params: string) => {
  console.log(typeof params);

  let err: IsNullError | null = null;
  if (params == null || params.trim().length == 0) {
    err = { Message: "Ce champs est requis" };
  }

  return err ? err : null;
};

export default IsNull;

const emailPattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;

/**
 *
 * @param email
 * @returns {boolean}
 */
const IsValidEmail: (email: string) => boolean = (email: string): boolean => emailPattern.test(email)

export default IsValidEmail;

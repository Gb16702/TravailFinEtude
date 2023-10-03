/**
 * Generate a secure password
 *
 * @param length
 * @returns {string}
 */
const GenerateSecurePassword: (length: number) => string = (length) => {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+~`|}{[]:;?><,./-=";
    let password: string = "";
    for(let i: number= 0; i <= length; ++i) {
        password += charset.charAt(Math.floor(Math.random() * charset.length));
    }

    return password;
};

export default GenerateSecurePassword;

/**
 * This function is used to check if the phone number is valid or not
 *
 * @param phone_number string
 * @returns {boolean}
 */
const ValidatePhoneNumber: (phone_number: string) => boolean = (phone_number) => {
    if(typeof phone_number !== "string") {
        return false;
    }
    const phoneRegex = /^\d{7,15}$/
    return phoneRegex.test(phone_number);
}

export default ValidatePhoneNumber;
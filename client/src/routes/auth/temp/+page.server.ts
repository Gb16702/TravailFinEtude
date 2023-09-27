import type { PageLoad } from "../../$types";
import PhoneNumberByCountry from "../../../utils/PhoneNumberByCountry.json";
import { SortCountries } from "../../../utils/sortCountries";

export const load: PageLoad = async () => {
    const sortedCountries: Record<string, string>[] = SortCountries(PhoneNumberByCountry);
    return Object.assign({ countries: sortedCountries });
};
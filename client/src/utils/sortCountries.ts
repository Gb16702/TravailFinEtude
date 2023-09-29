/**
 *  This function sorts the countries in alphabetical order
 *
 * @param obj Record<string, string>
 * @returns Array<Record<string, string>>
 */
export const SortCountries: (obj: Record<string, string>[]) => Record<string, string>[] = (obj: Record<string, string>[]) => {
  return obj.sort(
    (a: any, b: any) => a.name.localeCompare(b.name) as -1 | 0 | 1
  );
};

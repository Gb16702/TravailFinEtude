/**
 *  This function sorts the countries in alphabetical order
 *
 * @param obj Record<string, string>
 * @returns Array<Record<string, string>>
 */
export const SortCountries: (obj: Record<string, string>) => Record<string, string>[] = (obj) => {
  return Object.entries(obj)
  .sort(([x]:[string, string], [y]:[string, string]): number => x.localeCompare(y) as -1 | 0 | 1)
  .map(([key, value]) => ({name: key, code: value}));
}

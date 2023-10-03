import { writable, type Writable } from "svelte/store";

export type Country = {
  id:   number | null;
  code: string | null;
  logo: string;
};

export const countryStore: Writable<Record<string, string | null>> =
  writable({
    code: "+32",
    logo: "🇧🇪",
});

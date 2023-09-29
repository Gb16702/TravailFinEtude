import { writable, type Writable } from "svelte/store";

export const countryStore: Writable<string | null> = writable("+32");

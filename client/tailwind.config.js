/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,svelte,ts}"],
  theme: {
    extend: {
      colors: {
        background: "#fff",
        foreground: "#000",
        primary: "#363636",
        secondary: "#DADADA",
        accent: {
          stone: "#6E6E7C",
          "stone-b": "#737387",
          "stone-c": "#706F71",
          "stone-d": "#5c5d5b",
          green: "#5E786B",
          "green-dark": "#4F6A5D",
          brown: "#736464",
          lilas: "#DEDCFF",
          "violet-a": "#433BFF",
          "violet-b": "#827EEF",
          "violet-c": "#8B87FF",
          "violet-d": "#111",
          "slate-a": "#47525c",
          "dark-mode-main": "#151F2C",
          "dark-mode-secondary": "#000",
          "dark-mode-accent": "#3A22A1",
          "light-mode-a": "#9CC9FF",
          "light-mode-b": "#92A2FE",
        },
      },
    },
  },
  plugins: [],
};

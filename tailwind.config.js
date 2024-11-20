import { fontFamily } from "tailwindcss/defaultTheme";

/** @type {import('tailwindcss').Config} */
export default {
  darkMode: ["class"],
  content: ["./templates/layouts/*.html", "./templates/pages/*.html", "./assets/static/css/*.css"],
  extend: {
    fontFamily: {
      serif: ['Lora Variable', 'sans-serif'],
      sans: [
        "Nunito Sans",
        ...fontFamily.sans
      ],
    }
  }
}

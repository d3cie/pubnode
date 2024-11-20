import { fontFamily } from "tailwindcss/defaultTheme";

/** @type {import('tailwindcss').Config} */
export default {
  darkMode: ["class"],
  content: ["./templates/layouts/*.html", "./templates/pages/*.html", "./assets/static/css/*.css"],
  theme: {
    extend: {
      colors: {
        "gray": {
          50: '#F8F8F7',
          100: '#EFEDEC',
          200: '#DDD8D5',
          300: '#CBC4BF',
          400: '#B8AFA8',
          500: '#A69A92',
          600: '#93857C',
          700: '#766A61',
          800: '#574E47',
          900: '#38322E',
          950: '#292522'
        },
      },
      fontFamily: {
        serif: ['Lora Variable', 'sans-serif'],
        sans: [
          "Nunito Sans",
          ...fontFamily.sans
        ],
      }
    }
  }
}

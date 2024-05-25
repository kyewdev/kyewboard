/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.go", "./**/*.go],
    safelist: [],
    plugins: [require("daisyui")],
    daisyui: {
        themes: ["dark"]
    }
}


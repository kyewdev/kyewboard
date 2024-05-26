/** @type {import('tailwindcss').Config} */
    module.exports = {
        content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
         theme: {
            extend: {
                fontFamily: {
                    morpheus: ['Morpheus', 'serif'],
                },
                colors: {
                    questText: '#DAA520', // Color similar to WoW quest text
                },
            },
         },
        safelist: [],
    }


#!/bin/zsh

case "$1" in
    "templ")
        wezterm cli spawn -- templ generate --watch --proxy=http://localhost:42069
        ;;
    "air")
        wezterm cli spawn -- air
        ;;
    "tailwind")
        wezterm cli spawn -- npx tailwindcss -i ./pkg/view/app.css -o ./assets/style.css --watch
        ;;
    "all")
        wezterm cli spawn -- air
        wezterm cli spawn -- npx tailwindcss -i ./pkg/view/app.css -o ./assets/style.css --watch
        wezterm cli spawn -- templ generate --watch --proxy=http://localhost:42069
        ;;
    *)
        echo "Invalid argument. Please use templ, air, tailwind, or all."
        ;;
esac

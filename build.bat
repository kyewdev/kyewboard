@echo off
setlocal

if "%1" == "templ" (
    start cmd /k "templ generate -watch -proxy=http://localhost:42069"
) else if "%1" == "air" (
    start cmd /k "air"
) else if "%1" == "tailwind" (
    start cmd /k "npx tailwindcss -i .\pkg\view\app.css -o .\assets\style.css --watch"
) else if "%1" == "all" (
    start cmd /k "templ generate -watch -proxy=http://localhost:42069"
    start cmd /k "air"
    start cmd /k "npx tailwindcss -i .\pkg\view\app.css -o .\assets\style.css --watch"
)else (
    echo Invalid argument. Please use templ, air, or tailwind.
)

endlocal
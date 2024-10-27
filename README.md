# Fiber go server with templ, tailwindcss, FrankenUI, alpinejs and htmx (TEMPLATE)
1. fiber - https://gofiber.io/
1. templ templating - https://templ.guide/
1. tailwindcss - https://tailwindcss.com/
1. frankenUI components (remote dependency in [base_layout.templ](./views/layouts/base_layout.templ)) - https://franken-ui.dev/ 
1. alpinejs (remote dependency in [base_layout.templ](./views/layouts/base_layout.templ)) - https://alpinejs.dev/
1. htmx (remote dependency in [base_layout.templ](./views/layouts/base_layout.templ)) - https://htmx.org/

## Requirements
1. go >= 1.23.2, see [go.mod](./go.mod)
1. node.js - npm tailwind dependency
1. makefile runner (for hot reload with Makefile, optional) - on win [Cygwin](https://www.cygwin.com/) with package `make` and added to PATH is required

## How to run (via Makefile)
1. open repository root
1. run `make live` command
1. tailwindcss watch started
1. fiber webserver started at http://127.0.0.1:3000 with file watcher
1. templ proxy (hot reload) started at http://127.0.0.1:7331

> [!NOTE]
> Makefile uses AIR to detect file/templ changes and to restart the app.

> [!NOTE]
> Makefile on windows (terminal/pws window in VSCode) sometimes freezes when stopping with `CTRL+C` and needs to be killed.

## How to run (manually - does not require Makefile runner)
1. open repository root
1. run `templ generate` command - generates templ components
1. run `npm install`, then `npx tailwindcss -i assets/app.css -o ./public/app.css --minify` - generates css
1. run `go run .`

> [!NOTE]
> Tested only on windows, there should be no windows dependencies, but you were warn.

## How to make it yours
1. replace all `app.yourcompany` namespace occurences with `your own` namespace
1. you are ready to go

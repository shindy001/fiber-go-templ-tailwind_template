# Starts tailwindcss in watch mode to generate css, also calls 'npm install' to get tailwind package, requires NodeJS https://nodejs.org/
# Starts go server in watch mode - rebuils/rerun server on file change
# Starts templ generate in watch mode to detect all .templ files changes, default proxy (hot reload) runs at: http://localhost:7331

# Start this Makefile by running "make live"
# On Windows, Cygwin with package "make" and added 'C:\cygwin64\bin' folder to PATH is required to run this, https://www.cygwin.com/

npm-install-and-tailwind-watch:
	npm.cmd install
	npx.cmd tailwindcss -i assets/app.css -o ./public/app.css --minify --watch

# runs air to detect any go file changes to re-build and re-run the server.
server-watch-air:
	go run github.com/air-verse/air@v1.60.0 \
	--build.cmd "templ.exe generate --notify-proxy & go build -o ./tmp/main.exe ." \
	--build.bin "tmp\\main.exe" \
	--build.delay "100" \
	--build.rerun_delay "100" \
	--build.exclude_dir "node_modules" \
	--build.exclude_regex "_test.go,.*_templ.go" \
	--build.include_ext "go,tpl,tmpl,templ,html,js,css" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
templ-watch:
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false -v



# start all watch processes in parallel
live:
	make -j3 npm-install-and-tailwind-watch server-watch-air templ-watch


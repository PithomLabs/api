package lambdas

import (
	"fmt"
	"net/http"
)

// IndexHandler corresponds to the "/" endpoints
func IndexHandler(resp http.ResponseWriter, req *http.Request) {
	// Creating a Schema

	resp.Header().Set("Content-Type", "text/html")

	fmt.Fprint(resp, `
	<style>
	html, body {
		height: 100%;
		margin: 0
	}
	body {
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: column
	}
	h1 {
		font-size: calc(4vw + 2em)
	}
	</style>
	<h1>Komfy API root page.</h1>

	<h2>Docs: coming soon</h2>

	<h2>Komfy homepage: <a href="https://komfy.now.sh">here</a></h2>

	<h2>Github repo: <a href="https://github.com/komfy/api">here</a></h2>
	`)

}

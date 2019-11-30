package lambdas

import (
	"fmt"
	"net/http"
)

// IndexHandler is the root page of API for displaying useful info
func IndexHandler(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-Type", "text/html")

	fmt.Fprint(resp, `
	<style>
	html, body {
		height: 100%;
		margin: 0;
		background: rgb(33, 31, 48);
		
		font-family: monospace;
	}
	body {
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: column
	}
	svg {
		height: 10em;
	}
	* {
		color: white;
	}
	a {
		color: skyblue;
	}
	h1 {
		font-size: calc(3vw + 1.6em)
	}
	</style>
	<svg
		 xmlns="http://www.w3.org/2000/svg"
		 viewBox="0 0 1033 1033"
		 version="1.1"
		 id="svg177"
		 >
		<sodipodi:namedview
			 pagecolor="#ffffff"
			 bordercolor="#666666"
			 borderopacity="1"
			 id="namedview179"
			 showgrid="false"
			 showguides="true">
		</sodipodi:namedview>
		<defs
			 id="defs165">
			<style
				 id="style163">.cls-1{fill:#fff;stroke-linecap:round;}.cls-1,.cls-2{stroke:#000;stroke-miterlimit:10;stroke-width:50px;}.cls-2{fill:none;}</style>
		</defs>
		<title
			 id="title167">logo</title>
		<circle
			 class="cls-1"
			 cx="516.5"
			 cy="516.5"
			 r="491.5"
			 id="circle169"
			 style="stroke:none" />
		<path
			 d="M779.35477,380.46328H435.84645a129.02364,129.02364,0,1,0-.19227,63.9732H703.27854V549.90581h55.32817V444.43648h20.74806V549.90581h55.32817V380.46328ZM310.19924,490.95669a78.40038,78.40038,0,1,1,78.40033-78.40033A78.4004,78.4004,0,0,1,310.19924,490.95669Z"
			 transform="translate(4 5)"
			 id="path171"
			 style="fill:#232121;fill-opacity:1" />
		<path
			 class="cls-2"
			 d="M333.94166,796.886s81.0793,66.97855,188.5975,66.97855S727,796.886,727,796.886"
			 transform="translate(4 5)"
			 id="path173"
			 style="stroke:#232121;stroke-opacity:1" />
		<polygon
			 points="470 607 524 607 497 574 534 544 641 661 470 661 470 607"
			 id="polygon175"
			 style="fill:#232121;fill-opacity:1" />
	</svg>
	
	<h1>Komfy API root page.</h1>
	<div>
		<a href="https://komfy.now.sh/rand">Password generator</a> | 
		<a href="https://komfy.now.sh">Homepage</a> |

		<a href="https://github.com/komfy/api">Github</a> |
		<a href="https://t.me/komfy">Telegram</a>
	</div>

	`)

}

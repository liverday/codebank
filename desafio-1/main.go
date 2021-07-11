package main

import (
	"fmt"
	"net/http"
)	

func main() {
	http.HandleFunc("/", ServerHandler)
	http.ListenAndServe(":8000", nil)
}

func ServerHandler(writer http.ResponseWriter, r *http.Request) {
	content := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Imersão Full Cycle</title>
			<link rel="preconnect" href="https://fonts.googleapis.com">
			<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
			<link href="https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;700&display=swap" rel="stylesheet">
		
			<style>
				* {
					margin: 0;
					padding: 0;
					box-sizing: border-box;
				}
		
				body {
					font-family: 'Open Sans', sans-serif;
					height: 100vh;
					width: 100vw;
					display: flex;
					align-items: center;
					justify-content: center;
					background: url('https://images.unsplash.com/photo-1524481905007-ea072534b820?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1950&q=80') no-repeat center;
				}
		
				h1 {
					background: #363636;
					padding: 20px;
					border-radius: 8px;
					color: #fff;
					text-align: center;
				}
			</style>
		</head>
		<body>
			<h1>Imersão Full Cycle</h1>
		</body>
		</html>
	`

	fmt.Fprintf(writer, content)
}
package lib

import (
	"bytes"
	"fmt"
	"net/http"
)

func respond(w http.ResponseWriter, body []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func errorJSON(msg string) []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `{"error": "%s"}`, msg)
	return buf.Bytes()
}

type Playground struct{}

func (h Playground) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respond(w, errorJSON("only GET requests are supported"), http.StatusMethodNotAllowed)
		return
	}

	w.Write(graphiql)
}

var graphiql = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphql-playground-react@1.7.26/build/static/css/index.css"/>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/2.0.3/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/16.2.0/umd/react.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/16.2.0/umd/react-dom.production.min.js"></script>
		<script src="https://cdn.jsdelivr.net/npm/graphql-playground-react@1.7.26/build/static/js/middleware.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="root"/>
		<script type="text/javascript">
			window.addEventListener('load', function (event) {
				const root = document.getElementById('root');
				root.classList.add('playgroundIn');
				const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:'
				GraphQLPlayground.init(root, {
					endpoint: location.protocol + '//' + location.host + '/graphql',
					subscriptionsEndpoint: wsProto + '//' + location.host + '/graphql',
				   shareEnabled: true,
					settings: {
						'request.credentials': 'same-origin'
					}
				})
			})
		</script>
	</body>
</html>
`)
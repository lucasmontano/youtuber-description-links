module lucasmontano.com/yt-links

go 1.15

replace (
	lucasmontano.com/yt-links/models => ./models
	lucasmontano.com/yt-links/services => ./services
	lucasmontano.com/yt-links/tests => ./tests
)

require (
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	google.golang.org/api v0.32.0
)

init-golanggo:
	git clone --filter=blob:none --no-checkout git@github.com:golang/go.git _golanggo
	cd _golanggo && git sparse-checkout init --cone && git sparse-checkout add /src/encoding && git checkout go1.15.6

cp-golanggo:
	cp -R _golanggo/src/** internal/

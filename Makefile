PACKAGE = nautilus-print-server

build:
	mkdir -p build && \
	go build . && \
	cp ./scripts/* $(PACKAGE).service ./$(PACKAGE) ./build

tar: build
	tar -czvf print-server.tar.gz ./build && \
	rm -rf build

clean:
	rm -rf build && \
	rm -f print-server.tar.gz

install:
	./build/install.sh

uninstall:
	./build/uninstall.sh
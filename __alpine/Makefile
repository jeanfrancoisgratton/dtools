all: 
	cd ../src
	go build . -o ./dtools.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.19.3 amd64
	sudo apk update
	#TODO : update following line
	#sudo apk add --no-cache pkgconfig libvirt libvirt-dev gcc musl-dev


install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin

	install -pm 755 $(BUILDDIR)/dtools.exe $(DESTDIR)$(PREFIX)/bin/dtools


	

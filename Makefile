html: md2html
	./md2html
md2html: 
	cd gen; go build -o ../md2html; cd ..
clean:
	rm ./md2html; rm html/*  -rvf

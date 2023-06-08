.PHONY: pdf
pdf:
	xelatex essential-go.tex
	makeindex essential-go.idx
	xelatex essential-go.tex
	xelatex essential-go.tex

.PHONY: ebook
ebook: cover
	tex4ebook -c tex4ht.cfg -x essential-go.tex
	mkdir -p tmp
	cp essential-go.epub tmp
	cd tmp && unzip essential-go.epub && cd OEBPS && python3 ../../scripts/replace_svg.py
	cd tmp && zip -r ../essential-go-new.epub .
	ebook-convert essential-go-new.epub essential-go.epub --remove-first-image --cover cover/cover.png --authors "Baiju Muthukadan" --title "Essentials of Go Programming"
	ebook-convert essential-go-new.epub essential-go.mobi --remove-first-image --cover cover/cover.png --authors "Baiju Muthukadan" --title "Essentials of Go Programming"

.PHONY: html
html: html
	make4ht -x -d html essential-go.tex

.PHONY: cover
cover:
	cd cover && xelatex cover.tex
	cd cover && xelatex cover.tex
	cd cover && pdftoppm cover.pdf cover -png
	cd cover && convert cover-1.png -gravity West -chop 1132x0 cover.png
	convert cover/cover.png -resize "30%" cover/images/cover-small.png

.PHONY: sphinx
sphinx:
	pandoc -o sphinx/introduction.rst introduction.tex
	pandoc -o sphinx/quickstart.rst quickstart.tex
	pandoc -o sphinx/control-structures.rst control-structures.tex
	pandoc -o sphinx/data-structures.rst data-structures.tex
	pandoc -o sphinx/functions.rst functions.tex
	pandoc -o sphinx/objects.rst objects.tex
	pandoc -o sphinx/concurrency.rst concurrency.tex
	pandoc -o sphinx/packages.rst packages.tex
	pandoc -o sphinx/io.rst io.tex
	pandoc -o sphinx/testing.rst testing.tex
	pandoc -o sphinx/tooling.rst tooling.tex
	pandoc -o sphinx/answers.rst answers.tex
	pandoc -o sphinx/acknowledgements.rst acknowledgements.tex
	pandoc -o sphinx/colophon.rst colophon.tex
	pandoc -o sphinx/preface.rst preface.tex
	python3 scripts/code_block.py

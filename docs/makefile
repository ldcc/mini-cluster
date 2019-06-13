SRC = $(wildcard texs/*.tex)
PDFS = $(SRC:.tex=.pdf)

pdf: clean $(PDFS)

%.pdf: %.tex
	xelatex -output-directory ./pdfs $<
%.pdf: texs/%.tex
	xelatex -output-directory ./pdfs $<
%.svg: %.pdf
	inkscape -z -l ./svgs/$@ $<
%.svg: pdfs/%.pdf
	inkscape -z -l ./svgs/$@ $<

%.cls:
	rm -f ./pdfs/$*.*~ ./pdfs/$*.* ./svgs/$*.*

clean:
	rm -f ./pdfs/*
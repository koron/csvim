gui: run
	#gvim -u NONE -U NONE --noplugin -c "source preview.vim"
	gvim -u NONE -U NONE -c "source preview.vim"

.PHONY: run
run:
	go run . > monochromenote.vim

.PHONY: install
install: run
	@if [ -d ~/vimfiles ] ; then \
	  echo cp monochromenote.vim ~/vimfiles/colors/ ; \
	  cp monochromenote.vim ~/vimfiles/colors/ ; \
	fi
	@if [ -d ~/.vim ] ; then \
	  echo cp monochromenote.vim ~/.vim/colors/ ; \
	  cp monochromenote.vim ~/.vim/colors/ ; \
	fi

" vim:set ts=8 sts=2 sw=2 tw=0 et:
scriptencoding utf-8

syntax on
set foldcolumn=2
set conceallevel=1 listchars+=conceal:~

if has('gui_running')
  set gfn=Cica:h10.5
  set guioptions=
  winpos 0 0
  set lines=999
  set columns=165
endif

so blackandwhite.vim
so $VIMRUNTIME/syntax/hitest.vim
rightb vsp blackandwhite.vim
set nowrap number cursorline cursorcolumn

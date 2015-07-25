# transpose
transpose text files fast

A fast simple way to transpose delimited (single- or multi-character) text files. Accepts STDIN or a file name. Use -d to change the input delimiter and -D to change the output delimiter. Default is tab.

If you can make it "go" faster, please submit a pull request.

**Note:** works in memory. For larger files than can fit in memory, you can use:

cut -f 1-100 yourbigfile | transpose > file1

cut -f 101-200 yourbigfile | transpose > file2

...

cat file1 file2 . . . fileN > new_transposed_file

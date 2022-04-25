### grup

Benchmark numbers for searching 99k files (1.3GB).

Setup:
- Clone facebook/react@6d3b6d0f 
- Build it: `yarn build`
- Run each bench command five times to warm disk cache
- Run each bench command with `time` five times and take average (.000 accuracy)

Results (Apple M1 Pro 16GB RAM):

```
# all report the line number of the single text file match
# and that a single binary file has a match somewhere

rg -uuu -n packages/react-fetch react/  0.87s user 4.44s system 702% cpu 0.756 total
sift -n packages/react-fetch react/  1.55s user 4.33s system 688% cpu 0.854 total
grup -n packages/react-fetch react/  1.66s user 3.50s system 523% cpu 0.986 total
grep -r -n packages/react-fetch react/  8.63s user 1.32s system 99% cpu 9.954 total
``` 


Versions:
- fgrep (BSD grep, GNU compatible) 2.6.0-FreeBSD
- grep (BSD grep, GNU compatible) 2.6.0-FreeBSD
- grup (f6a04efd)
- ripgrep 13.0.0 -SIMD -AVX (compiled)
- sift 0.9.0 (darwin/arm64)

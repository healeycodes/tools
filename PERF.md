### grup

Some draft benchmark numbers on a Apple M1 Pro.

Searching: facebook/react (6d3b6d0f), 98915 files (a mix of text and binary), ~1.3GB.

Results:

```
grep -F -r -n packages/react-fetch react/  8.64s user 1.31s system 99% cpu 9.948 total
grup -n packages/react-fetch react/  0.85s user 1.41s system 107% cpu 2.091 total
rg -n packages/react-fetch react/  0.04s user 0.19s system 488% cpu 0.046 total
rg -uuu -n packages/react-fetch react/  0.87s user 4.08s system 717% cpu 0.690 total
``` 


`rg` clearly winning (even if you force it to search ignored files and binary files).

Surprising that `grup` is 2x as fast as `grep`.


Versions:
- fgrep (BSD grep, GNU compatible) 2.6.0-FreeBSD
- grep (BSD grep, GNU compatible) 2.6.0-FreeBSD
- grup (d5d43f90)
- ripgrep 13.0.0 -SIMD -AVX (compiled)
- sift 0.9.0 (darwin/arm64)

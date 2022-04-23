### grup

Some draft benchmark numbers on a Apple M1 Pro.

Searching: facebook/react (6d3b6d0f), 98915 files (a mix of text and binary), ~1.3GB.

Results:

- fgrep -r -n packages/react-fetch react/  8.77s user 1.98s system 67% cpu 15.852 total
- grep -r -n packages/react-fetch react/  8.67s user 1.93s system 69% cpu 15.256 total
- grup -n packages/react-fetch react/  0.96s user 2.10s system 41% cpu 7.436 total
- rg -n packages/react-fetch react/  0.03s user 0.15s system 357% cpu 0.050 total
- rg -uuu -n packages/react-fetch react/  0.80s user 3.49s system 269% cpu 1.595 total
- sift -n packages/react-fetch react/  1.45s user 3.44s system 254% cpu 1.918 total

`rg` clearly winning (even if you force it to search ignored files and binary files).

Surprising that `grup` is 2x as fast as `grep`.


Versions:
- fgrep (BSD grep, GNU compatible) 2.6.0-FreeBSD
- grep (BSD grep, GNU compatible) 2.6.0-FreeBSD
- grup (d5d43f90)
- ripgrep 13.0.0 -SIMD -AVX (compiled)
- sift 0.9.0 (darwin/arm64)

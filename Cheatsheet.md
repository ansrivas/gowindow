1. Avoid pointers if you can.
2. Return larger objects, pack them into structs ( works like wonders)
3. Slices should be used efficiently
  - tempslice := baseslice[:0]
  -
4. Size of empty struct = 0 ( just declare it `var s struct`)
5.
Avoid unnecessary heap allocations.
Prefer values over pointers for not big structures.
Preallocate maps and slices if you know the size beforehand.
Don't log if you don't have to.
Use buffered I/O if you do many sequential reads or writes.
If your application extensively uses JSON, consider utilizing parser/serializer generators (I personally prefer easyjson).
Every operation matters in a hot path.

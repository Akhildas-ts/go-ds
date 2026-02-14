1. What is the difference between an array and a slice in Go?
Arrays are value types
When you assign or pass an array, the entire array is copied

Dynamic size
Can grow or shrink using append
Under the hood, it uses an array


2. What does the cap() function return, and why does it matter for slices?


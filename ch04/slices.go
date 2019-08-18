package ch04

/*
 * Slices represent variable-length sequences whose elements all have
 * the same type. A slice type is written as []T, where the elements
 * have type T; it looks like an array type without a size.
 *
 * Arrays and slices are intimately connected. A slice is a lightweight
 * data structure that gives access to a subsequence (or perhaps all) of
 * the elements of an array, which is known as the slice’s underlying array.
 *
 * A slice has three components: a pointer, a length, and a capacity.
 *
 * The pointer points to the first element of the array that is reachable
 * through the slice, which is not necessarily the array’s first element.
 *
 * The length is the number of slice elements; it can’t exceed the capacity,
 * which is usually the number of elements between the start of the slice and
 * the end of the underlying array.
 *
 * The built-in functions len and cap return those values.
 */

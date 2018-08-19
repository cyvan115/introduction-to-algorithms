# 插入排序

> 针对某一数组 `A`，其长度为 `A.length` 。插入排序 **原址** 排序数组中的数，排序过程中最多只有常数个数字存储在数组外面。在排序结束时，输出数组 `A` 包含排序好的输出序列。

## 伪代码描述

``` None
  // ascending order
  for j = 2 to A.length
    key = A[j]
    // Insert A[j] into the sorted sequence A[1..j-1].
    i = j - 1
    while i > 0 and A[i] > key
      A[i + 1] = A[i]
      i = i - 1
    A[i + 1] = key
```

## 伪代码解读

首先，我们得先了解到 **插入排序** 的核心思想：对于即将要进行排序的数组元素 `A[j]` 来说，一定可以保证 `A[1..j-1]` 一定是已经排好序的，所以我们只需确定 `A[j]` 在数组 `A[1..j-1]` 中正确的位置就行。

伪代码中第一行 `for j = 2 to A.length` 的 `j` 代表着数组中的第`j`个元素，而不是我们习惯性认为的数组元素下标；针对于即将要插入到有序序列 `A[1..j-1]` 的 `A[j]` 来说，我们可以通过从当前位置向前遍历的方式找到 `A[j]` 的正确位置：当 `A[i] <= A[j] and A[j] <= A[i + 1]` 时，`A[j]` 应当插入到 `A[i]` 的后面，即 `A[i+1] = A[j]`，当然，在这之前我们得先将 `A[i+1..j]` 的所有元素向后移动一个单位，目的是为了给 `A[j]` 的插入腾出位置且避免先插入导致丢失 `A[i+1]` 的原始值（`A[j]` 的值早已保存在了变量 `key` 中，所以不用担心后移元素时丢失 `A[j]` 的值）。

## 额外说明

### 循环不变式

- 上述 `A[1..j-1]` 称为一个循环不变式：每一次 `for` 循环开始迭代时， `A[1..j-1]` 都是由原来的 `A[1..j-1]` 中的元素构成，但已按序排列。
- 循环不变式主要用来帮助理解算法的正确性，关于循环不变式，必须要证明三条性质：
  - **初始化**：循环的第一次迭代之前，它为真
  - **保持**：循环中某一次迭代它为真，则下一次也为真
  - **终止**：在循环终止时，不变式提供了一个有用的性质，该性质有助于证明算法是正确的

### 循环不变式性质证明

- 初始化：循环刚开始时 `j` 被赋值为2，也就是说此时的 `A[j]` 代指数组中的第二个元素，此时的“循环不变式”就是 `A[0]`，一个元素显然是有序的，得证
- 保持：保持在插入排序中代指的就是在某一次循环过程中，给 `A[j]` 在 `A[1..j-1]` 中寻找正确的插入位置，完成插入操作后列表仍有序；由于 `A[1..j-1]` 在插入前已经排好序了，所以在有序列表中按照规则插入一个数 `A[j]`，列表必然仍有序，得证
- 终止：循环终止时，`j` 被赋值为 `A.length`，指向了最后一个元素并完成了插入操作；此时已经遍历（正向遍历）了除第一个元素外的所有元素，所有的元素都按序排在了列表中，得证

## Go 代码实现

[插入排序的Go语言实现](./insertion_sort.go)
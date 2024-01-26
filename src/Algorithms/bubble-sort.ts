function sort(numbers: number[]) {
  for (let i = 0; i < numbers.length; i++) {
    for (let j = 0; j < numbers.length; j++) {
      const left = numbers[j];
      const right = numbers[j + 1];

      if (left > right) {
        numbers[j] = right;
        numbers[j + 1] = left;
      }
    }
  }
  return numbers;
}

console.log(sort([2, 1, 6, 5, 9, 7, 4, 3]));
console.log(sort([-8, 1, 3, 5, -3, 7, 2, 4]));
console.log(sort([10,2,-1,3,-22]));

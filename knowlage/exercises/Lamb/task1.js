console.log("Task 1");

let numbers = [
  [6, 4, 6, 7],
  [2, 5, 7, 7],
  [5, 7, 5, 7],
];

console.log(numbers);

let sum = 0;

for (i = 0; i < numbers.length; i++) {
  for (j = 0; j < numbers[i].length; j++) {
    if (numbers[i][j] % 2 == 0) {
      sum += numbers[i][j];
    }
  }
}

console.log("сумма четных чисел: ", sum);

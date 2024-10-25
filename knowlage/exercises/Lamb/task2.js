console.log("Hello! i am task 2");

console.log("отсортировать массив по возрастанию");
let numbers = [2, 5, 8, 4, 1, 9, 12, 5];
console.log(numbers);
let min = 1000;
let minNumbers = [];
let index = 0;
for (let j = 0; j < numbers.length; j++) {
  for (i = 0; i < numbers.length; i++) {
    if (numbers[i] <= min && numbers[i] != -1) {
      min = numbers[i];
      index = i;
    }
  }

  numbers[index] = -1;
  minNumbers.push(min);
  min = 1000;
}
console.log(minNumbers);

const assert = require('assert');

// Mock statistics calculation function
function calculateStatistics(matrix) {
  const flattened = matrix.flat();

  if (flattened.length === 0) {
    return {
      maxValue: null,
      minValue: null,
      average: null,
      totalSum: null,
      isDiagonal: false,
    };
  }

  const maxValue = Math.max(...flattened);
  const minValue = Math.min(...flattened);
  const totalSum = flattened.reduce((sum, val) => sum + val, 0);
  const average = totalSum / flattened.length;

  const isDiagonal = checkDiagonal(matrix);

  return {
    maxValue,
    minValue,
    average: parseFloat(average.toFixed(2)),
    totalSum,
    isDiagonal,
  };
}

function checkDiagonal(matrix) {
  if (matrix.length === 0) return false;

  const rows = matrix.length;
  const cols = matrix[0].length;

  if (rows !== cols) return false;

  for (let i = 0; i < rows; i++) {
    for (let j = 0; j < cols; j++) {
      if (i !== j && matrix[i][j] !== 0) {
        return false;
      }
    }
  }

  return true;
}

// Tests
console.log('Running tests...\n');

// Test 1: Basic 3x3 matrix
const matrix1 = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
];
const stats1 = calculateStatistics(matrix1);
assert.strictEqual(stats1.maxValue, 9, 'Test 1: Max value should be 9');
assert.strictEqual(stats1.minValue, 1, 'Test 1: Min value should be 1');
assert.strictEqual(stats1.average, 5, 'Test 1: Average should be 5');
assert.strictEqual(stats1.totalSum, 45, 'Test 1: Total sum should be 45');
assert.strictEqual(stats1.isDiagonal, false, 'Test 1: Should not be diagonal');
console.log('✓ Test 1 passed: Basic 3x3 matrix');

// Test 2: Diagonal matrix
const matrix2 = [
  [5, 0, 0],
  [0, 5, 0],
  [0, 0, 5],
];
const stats2 = calculateStatistics(matrix2);
assert.strictEqual(stats2.isDiagonal, true, 'Test 2: Should be diagonal');
assert.strictEqual(stats2.maxValue, 5, 'Test 2: Max value should be 5');
assert.strictEqual(stats2.minValue, 0, 'Test 2: Min value should be 0');
console.log('✓ Test 2 passed: Diagonal matrix');

// Test 3: Single element
const matrix3 = [[42]];
const stats3 = calculateStatistics(matrix3);
assert.strictEqual(stats3.maxValue, 42, 'Test 3: Max value should be 42');
assert.strictEqual(stats3.minValue, 42, 'Test 3: Min value should be 42');
assert.strictEqual(stats3.average, 42, 'Test 3: Average should be 42');
assert.strictEqual(stats3.totalSum, 42, 'Test 3: Total sum should be 42');
console.log('✓ Test 3 passed: Single element matrix');

// Test 4: Rectangular matrix
const matrix4 = [
  [1, 2, 3, 4],
  [5, 6, 7, 8],
];
const stats4 = calculateStatistics(matrix4);
assert.strictEqual(stats4.maxValue, 8, 'Test 4: Max value should be 8');
assert.strictEqual(stats4.minValue, 1, 'Test 4: Min value should be 1');
assert.strictEqual(stats4.totalSum, 36, 'Test 4: Total sum should be 36');
assert.strictEqual(stats4.isDiagonal, false, 'Test 4: Should not be diagonal (not square)');
console.log('✓ Test 4 passed: Rectangular matrix');

// Test 5: Empty matrix
const matrix5 = [];
const stats5 = calculateStatistics(matrix5);
assert.strictEqual(stats5.maxValue, null, 'Test 5: Max value should be null');
assert.strictEqual(stats5.minValue, null, 'Test 5: Min value should be null');
console.log('✓ Test 5 passed: Empty matrix');

// Test 6: Negative numbers
const matrix6 = [
  [-5, -2, 0],
  [3, -1, 4],
];
const stats6 = calculateStatistics(matrix6);
assert.strictEqual(stats6.maxValue, 4, 'Test 6: Max value should be 4');
assert.strictEqual(stats6.minValue, -5, 'Test 6: Min value should be -5');
console.log('✓ Test 6 passed: Matrix with negative numbers');

console.log('\n✅ All tests passed!');

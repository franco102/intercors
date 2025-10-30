const matrixService = {
  calculateStatistics: (matrix) => {
    const flattened = matrix.flat();

    if (flattened.length === 0) {
      return {
        maxValue: null,
        minValue: null,
        average: null,
        totalSum: null,
        isDiagonal: false,
        elementCount: 0
      };
    }

    const maxValue = Math.max(...flattened);
    const minValue = Math.min(...flattened);
    const totalSum = flattened.reduce((sum, val) => sum + val, 0);
    const average = totalSum / flattened.length;

    // Check if matrix is diagonal
    const isDiagonal = matrixService.checkDiagonal(matrix);

    return {
      maxValue,
      minValue,
      average: parseFloat(average.toFixed(2)),
      totalSum,
      isDiagonal,
      elementCount: flattened.length,
      dimensions: {
        rows: matrix.length,
        columns: matrix[0].length
      }
    };
  },

  checkDiagonal: (matrix) => {
    if (matrix.length === 0) return false;

    const rows = matrix.length;
    const cols = matrix[0].length;

    // A diagonal matrix must be square
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
};

module.exports = matrixService;

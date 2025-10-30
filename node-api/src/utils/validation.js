const validation = {
  validateMatrix: (matrix) => {
    if (!Array.isArray(matrix)) {
      return { isValid: false, error: 'Input must be an array' };
    }

    if (matrix.length === 0) {
      return { isValid: false, error: 'Matrix cannot be empty' };
    }

    const firstRowLength = matrix[0].length;

    for (let i = 0; i < matrix.length; i++) {
      const row = matrix[i];

      if (!Array.isArray(row)) {
        return { isValid: false, error: `Row ${i} is not an array` };
      }

      if (row.length !== firstRowLength) {
        return { isValid: false, error: 'All rows must have the same length' };
      }

      // Validate that all elements are numbers
      for (let j = 0; j < row.length; j++) {
        if (typeof row[j] !== 'number' || isNaN(row[j])) {
          return { isValid: false, error: `Element at [${i}][${j}] is not a valid number` };
        }
      }
    }

    return { isValid: true };
  }
};

module.exports = validation;

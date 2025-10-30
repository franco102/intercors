const matrixService = require('../services/matrixService');
const validation = require('../utils/validation');

const statisticsController = {
  calculateStatistics: (req, res) => {
    try {
      const { data, originalDiagonal } = req.body;

      // Validate input
      if (!data || !Array.isArray(data)) {
        return res.status(400).json({
          error: 'Invalid input. Expected an object with a "data" property containing an array of arrays.'
        });
      }

      // Validate matrix structure
      const validationResult = validation.validateMatrix(data);
      if (!validationResult.isValid) {
        return res.status(400).json({
          error: validationResult.error
        });
      }

      // Calculate statistics
      const statistics = matrixService.calculateStatistics(data);

      // Use originalDiagonal from Go API if provided
      if (originalDiagonal !== undefined) {
        statistics.isDiagonal = originalDiagonal;
      }

      res.json({
        rotatedMatrix: data,
        statistics,
        processedBy: 'node-api',
        user: req.user.username
      });
    } catch (error) {
      console.error('Statistics calculation error:', error);
      res.status(500).json({ error: 'Internal server error during statistics calculation' });
    }
  }
};

module.exports = statisticsController;

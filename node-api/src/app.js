const express = require('express');
const cors = require('cors');
const authMiddleware = require('./middleware/authMiddleware');
const errorMiddleware = require('./middleware/errorMiddleware');
const healthController = require('./controllers/healthController');
const authController = require('./controllers/authController');
const statisticsController = require('./controllers/statisticsController');

const app = express();

// Middleware
app.use(cors());
app.use(express.json());

// Public routes
app.get('/health', healthController.healthCheck);
app.post('/login', authController.login);

// Protected routes
app.post('/api/statistics', authMiddleware.verifyToken, statisticsController.calculateStatistics);

// 404 handler
app.use('*', (req, res) => {
  res.status(404).json({ error: 'Endpoint not found' });
});

// Error handling middleware
app.use(errorMiddleware.errorHandler);

module.exports = app;

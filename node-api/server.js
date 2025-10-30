const express = require('express');
const cors = require('cors');
const jwt = require('jsonwebtoken');
const app = express();

// JWT Secret
const JWT_SECRET = process.env.JWT_SECRET || 'your-secret-key-change-in-production';

// Middleware
app.use(cors());
app.use(express.json());

// JWT Middleware
function verifyToken(req, res, next) {
  const authHeader = req.headers['authorization'];
  if (!authHeader) {
    return res.status(401).json({ error: 'Missing authorization header' });
  }

  const token = authHeader.split(' ')[1]; // Extract token from "Bearer <token>"
  if (!token) {
    return res.status(401).json({ error: 'Invalid authorization header format' });
  }

  jwt.verify(token, JWT_SECRET, (err, decoded) => {
    if (err) {
      return res.status(401).json({ error: 'Invalid or expired token' });
    }
    req.user = decoded;
    next();
  });
}

// Calculate statistics from rotated matrix
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

  // Check if matrix is diagonal
  const isDiagonal = checkDiagonal(matrix);

  return {
    maxValue,
    minValue,
    average: parseFloat(average.toFixed(2)),
    totalSum,
    isDiagonal,
  };
}

// Check if a matrix is diagonal
function checkDiagonal(matrix) {
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

// Health check endpoint (public)
app.get('/health', (req, res) => {
  res.json({ status: 'healthy' });
});

// Login endpoint (public)
app.post('/login', (req, res) => {
  const { username, password } = req.body;

  // Simple authentication (in production, use proper password hashing)
  if (username === 'admin' && password === 'password') {
    const token = jwt.sign(
      { username },
      JWT_SECRET,
      { expiresIn: '24h' }
    );
    return res.json({ token });
  }

  res.status(401).json({ error: 'Invalid credentials' });
});

// Statistics endpoint (protected)
app.post('/api/statistics', verifyToken, (req, res) => {
  try {
    const { data, originalDiagonal } = req.body;

    if (!data || !Array.isArray(data)) {
      return res.status(400).json({
        error: 'Invalid input. Expected an object with a "data" property containing an array of arrays.',
      });
    }

    // Validate matrix structure
    if (data.length === 0) {
      return res.status(400).json({ error: 'Matrix cannot be empty' });
    }

    const firstRowLength = data[0].length;
    for (let row of data) {
      if (!Array.isArray(row) || row.length !== firstRowLength) {
        return res.status(400).json({
          error: 'Invalid matrix structure. All rows must have the same length.',
        });
      }
    }

    // Calculate statistics
    const statistics = calculateStatistics(data);
    
    // Use originalDiagonal from Go API if provided, otherwise check rotated matrix
    if (originalDiagonal !== undefined) {
      statistics.isDiagonal = originalDiagonal;
    }

    res.json({
      rotatedMatrix: data,
      statistics,
    });
  } catch (error) {
    console.error('Error processing request:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

// 404 handler
app.use((req, res) => {
  res.status(404).json({ error: 'Endpoint not found' });
});

// Error handler
app.use((err, req, res, next) => {
  console.error('Error:', err);
  res.status(500).json({ error: 'Internal server error' });
});

const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
  console.log(`Node.js API server running on port ${PORT}`);
});

const jwt = require('jsonwebtoken');

const JWT_SECRET = process.env.KEY_JWT || 'your-secret-key-change-in-production';

const authMiddleware = {
  verifyToken: (req, res, next) => {
    try {
      const authHeader = req.headers['authorization'];
      if (!authHeader) {
        return res.status(401).json({ error: 'Missing authorization header' });
      }

      const token = authHeader.split(' ')[1];
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
    } catch (error) {
      console.error('Auth middleware error:', error);
      res.status(500).json({ error: 'Authentication error' });
    }
  }
};

module.exports = authMiddleware;

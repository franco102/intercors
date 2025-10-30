const jwt = require('jsonwebtoken');

const JWT_SECRET = process.env.JWT_SECRET || 'your-secret-key-change-in-production';

const authController = {
  login: (req, res) => {
    try {
      const { username, password } = req.body;

      if (!username || !password) {
        return res.status(400).json({
          error: 'Username and password are required'
        });
      }

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
    } catch (error) {
      console.error('Login error:', error);
      res.status(500).json({ error: 'Internal server error during login' });
    }
  }
};

module.exports = authController;

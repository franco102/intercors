const healthController = {
  healthCheck: (req, res) => {
    res.json({
      status: 'healthy',
      timestamp: new Date().toISOString(),
      service: 'node-api'
    });
  }
};

module.exports = healthController;

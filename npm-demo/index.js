const express = require('express');
const axios = require('axios');

const app = express();
const PORT = 3000;

app.use(express.json());

// Root endpoint
app.get('/', (req, res) => {
  res.json({
    message: 'npm Demo Application',
    tool: 'npm',
    language: 'JavaScript',
    framework: 'Express',
    version: '1.0.0'
  });
});

// Info endpoint
app.get('/info', (req, res) => {
  res.json({
    dependencies: {
      express: '^4.18.2',
      axios: '^1.6.0'
    },
    description: 'REST API demonstrating npm package management',
    endpoints: ['/', '/info', '/external']
  });
});

// External API call using axios
app.get('/external', async (req, res) => {
  try {
    const response = await axios.get('https://api.github.com/repos/npm/cli');
    res.json({
      message: 'External API call successful',
      npmRepo: {
        name: response.data.name,
        stars: response.data.stargazers_count,
        language: response.data.language
      }
    });
  } catch (error) {
    res.status(500).json({ error: 'Failed to fetch external data' });
  }
});

app.listen(PORT, () => {
  console.log('=== npm Demo Application ===\n');
  console.log(`✅ Server running on http://localhost:${PORT}`);
  console.log('\nAvailable endpoints:');
  console.log(`  GET http://localhost:${PORT}/`);
  console.log(`  GET http://localhost:${PORT}/info`);
  console.log(`  GET http://localhost:${PORT}/external`);
  console.log('\nPress Ctrl+C to stop\n');
});

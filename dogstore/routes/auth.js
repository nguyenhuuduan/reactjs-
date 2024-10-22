const express = require('express');
const { register, login } = require('../controllers/authController');

const router = express.Router();

// Register new customer
router.post('/register', register);

// Login customer
router.post('/login', login);

module.exports = router;

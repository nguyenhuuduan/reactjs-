const Product = require('../models/product');

exports.createProduct = async (req, res) => {
  const { name, category_id, price, description, images, sizes, colors } = req.body;
  try {
    const newProduct = new Product({ name, category_id, price, description, images, sizes, colors });
    await newProduct.save();
    res.status(201).json(newProduct);
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
};

exports.getProducts = async (req, res) => {
  const { page = 1, limit = 10, search = '', category, minPrice, maxPrice } = req.query;
  try {
    const query = {
      name: { $regex: search, $options: 'i' },
      ...(category && { category_id: category }),
      ...(minPrice && { price: { $gte: minPrice } }),
      ...(maxPrice && { price: { $lte: maxPrice } }),
    };
    const products = await Product.find(query)
      .limit(Number(limit))
      .skip((Number(page) -


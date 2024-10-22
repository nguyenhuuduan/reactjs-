const mongoose = require('mongoose');

const productSchema = new mongoose.Schema({
  name: { type: String, required: true },
  category_id: { type: mongoose.Schema.Types.ObjectId, ref: 'Category', required: true },
  price: { type: Number, required: true },
  description: { type: String, required: true },
  images: [String],
  sizes: [String],
  colors: [String],
});

module.exports = mongoose.model('Product', productSchema);

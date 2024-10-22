const express = require('express');
const {
  createProduct,
  getProducts,
  getProductById,
  updateProduct,
  deleteProduct
} = require('../controllers/productController');

const router = express.Router();

/**
 * @swagger
 * components:
 *   schemas:
 *     Product:
 *       type: object
 *       required:
 *         - name
 *         - category_id
 *         - price
 *         - description
 *         - images
 *         - sizes
 *         - colors
 *       properties:
 *         id:
 *           type: string
 *           description: The auto-generated id of the product
 *         name:
 *           type: string
 *           description: The name of the product
 *         category_id:
 *           type: string
 *           description: The category id of the product
 *         price:
 *           type: number
 *           description: The price of the product
 *         description:
 *           type: string
 *           description: The description of the product
 *         images:
 *           type: array
 *           items:
 *             type: string
 *           description: The images of the product
 *         sizes:
 *           type: array
 *           items:
 *             type: string
 *           description: The sizes of the product
 *         colors:
 *           type: array
 *           items:
 *             type: string
 *           description: The colors of the product
 *       example:
 *         name: "Golden Retriever"
 *         category_id: "60b6ae8f5b3d3c44f8fae635"
 *         price: 500
 *         description: "A friendly and energetic dog"
 *         images: ["url1", "url2"]
 *         sizes: ["small", "medium"]
 *         colors: ["golden", "cream"]
 */

/**
 * @swagger
 * tags:
 *   name: Products
 *   description: The products managing API
 */

/**
 * @swagger
 * /products:
 *   post:
 *     summary: Create a new product
 *     tags: [Products]
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             $ref: '#/components/schemas/Product'
 *     responses:
 *       201:
 *         description: The product was successfully created
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/Product'
 *       500:
 *         description: Some server error
 */
router.post('/', createProduct);

/**
 * @swagger
 * /products:
 *   get:
 *     summary: Returns the list of all the products
 *     tags: [Products]
 *     parameters:
 *       - in: query
 *         name: page
 *         schema:
 *           type: integer
 *         description: The page number
 *       - in: query
 *         name: limit
 *         schema:
 *           type: integer
 *         description: The number of items to return
 *       - in: query
 *         name: search
 *         schema:
 *           type: string
 *         description: The search term
 *       - in: query
 *         name: category
 *         schema:
 *           type: string
 *         description: The category id to filter
 *       - in: query
 *         name: minPrice
 *         schema:
 *           type: number
 *         description: The minimum price
 *       - in: query
 *         name: maxPrice
 *         schema:
 *           type: number
 *         description: The maximum price
 *     responses:
 *       200:
 *         description: The list of the products
 *         content:
 *           application/json:
 *             schema:
 *               type: array
 *               items:
 *                 $ref: '#/components/schemas/Product'
 */
router.get('/', getProducts);

/**
 * @swagger
 * /products/{id}:
 *   get:
 *     summary: Get the product by id
 *     tags: [Products]
 *     parameters:
 *       - in: path
 *         name: id
 *         schema:
 *           type: string
 *         required: true
 *         description: The product id
 *     responses:
 *       200:
 *         description: The product description by id
 *         content:
 *           application/json:
 *             schema:
 *               $ref: '#/components/schemas/Product'
 *       404:
 *         description: The product was not found
 */
router.get('/:id', getProductById);

/**
 * @swagger
 * /products/{id}:
 *   put:
 *     summary: Update the product by id
 *     tags: [Products]
 *     parameters:
 *       - in: path
 *         name: id
 *         schema:
 *           type: string
 *         required: true
 *         description: The product id
 *     requestBody:
 *       required:

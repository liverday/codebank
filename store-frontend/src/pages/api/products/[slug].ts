import type { NextApiRequest, NextApiResponse } from 'next'
import { Product, products } from '../../model'

export default function handler(req: NextApiRequest, res: NextApiResponse<Product | { message: string}>) {
    const { slug } = req.query;
    const product: Product | undefined = products.find(({ slug: productSlug }) => productSlug === slug)
    
    if (!product) {
        res.status(404).json({ message: 'Product not found'})
    } else { 
        res.json(product)
    }
}
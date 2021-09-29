export interface Product {
    id: string;
    name: string;
    description: string;
    image_url: string;
    slug: string;
    price: number;
    created_at: String;
    updated_at: String
}

export const products: Product[] = [
    {
        id: 'uuid',
        name: 'Produto teste',
        description: 'muito texto descritivo ',
        price: 50.50,
        image_url: 'https://source.unsplash.com/random?product,' + Math.random(),
        slug: 'produto-teste',
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
    },
    {
        id: 'uuid2',
        name: 'Produto teste 2',
        description: 'muito texto descritivo 2',
        price: 50.50,
        image_url: 'https://source.unsplash.com/random?product,' + Math.random(),
        slug: 'produto-teste',
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
    },
]
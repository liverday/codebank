import React from "react";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  CardMedia,
  Typography,
} from "@material-ui/core";
import type { GetStaticPaths, GetStaticProps, NextPage } from "next";
import Head from "next/head";
import Link from "next/link";
import { Product } from "../../model";
import api from "../../../services/api";
import axios from "axios";

interface ProductDetailPageProps {
  product: Product;
}

const ProductDetailsPage: NextPage<ProductDetailPageProps> = ({ product }) => {
  console.log(product);
  return (
    <div>
      <Head>
        <title>{product.name} - Detalhes</title>
      </Head>
      <Card>
        <CardHeader
          title={product.name.toUpperCase()}
          subheader={`R$ ${product.price}`}
        />
        <CardActions>
          <Link
            href="/products/[slug]/order"
            as={`/products/${product.slug}/order`}
            passHref
          >
            <Button size="small" color="primary" component="a">
              Comprar
            </Button>
          </Link>
        </CardActions>
        <CardMedia style={{ paddingTop: "30%" }} image={product.image_url} />
        <CardContent>
          <Typography
            component="p"
            variant="body2"
            color="textSecondary"
            gutterBottom
          >
            {product.description}
          </Typography>
        </CardContent>
      </Card>
    </div>
  );
};

export const getStaticPaths: GetStaticPaths = async (ctx) => {
  const { data: products } = await api.get<Product[]>(`/products`);

  return {
    paths: products.map((product) => ({
      params: { slug: product.slug },
    })),
    fallback: "blocking",
  };
};

export const getStaticProps: GetStaticProps<
  ProductDetailPageProps,
  { slug: string }
> = async (ctx) => {
  const { slug } = ctx.params!;
  try {
    const { data: product } = await api.get(`/products/${slug}`);

    return {
      props: {
        product,
      },
      revalidate: 1 * 60 * 2,
    };
  } catch (e) {
    if (axios.isAxiosError(e) && e.response?.status === 404) {
      return { notFound: true };
    }

    throw e;
  }
};

export default ProductDetailsPage;

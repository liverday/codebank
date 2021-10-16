import React from "react";
import {
  Grid,
  Button,
  Card,
  CardActions,
  CardContent,
  CardMedia,
  Typography,
} from "@material-ui/core";
import type { GetServerSideProps, NextPage } from "next";
import Head from "next/head";
import Link from "next/link";

import api from "../services/api";

import { Product } from "./model";

interface ProductsListPageProps {
  products: Product[];
}

const ProductsListPage: NextPage<ProductsListPageProps> = ({ products }) => {
  return (
    <div>
      <Head>
        <title>Listagem de produtos</title>
      </Head>

      <Typography component="h1" variant="h3" color="textPrimary" gutterBottom>
        Produtos
      </Typography>

      <Grid container spacing={4}>
        {products.map((product, key) => (
          <Grid key={key} item xs={12} sm={6} md={4} lg={3}>
            <Card>
              <CardMedia
                style={{ paddingTop: "56%" }}
                image={product.image_url}
              />
              <CardContent>
                <Typography
                  component="h2"
                  variant="h5"
                  color="textPrimary"
                  gutterBottom
                >
                  {product.name}
                </Typography>
              </CardContent>
              <CardActions>
                <Link href="/products/[slug]" as={`/products/${product.slug}`} passHref>
                  <Button size="small" color="primary" component="a">
                    Detalhes
                  </Button>
                </Link>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
    </div>
  );
};

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const { data: products } = await api.get("products");
  return {
    props: {
      products,
    },
  };
};

export default ProductsListPage;

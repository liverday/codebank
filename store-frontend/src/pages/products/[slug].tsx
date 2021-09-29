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
import type { NextPage } from "next";
import Head from "next/head";
import { Product } from "../model";

interface ProductDetailPageProps {
  product: Product;
}

const ProductDetailsPage: NextPage<ProductDetailPageProps> = ({ product }) => {
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
          <Button size="small" color="primary" component="a">
            Detalhes
          </Button>
        </CardActions>
        <CardMedia style={{ paddingTop: '30%'}} image={product.image_url} />
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

export default ProductDetailsPage;

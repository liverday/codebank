import React from "react";
import {
  Button,
  Avatar,
  ListItem,
  ListItemAvatar,
  ListItemText,
  TextField,
  Typography,
  Grid,
  Box,
} from "@material-ui/core";

import type { GetServerSideProps, NextPage } from "next";
import Head from "next/head";
import { Product } from "../../model";
import api from "../../../services/api";
import axios from "axios";

interface OrderPageProps {
  product: Product;
}

const OrderPage: NextPage<OrderPageProps> = ({ product }) => {
  return (
    <div>
      <Head>
        <title>Pagamento</title>
      </Head>
      <Typography component="h1" variant="h3" color="textPrimary" gutterBottom>
        Checkout
      </Typography>
      <ListItem>
        <ListItemAvatar>
          <Avatar src={product.image_url} />
        </ListItemAvatar>
        <ListItemText
          primary={product.name}
          secondary={`R$ ${product.price}`}
        />
      </ListItem>
      <Typography component="h2" variant="h6" gutterBottom>
        Pague com cartão de crédito
      </Typography>
      <form>
        <Grid container spacing={3}>
          <Grid item xs={12} md={6}>
            <TextField required label="Nome" fullWidth />
          </Grid>
          <Grid item xs={12} md={6}>
            <TextField
              required
              inputProps={{ maxLength: 16 }}
              label="Número do Cartão"
              fullWidth
            />
          </Grid>
          <Grid item xs={12} md={6}>
            <TextField required type="number" label="CVV" fullWidth />
          </Grid>
          <Grid item xs={12} md={6}>
            <Grid container spacing={3}>
              <Grid item xs={6}>
                <TextField
                  required
                  type="number"
                  label="Expiração mês"
                  fullWidth
                />
              </Grid>
              <Grid item xs={6}>
                <TextField
                  required
                  type="number"
                  label="Expiração ano"
                  fullWidth
                />
              </Grid>
            </Grid>
          </Grid>
        </Grid>
        <Box marginTop={3}>
          <Button type="submit" variant="contained" color="primary" fullWidth>
            Pagar
          </Button>
        </Box>
      </form>
    </div>
  );
};

export const getServerSideProps: GetServerSideProps<
  OrderPageProps,
  { slug: string }
> = async (ctx) => {
  const { slug } = ctx.params!;
  try {
    const { data: product } = await api.get(`/products/${slug}`);

    return {
      props: {
        product,
      },
    };
  } catch (e) {
    if (axios.isAxiosError(e) && e.response?.status === 404) {
      return { notFound: true };
    }

    throw e;
  }
};

export default OrderPage;

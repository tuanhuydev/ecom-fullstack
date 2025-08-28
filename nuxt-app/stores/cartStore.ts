import { defineStore } from "pinia";
import type { Product } from "~/types/product";

export const useCartStore = defineStore("cart", {
  state: () => {
    if (typeof window !== 'undefined') {
      const storedCarts = localStorage.getItem("cart");
      if (!storedCarts) {
        return { carts: [] as Product[] };
      }
      const parsedCarts: { carts: [] } = JSON.parse((storedCarts as string));
      console.log(parsedCarts);
      if (!Array.isArray(parsedCarts.carts)) {
        return { carts: [] as Product[] };
      }

      return { carts: (parsedCarts).carts as Product[] };
    }
    return { carts: [] as Product[] };
  },
  actions: {
    addToCart(item: Product) {
      this.carts.push(item);
      if (typeof window !== 'undefined') {
        localStorage.setItem("cart", JSON.stringify(this));
      }
    }
  },
  getters: {
    cartCount: (state) => state.carts.length,
    cartTotal: (state) => state.carts.reduce((total, item) => total + (item as Product).price, 0),
  },

});
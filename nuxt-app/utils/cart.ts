export const CART_KEY = 'user_cart';

// Get the cart from local storage
export const getCart = (): unknown[] => {
  if (typeof window === 'undefined') {
    return []; // Return an empty cart during SSR
  }
  const cart = localStorage.getItem(CART_KEY);
  return cart ? JSON.parse(cart) : [];
};

// Save the cart to local storage
export const saveCart = (cart: unknown[]): void => {
  if (typeof window !== 'undefined') {
    localStorage.setItem(CART_KEY, JSON.stringify(cart));
  }
};

// Add an item to the cart
export const addToCart = (item: unknown): void => {
  const cart = getCart();
  cart.push(item);
  saveCart(cart);
};

// Remove an item from the cart
export const removeFromCart = (itemId: string): void => {
  const cart = getCart();
  const updatedCart = cart.filter((item) => (item as { id: string }).id !== itemId);
  saveCart(updatedCart);
};
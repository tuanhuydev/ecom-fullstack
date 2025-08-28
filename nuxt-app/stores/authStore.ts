import { defineStore } from "pinia";

interface User {
  id: string;
  email: string;
  name?: string;
}

interface AuthState {
  token: string | null;
  user: User | null;
  isAuthenticated: boolean;
}

export const useAuthStore = defineStore("auth", {
  state: (): AuthState => ({
    token: null,
    user: null,
    isAuthenticated: false,
  }),

  actions: {
    async login(email: string, password: string) {
      try {
        const config = useRuntimeConfig();
        // Make the API call to your backend
        const response: { token: string} = await $fetch(`${config.public.apiUrl}/auth/sign-in`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ email, password }),
        });
        if (!response.token) {
          throw new Error('Login failed');
        }

        
        // Store the token in localStorage for persistence
        localStorage.setItem('auth_token', response.token);
        
        // Update the store state
        this.token = response.token;
        this.isAuthenticated = true;
        
        // Return success
        return { success: true };
      } catch (error: unknown) {
        console.error('Login error:', error);
        return { 
          success: false, 
          error: (error as Error).message || 'Authentication failed' 
        };
      }
    },

    async logout() {
      // Clear token from localStorage
      if (typeof localStorage !== 'undefined') {
        localStorage.removeItem('auth_token');
      }
      
      // Reset the store state
      this.token = null;
      this.user = null;
      this.isAuthenticated = false;
    },

    // Check token on app initialization
    async initAuth() {
      if (typeof localStorage === 'undefined') return;
      
      const token = localStorage.getItem('auth_token');
      if (token) {
        this.token = token;
      }
    }
  },

  getters: {
    loggedIn: (state) => state.isAuthenticated,
    currentUser: (state) => state.user,
    authToken: (state) => state.token,
  },
});